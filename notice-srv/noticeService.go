package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	wrapperTrace "github.com/opentracing/opentracing-go"
	_ "notice-srv/database"
	"notice-srv/handler"
	"notice-srv/lib"
	proto "notice-srv/submodules/common/protob"
	"notice-srv/submodules/common/tracer"
	"time"
)

func main() {

	//----------------
	etcdAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Etcd.Port
	jaegerAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Jaeger.Port
	noticeSrvName := lib.Config.GoMicro.Srv.Notice.Name
	noticeSrvPort := ":" + lib.Config.GoMicro.Srv.Notice.Port

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(noticeSrvName, jaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	wrapperTrace.SetGlobalTracer(jaegerTracer)

	service := micro.NewService(
		micro.Name(noticeSrvName),
		micro.Address(noticeSrvPort),
		micro.Version("latest"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(etcdAddr),
		)),
		micro.RegisterTTL(time.Second*30),      //注册延迟，30s 内没有注册则失效，etcd 会自动删除服务
		micro.RegisterInterval(time.Second*20), //注册间隔，每隔 20s 注册一次
		// 配置链路追踪为 jaeger
		micro.WrapHandler(opentracing.NewHandlerWrapper(wrapperTrace.GlobalTracer())),
	)
	service.Init()

	//将处理器注册到服务上
	if err := proto.RegisterNoticeServiceHandler(service.Server(), new(handler.NoticeService)); err != nil {
		log.Error(noticeSrvName, " handler.NoticeService 绑定处理器到服务失败:", err)
		return
	}

	if err := service.Run(); err != nil {
		log.Info("服务启动失败:", err)
		return
	}
}
