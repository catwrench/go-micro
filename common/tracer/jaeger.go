package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	config := jaegercfg.Configuration{
		ServiceName: serviceName,
		//配置采样器
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		//配置报告间隔
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	//创建udp发送实例
	sender, err := jaeger.NewUDPTransport(addr, 0)
	if err != nil {
		return nil, nil, err
	}

	reporter := jaeger.NewRemoteReporter(sender)

	//创建链路追踪实例
	tracer, closer, err := config.NewTracer(
		jaegercfg.Reporter(reporter),
	)

	return tracer, closer, err
}
