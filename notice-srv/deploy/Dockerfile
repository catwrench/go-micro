FROM golang-base:1.14.4
ARG EXPOSE_PORT

ADD notice-srv /go/src/gomicro
WORKDIR /go/src/gomicro
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build --ldflags "-extldflags -static" -tags netgo -o notice-srv .

EXPOSE ${EXPOSE_PORT}
CMD [ "/go/src/gomicro/notice-srv" ]