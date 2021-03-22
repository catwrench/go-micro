package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

const contextTracerKey = "Tracer-context"

//jaeger中间件，记录token和span信息
func JaegerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//元数据metadata
		md := make(map[string]string)
		//获取请求投的 spanCtx
		spanCtx, sErr := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if sErr != nil {
			log.Info("sErr:", sErr)
		}
		//基于提取的 spanCtx 创建新的子span
		sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		defer sp.Finish()

		//注入span到tracer
		if err := opentracing.GlobalTracer().Inject(
			sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md)); err != nil {
			log.Log(err)
		}

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		ctx = metadata.NewContext(ctx, md)
		c.Set(contextTracerKey, ctx)

		c.Next()

		//通过ext可以为追踪设置额外的一些信息
		statusCode := c.Writer.Status()
		ext.HTTPStatusCode.Set(sp, uint16(statusCode))
		ext.HTTPMethod.Set(sp, c.Request.Method)
		ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
		if statusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		}
	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get(contextTracerKey)
	if exist == false {
		ok = false
		ctx = context.TODO()
		return
	}

	ctx, ok = v.(context.Context)
	return
}
