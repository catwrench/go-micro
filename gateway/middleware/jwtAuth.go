package middleware

import (
	"fmt"
	"gateway/lib"
	"gateway/serviceclient"
	proto "gateway/submodules/common/protob"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"strconv"
)

//验证token中间件
func ValidateTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取请求头的token
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			log.Info("lib.NewResponse(1006): token不存在或已过期")
			ctx.JSON(200, lib.NewResponse(1006))
			ctx.Abort()
			return
		}

		//远程调用用户服务进行验证
		var cid string
		if ctx.Request.Method == "GET" {
			//get请求从参数中获取c_id
			cid = ctx.Query("c_id")
		} else {
			//非get请求从body中获取c_id
			type req struct {
				Cid int64 `json:"c_id"`
			}
			var r req
			if err := lib.GetFromGinBody(ctx, &r); err != nil {
				ctx.JSON(200, lib.NewResponse(1001).WithMsg("获取c_id失败"))
				ctx.Abort()
				return
			}
			cid = strconv.FormatInt(r.Cid, 10)
		}
		fmt.Println("cid:", cid, "METHOD:", ctx.Request.Method)

		//添加链路追踪子span
		c, ok := lib.ContextWithSpan(ctx)
		if ok == false {
			log.Error("get spanContext err")
		}
		//远程调用UserService.ValidateToken
		res, err := serviceclient.UserServiceClient.ValidateToken(c, &proto.ReqValidateToken{Token: token, CId: cid})
		fmt.Println("ValidateToken res:", res, "| cid:", cid, "| token:", token)

		//远程调用UserService异常
		if err != nil {
			log.Info("远程调用UserService.ValidateToken失败:", err)
			ctx.JSON(500, lib.NewResponse(30100))
			ctx.Abort()
			return
		}
		//返回结果非200
		if res.Code != lib.Err.GetInt64("error.ok") {
			ctx.JSON(200, lib.NewResponse(res.Code).WithMsg(res.Message))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
