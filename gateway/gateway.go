package main

import (
	"fmt"
	"gateway/lib"
	"gateway/router"
	"gateway/serviceclient"
	proto "gateway/submodules/common/protob"
	"gateway/submodules/common/tracer"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	wrapperTrace "github.com/opentracing/opentracing-go"
)

func main() {
	//注册服务
	registerService()
}

//注册服务
func registerService() {
	etcdAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Etcd.Port
	jaegerAddr := lib.Config.Etcd.Addr + ":" + lib.Config.Jaeger.Port
	webServiceAddr := ":" + lib.Config.GoMicro.Gateway.Port
	gatewayName := lib.Config.GoMicro.Gateway.Name
	gatewayWeb := lib.Config.GoMicro.Gateway.Web
	userSrvName := lib.Config.GoMicro.Srv.User.Name
	meetingApiName := lib.Config.GoMicro.Api.Meeting.Name

	//etcd注册实例
	etcdRegister := etcd.NewRegistry(
		registry.Addrs(etcdAddr),
	)
	//----------------注册网关-----------------------------------
	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(gatewayName, jaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	wrapperTrace.SetGlobalTracer(jaegerTracer)

	//注册网关服务
	//这个服务实际不会调用.run方法，实际会启动的是下面的webService
	gwtService := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "p",
				Usage: "port",
			},
		),
		micro.Name(gatewayName),
		micro.Version("latest"),
		// 配置etcd为注册中心，配置etcd路径，默认端口是2379
		micro.Registry(etcdRegister),
		// 配置链路追踪为 jaeger
		micro.WrapHandler(opentracing.NewHandlerWrapper(wrapperTrace.GlobalTracer())),
	)

	//解析命令行参数使用 -p 指定gateway服务启动端口
	//gwtService.Init(micro.Action(func(context *cli.Context) error {
	//	Port = context.String("p")
	//	if len(Port) == 0 {
	//		Port = "8888"
	//	}
	//	return nil
	//}))

	//---------------注册web服务-------------------------------
	//会议室预订服务的restful api映射
	webService := web.NewService(
		web.Name(gatewayWeb),
		web.Address(webServiceAddr),
		web.Registry(etcdRegister),
	)

	//绑定服务客户端
	//用户服务
	serviceclient.UserServiceClient = proto.NewUserService(userSrvName, gwtService.Client())
	//会议室预订api服务
	node, err := getSrvNode(etcdRegister, meetingApiName)
	if err != nil {
		log.Error("从注册中心获取 " + meetingApiName + " 失败")
	}
	serviceclient.MeetingApiNode = node

	//注册路由处理器
	webService.Handle("/", router.NewRouter())

	//启动服务
	if err := webService.Run(); err != nil {
		fmt.Println("webService.Run error:", err)
	}
}

//获取服务节点
func getSrvNode(reg registry.Registry, srvName string) (*registry.Node, error) {
	//获取服务列表
	services, err := reg.GetService(srvName)
	if err != nil {
		log.Info("未获取到服务 " + srvName + ",请确认服务是否存在")
		return nil, err
	}
	//获取随机服务
	next := selector.Random(services)
	node, err := next()
	if err != nil {
		log.Info("随机获取服务 " + srvName + " 实例失败")
		return nil, err
	}
	return node, nil
}
