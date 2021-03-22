FROM golang-base:1.14.4
ARG EXPOSE_PORT

ADD gateway /go/src/gomicro
WORKDIR /go/src/gomicro
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
# -o :指定编译输出为user-srv，代替默认包名
# -tags :构建出带tag的版本
# --ldflags :"-extldflags -static"静态编译
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags "-extldflags -static" -tags netgo -o gateway .

EXPOSE ${EXPOSE_PORT}
CMD [ "/go/src/gomicro/gateway" ]