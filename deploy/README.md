# docker部署流程

ps:开始前请确认

1、将common服务复制到每个项目的submodules路径下（使用git submodules引入公共模块，配置其实可以使用etcd存储）

2、是否将`submodules/common/config/config.dev.yml`重命名为`config.yml`，并填写正确参数

3、确认`deploy/.env`是否配置正确
***

### 先构建golang基础镜像，打上标签

```
cd golang
docker build -t golang:base .
docker tag golang:base golang-base:1.14.4
```

### 通过docker-compose 构建微服务镜像并启动电容器

```
//在deploy根目录执行
docker-compose up --build -d  
```

