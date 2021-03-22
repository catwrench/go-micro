# 基础镜像
FROM mysql:latest

# 设置时区
RUN echo "Asia/shanghai" > /etc/timezone;
RUN rm -rf /etc/localtime;
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

# 将本地mysql.conf.d下的文件复制到镜像中
ADD ./etc/mysql/my.cnf /etc/my.cnf


