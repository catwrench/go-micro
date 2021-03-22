# 悦来微服务
***

## 服务部署说明 
ps:建议`deploy/docker-compose`部署，查看`deploy/README.md`可以跳过以下步骤

### etcd 注册中心安装
- docker安装：
```
docker run -d --name etcd \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd:2379 \
    bitnami/etcd:latest
```
### jaeger 链路追踪安装
官方文档：https://www.jaegertracing.io/docs/1.21/deployment/
- docker安装
链路追踪查看地址：http://localhost:16686/
```
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 14250:14250 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest
```
***

## 启动服务
ps: 
```
1、启动服务的命令都需要到相应服务根目录下执行
2、下面的服务按顺序启动即可
3、服务启动前先在服务根目录执行命令更新依赖：`go mod tidy`
```

- 启动 etcd服务的web管理界面 【需要查看才启动】
访问 http://localhost:8082 可以查看已注册的服务
```
micro --registry=etcd --registry_address=127.0.0.1:2379 web
```

- 启动 gateway网关
```
cd gateway
go run . -p 8888 --registry=etcd
```

- 启动 会议室预订服务[api]
```
cd meeting-api
go run . --registry=etcd
```

- 启动 会议室预订服务[server]
```
cd meeting-srv
go run . --registry=etcd
```

- 启动 消息通知服务[server]
```
cd notice-srv
go run . --registry=etcd
```

- curl调用会议室接口进行测试
```
curl -X POST -d "username=yw&mobile=18612345678" http://127.0.0.1:8091/meeting-api/meeting/test
```

***

# 引入的第三方组件说明
- git submodules
使用参考：https://www.jianshu.com/p/a47493cb31b0

- protobuf
使用参考：https://juejin.cn/post/6844904147339198472#heading-4
缺失值问题解决，可以使用one_of: https://developers.google.com/protocol-buffers/docs/proto3#oneof

在common根目录下生成所有proto文件命令：
```
for x in **/*.proto; do protoc --go_out=protob --micro_out=protob $x; done
```

protobuf生成的文件管理方案：

参考：http://www.liuhaihua.cn/archives/681093.html
```
A. 在不同的 server / client 代码项目中复制 proto 文件并各自生成
B. 使用 git submodule 使用同一个仓库的 proto 文件
C. 使用 semantic version 并使用第三方的方式管理 proto 文件
```

Proto 语言文件的规范：
```
proto 文件遵循只增不减的原则
proto 文件中的接口遵循只增不减的原则
proto 文件中的 message 字段遵循只增不减的原则
proto 文件中的 message 字段类型和序号不得修改
```

- gorm组件

操作数据库的orm框架-gorm
官方文档: https://gorm.io/zh_CN/docs/
逆向表到gorm模型：https://github.com/xxjwxc/gormt/blob/master/README_zh_cn.md

- 字段参数验证库-validator

repo: https://github.com/go-playground/validator 
使用:  https://www.cnblogs.com/jiujuan/p/13823864.html
***

## 服务命名规范（建议）
【go.micro.srv.student】
- 参数说明：
```
前两位：固定为`go.micro`
第三位：服务类型：`api(web应用，提供给外部调用的api服务)、rpc(远程过程调用)、proxy(反向代理)、event(事件)、handler`
第四位：服务名称
```

***

# 备注
- http调用服务
```
req := myClient.NewRequest("serviceName", "/route", requestType)
err := myClient.Call(context.Background(), req, &rsp)
```

- rpc调用服务
需要client和server同时持有.pb.go和.pb.micro.go，handler和服务进行绑定

- gorm的数据库model和proto

gorm的数据库model和proto定义的消息体结构不兼容（因为proto不支持自定义标签）
```
解决方案一（推荐）：需要分开定义gorm和proto的两个model,再使用中间层进行转换，遇到时间和多条数据的情况需要单独处理，算是比较灵活的方案
解决方案二：不用gorm,自己写sql
解决方案三：不使用rpc,用http完成服务间调用，就能避免使用proto文件
```

***

已知问题：
```
1、jaeger的链路追踪有点问题，只能追踪到两个服务，推测和注册时用的wrapper类型有关（已处理）
2、微服务传输层的数据结构还为调整为json，接口请求参数暂时使用body的form-data（已处理）
3、公共配置暂时使用统一定义const文件，每个服务独立，后续考虑将公共配置集中到同一个项目，
    使用git submodule将公共项目作为依赖引入（已处理）
4、网关进行反向代理的配置目前是写死的，导致网关必须在api服务后面启动，api服务重启后端口会变动，网关也需要重启一次才能正常访问。
    临时解决办法是将每个服务的端口固定。(已处理，通过固定服务端口)
5、gorm model好像可以直接映射protobuf字段，后续进行测试 （https://www.bilibili.com/video/BV1LT4y1M7mV?p=18）
```

踩坑记录：
1、protobuf缺省值问题
原因：proto3传输时，0、""、null会被忽略
解决：可以使用one_of,并且将字段类型设置为string,避免int类型默认值被设置为0 
https://developers.google.com/protocol-buffers/docs/proto3#oneof

2、gorm使用update进行更新struct时忽略了0值
原因：update存储时会先将struct转map,转换过程中0值会被忽略
解决一：使用save方法进行保存
解决二：使用update方法进行保存，但传入参数手动转换为map类型

3、引入viper组件，读取配置文件后，开启调试模式报错
原因：ide配置错误
解决：debug配置->go build->Run kind(package)->working directory(debug服务的根目录)

4、读取ctx.request.body时报错"unexpected EOF"
原因：将数据重新写入body，因为readall读取一次后就不在了，会导致后面的api服务读取body是报错"unexpected EOF"
解决：ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))