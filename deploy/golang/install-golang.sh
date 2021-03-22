# !/bin/bash
mkdir /root/gopath
yum install -y wget
cd /usr/local/ && wget https://golang.google.cn/dl/go1.14.4.linux-amd64.tar.gz
tar zxf go1.14.4.linux-amd64.tar.gz && rm -rf go1.14.4.linux-amd64.tar.gz
echo "export GOROOT=/usr/local/go" >> /etc/profile
echo "export GOPATH=/root/gopath" >> /etc/profile
echo "export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin" >> /etc/profile
source /etc/profile
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct