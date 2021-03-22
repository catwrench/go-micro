# 基于centos:7.8.2003构建golang环境
FROM centos:7.8.2003

# 安装golang
COPY ./install-golang.sh /root/install-golang.sh
RUN chmod +x /root/install-golang.sh
RUN sh /root/install-golang.sh
RUN rm -rf /root/install-golang.sh
ENV PATH /usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GOROOT /usr/local/go
ENV GOPATH /root/gopath