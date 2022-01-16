# Module 3 Homework

- 构建本地镜像
- 编写 `Dockerfile` 将练习 2.2 编写的 `httpserver` 容器化
    
采用对阶段构建

[Dockerfile](../../module02/homework/httpserver/Dockerfile)
```dockerfile
FROM golang:1.17 AS build

RUN mkdir /app
COPY . /app
WORKDIR /app
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go mod download && GOOS=linux go build -installsuffix cgo -o httpserver ./cmd/main.go

###
FROM scratch as final
COPY --from=build /app/httpserver .
EXPOSE 80
ENTRYPOINT ["/httpserver"]
```
```shell
$ cd module02/httpserver
$ docker build . -t my-httpserver:0.1.1
$ docker images

REPOSITORY      TAG       IMAGE ID       CREATED         SIZE
my-httpserver   0.1.1     e9b767f9d082   8 minutes ago   6.39MB
<none>          <none>    70093f634e28   8 minutes ago   999MB
golang          1.17      8b86bf336a01   9 days ago      941MB
 ```

- 将镜像推送至 `docker` 官方镜像仓库
```shell
$ docker login --username=<muhubusername> --email=my@163.com

Login Succeeded

$ docker push my-httpserver:0.1.1
```
- 通过 `docker` 命令本地启动 `httpserver`

```shell
$ docker run -d -p 8080:80 my-httpserver:0.1.1

# test httpserver
$ curl -i -X GET http://localhost:8080
HTTP/1.1 200 OK
Accept: */*
User-Agent: curl/7.71.1
Version:
Date: Sun, 16 Jan 2022 12:28:30 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8

Hello World%
```
- 通过 `nsenter` 进入容器查看 IP 配置
```shell
# in Ubuntu 20.04
$ docker ps

CONTAINER ID   IMAGE                 COMMAND         CREATED         STATUS         PORTS                                   NAMES
263aa299962e   my-httpserver:0.1.1   "/httpserver"   7 minutes ago   Up 7 minutes   0.0.0.0:8080->80/tcp, :::8080->80/tcp   magical_mclaren

$ PID=$(docker inspect --format "{{ .State.Pid }}" 263aa299962e)
$ sudo nsenter -t $PID -n ip a

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
8: eth0@if9: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
