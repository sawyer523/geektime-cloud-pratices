# 操作步骤
```shell
make build.docker
```
执行结果
```shell
# shellcheck disable=SC2084
[+] Building 1.5s (8/8) FINISHED                                                                                                                                            
 => [internal] load build definition from Dockerfile                                                                                                                   0.0s
 => => transferring dockerfile: 37B                                                                                                                                    0.0s
 => [internal] load .dockerignore                                                                                                                                      0.0s
 => => transferring context: 2B                                                                                                                                        0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                                                                       1.3s
 => [1/3] FROM docker.io/library/alpine@sha256:e1c082e3d3c45cccac829840a25941e679c25d438cc8412c2fa221cf1a824e6a                                                        0.0s
 => [internal] load build context                                                                                                                                      0.1s
 => => transferring context: 6.22MB                                                                                                                                    0.1s
 => CACHED [2/3] WORKDIR /opt/httpserver                                                                                                                               0.0s
 => CACHED [3/3] COPY httpserver /opt/httpserver/                                                                                                                      0.0s
 => exporting to image                                                                                                                                                 0.0s
 => => exporting layers                                                                                                                                                0.0s
 => => writing image sha256:aa357dfe6f15d47623832dae087d3171fb9c539d0a45ec9468a8b259face084d                                                                           0.0s
 => => naming to docker.io/library/httpserver:v1.0.0.4                                                                                                                 0.0s

```
运行容器
```shell
 docker run -dit httpserver:v1.0.0.4
```
执行结果
```shell
e56234aa054b4c5e23509bcc9f4973e701377440200dd77d3dd919851017bd1d
```
推送到docker仓库
```shell
docker login
docker tag httpserver:v1.0.0.4 sawyer523/httpserver:v1.0.0.4
docker push sawyer523/httpserver:v1.0.0.4

The push refers to repository [docker.io/sawyer523/httpserver]
27a5bd1040d3: Pushed 
bc27313a33a4: Pushed 
e2eb06d8af82: Pushed 
v1.0.0.4: digest: sha256:2bc4fe6669d18fb248172a2c5fdd526fa0b6de4351995676a72b55a0b30c9b5e size: 946
```

进入容器
```shell
nsenter e56234a
```