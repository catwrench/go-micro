package router

import (
	"gateway/lib"
	"gateway/middleware"
	"gateway/serviceclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	//跨域处理
	route.Use(middleware.Cors())

	//反向代理用户中心接口
	route.Any("/saasuims/uims/*any", ReverseProxy(lib.Config.Url.Uims.Host, lib.Config.Url.Uims.Scheme))

	route.Use(
		lib.JaegerMiddleware(),               //jaeger中间件
		middleware.ValidateTokenMiddleware(), //token验证中间件(依赖jaeger)
	)

	//通配路由，以 meeting 为前缀的都转发到会议室预订服务去
	uriMeeting := serviceclient.MeetingApiNode.Address
	route.Any("/api/meeting/*any", ReverseProxy(uriMeeting, ""))
	route.Any("/api/meetingApplet/*any", ReverseProxy(uriMeeting, ""))

	return route
}

//反向代理-地址
//host eg: "www.baidu.com" , "https"
func ReverseProxy(host string, scheme string) gin.HandlerFunc {
	return func(context *gin.Context) {
		director := func(req *http.Request) {
			if scheme == "" {
				scheme = "http"
			}
			req.URL.Scheme = scheme
			req.URL.Host = host
			req.Host = host //一个ip对应多个域名的情况需要设置这项
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(context.Writer, context.Request)
	}
}

//NewSingleHostReverseProxy导致线上代理循环
//func ReverseProxy(target string) gin.HandlerFunc {
//	return func(context *gin.Context) {
//		u, err := url.Parse(target)
//		if err != nil {
//			log.Error("ReverseProxy target:", target)
//			panic(err)
//		}
//		fmt.Println("lib.Config.Url.Uims:", lib.Config.Url.Uims)
//		fmt.Println("ReverseProxy target:", context.Request.URL.Path)
//		fmt.Println("ReverseProxy u:", context.Request.URL.Path)
//		proxy := httputil.NewSingleHostReverseProxy(u)
//		proxy.ServeHTTP(context.Writer, context.Request)
//	}
//}
