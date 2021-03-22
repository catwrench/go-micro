package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	wrapperTrace "github.com/opentracing/opentracing-go"
	_ "meeting-srv/database"
	"meeting-srv/handler"
	"meeting-srv/lib"
	"meeting-srv/serviceclient"
	proto "meeting-srv/submodules/common/protob"
	"meeting-srv/submodules/common/tracer"
	"time"
)

func main() {

	etcdAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Etcd.Port
	jaegerAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Jaeger.Port
	meetingSrvName := lib.Config.GoMicro.Srv.Meeting.Name
	meetingSrvPort := ":" + lib.Config.GoMicro.Srv.Meeting.Port
	noticeSrvName := lib.Config.GoMicro.Srv.Notice.Name
	userSrvName := lib.Config.GoMicro.Srv.User.Name

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(meetingSrvName, jaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	wrapperTrace.SetGlobalTracer(jaegerTracer)

	//服务初始化
	service := micro.NewService(
		micro.Name(meetingSrvName),
		micro.Address(meetingSrvPort),
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

	//绑定会议室预订服务
	if err := proto.RegisterReservationServiceHandler(service.Server(), new(handler.ReservationService)); err != nil {
		log.Error(meetingSrvName, " 绑定服务 ReservationService 失败:", err)
		return
	}
	if err := proto.RegisterSpaceServiceHandler(service.Server(), new(handler.SpaceService)); err != nil {
		log.Error(meetingSrvName, " 绑定服务 SpaceService 失败:", err)
		return
	}
	if err := proto.RegisterDeviceServiceHandler(service.Server(), new(handler.DeviceService)); err != nil {
		log.Error(meetingSrvName, " 绑定服务 DeviceService 失败:", err)
		return
	}
	if err := proto.RegisterRoomServiceHandler(service.Server(), new(handler.RoomService)); err != nil {
		log.Error(meetingSrvName, " 绑定服务 RoomService 失败:", err)
		return
	}

	//给 NoticeServiceClient 绑定实例
	serviceclient.NoticeServiceClient = proto.NewNoticeService(noticeSrvName, service.Client())
	serviceclient.UserServiceClient = proto.NewUserService(userSrvName, service.Client())

	if err := service.Run(); err != nil {
		log.Info("服务启动失败:", err)
		return
	}
}
