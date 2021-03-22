package lib

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"math/rand"
	"net/http"
)

//参考：https://studygolang.com/articles/25337
//https://www.ctolib.com/xinliangnote-go-jaeger-demo.html

const contextTracerKey = "Tracer-context"
//采样率 sf sampling frequency
var sf = 100
// 0 <= n <= 100
func SetSamplingFrequency(n int) {
	sf = n
}

//jaeger中间件，记录token和span信息
func JaegerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
		tracer := opentracing.GlobalTracer()
		//元数据metadata
		md := make(map[string]string)
		//获取请求投的 spanCtx
		spanCtx, sErr := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if sErr != nil {
			sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
			tracer = sp.Tracer()
		}
		//基于提取的 spanCtx 创建新的子span
		sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
		sp.SetTag("Authorization", c.GetHeader("Authorization"))
		defer sp.Finish()

		//注入span到tracer.text
		if err := tracer.Inject(
			sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md),
		); err != nil {
			log.Log(err)
		}
		//注入span到tracer.text
		if err := tracer.Inject(
			sp.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		); err != nil {
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
		} else if rand.Intn(100) > sf {
			ext.SamplingPriority.Set(sp, 0)
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
