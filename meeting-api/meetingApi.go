package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	wrapperTrace "github.com/opentracing/opentracing-go"
	"meeting-api/lib"
	"meeting-api/router"
	"meeting-api/serviceclient"
	proto "meeting-api/submodules/common/protob"
	"meeting-api/submodules/common/tracer"
	"time"
)

func main() {
	etcdAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Etcd.Port
	jaegerAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Jaeger.Port
	meetingApiName := lib.Config.GoMicro.Api.Meeting.Name
	meetingApiPort := ":" + lib.Config.GoMicro.Api.Meeting.Port
	meetingSrvName := lib.Config.GoMicro.Srv.Meeting.Name

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(meetingApiName, jaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	wrapperTrace.SetGlobalTracer(jaegerTracer)

	//服务初始化
	service := micro.NewService(
		micro.Name(meetingApiName),
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

	//http服务初始化
	httpService := web.NewService(
		web.Name(meetingApiName),
		web.Address(meetingApiPort),
		web.Version("latest"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		web.Registry(etcd.NewRegistry(
			registry.Addrs(etcdAddr),
		)),
		web.RegisterTTL(time.Second*30),      //注册延迟，30s 内没有注册则失效，etcd 会自动删除服务
		web.RegisterInterval(time.Second*20), //注册间隔，每隔 20s 注册一次
		web.Handler(router.NewRouter()),
	)

	//绑定服务
	serviceclient.ReservationServiceClient = proto.NewReservationService(meetingSrvName, service.Client())
	serviceclient.SpaceServiceClient = proto.NewSpaceService(meetingSrvName, service.Client())
	serviceclient.DeviceServiceClient = proto.NewDeviceService(meetingSrvName, service.Client())
	serviceclient.RoomServiceClient = proto.NewRoomService(meetingSrvName, service.Client())

	//http服务启动
	if err := httpService.Run(); err != nil {
		log.Info("服务启动失败:", err)
		return
	}
}
