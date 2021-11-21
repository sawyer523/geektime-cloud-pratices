# 操作步骤
```shell
make build.docker
```
执行结果
```shell
# shellcheck disable=SC2084
 03> docker build -t httpserver .
Sending build context to Docker daemon  23.91MB
Step 1/13 : FROM golang:1.17.2-alpine3.14 AS builder
 ---> 35cd8c8897b1
Step 2/13 : WORKDIR /go/src
 ---> Using cache
 ---> 600559829a32
Step 3/13 : COPY go.mod .
 ---> Using cache
 ---> f94f80ca6e7d
Step 4/13 : COPY main.go .
 ---> Using cache
 ---> b7f10e7a8aba
Step 5/13 : RUN go mod vendor
 ---> Using cache
 ---> 0b043abb0c39
Step 6/13 : ENV CGO_ENABLED=0
 ---> Using cache
 ---> 3cb142bccd25
Step 7/13 : RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/httpserver main.go
 ---> Using cache
 ---> 6dbddab8fbc5
Step 8/13 : FROM alpine:3.14
 ---> 14119a10abf4
Step 9/13 : LABEL name="httpserver"
 ---> Running in 7985545939fc
Removing intermediate container 7985545939fc
 ---> 01b054a9c8c8
Step 10/13 : WORKDIR /opt/httpserver
 ---> Running in da015133e6a8
Removing intermediate container da015133e6a8
 ---> 8792a798160a
Step 11/13 : COPY --from=builder /go/bin/httpserver /opt/httpserver/
 ---> fa2e67417f25
Step 12/13 : EXPOSE 8080
 ---> Running in d1f84933b2e0
Removing intermediate container d1f84933b2e0
 ---> 24cdd313893d
Step 13/13 : ENTRYPOINT  ["/opt/httpserver/httpserver"]
 ---> Running in 952052f0b640
Removing intermediate container 952052f0b640
 ---> 013954043765
Successfully built 013954043765
Successfully tagged httpserver:v1.0.0.1
```
运行容器
```shell
 docker run -dit httpserver:v1.0.0.1
```
执行结果
```shell
203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d
```
推送到docker仓库
```shell
docker login
docker tag httpserver:v1.0.0.1 sawyer523/httpserver:v1.0.0.1
docker push sawyer523/httpserver:v1.0.0.1

The push refers to repository [docker.io/sawyer523/httpserver]
27a5bd1040d3: Pushed 
bc27313a33a4: Pushed 
e2eb06d8af82: Pushed 
v1.0.0.4: digest: sha256:2bc4fe6669d18fb248172a2c5fdd526fa0b6de4351995676a72b55a0b30c9b5e size: 946
```

进入容器
```shell
03> docker inspect -f '{{ .State.Pid }}' 203a430ca0
222842

03> chsh -s /bin/bash
root@k8s:/home/cxn# nsenter --target 240808 --mount --uts --ipc --net --pid su - root
203a430ca01f:~# cd /opt/httpserver/
203a430ca01f:/opt/httpserver# ls
httpserver
203a430ca01f:/opt/httpserver#
203a430ca01f:/opt/httpserver# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
94: eth0@if95: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```
