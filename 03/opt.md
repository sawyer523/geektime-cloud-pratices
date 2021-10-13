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
03> docker inspect 203a430ca0
[
    {
        "Id": "203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d",
        "Created": "2021-10-13T15:06:51.289115506Z",
        "Path": "/opt/httpserver/httpserver",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 222842,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-10-13T15:06:51.623131044Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:01395404376526c6af77cfeb67f33cc0188b397e0fac16239e066289e266d01b",
        "ResolvConfPath": "/var/lib/docker/containers/203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d/hostname",
        "HostsPath": "/var/lib/docker/containers/203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d/hosts",
        "LogPath": "/var/lib/docker/containers/203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d/203a430ca01f8e7e62c3bce18212aafd2a7a2649efb8466d0510b436ba74197d-json.log",
        "Name": "/sleepy_engelbart",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "docker-default",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "CgroupnsMode": "host",
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/947e4f5fcc40022fa25ccbe47d2b4a9791634343de01fdd31ed6857ca6e7b819-init/diff:/var/lib/docker/overlay2/0a725266a4bab4c3246b354f3036abba20b299e9bca2dcf96f0d4487960ded3d/diff:/var/lib/docker/overlay2/109077f5b6ae4059ed2b2f358802efeb448e0816d18f5a550d6db10e08308b52/diff:/var/lib/docker/overlay2/71c8433034e8557266ad95e8fc941a99f4039b4a94670b59e48e5f7e1e0cb3c5/diff",
                "MergedDir": "/var/lib/docker/overlay2/947e4f5fcc40022fa25ccbe47d2b4a9791634343de01fdd31ed6857ca6e7b819/merged",
                "UpperDir": "/var/lib/docker/overlay2/947e4f5fcc40022fa25ccbe47d2b4a9791634343de01fdd31ed6857ca6e7b819/diff",
                "WorkDir": "/var/lib/docker/overlay2/947e4f5fcc40022fa25ccbe47d2b4a9791634343de01fdd31ed6857ca6e7b819/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "203a430ca01f",
            "Domainname": "",
            "User": "",
            "AttachStdin": false,
            "AttachStdout": false,
            "AttachStderr": false,
            "ExposedPorts": {
                "8080/tcp": {}
            },
            "Tty": true,
            "OpenStdin": true,
            "StdinOnce": false,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": null,
            "Image": "httpserver:latest",
            "Volumes": null,
            "WorkingDir": "/opt/httpserver",
            "Entrypoint": [
                "/opt/httpserver/httpserver"
            ],
            "OnBuild": null,
            "Labels": {
                "name": "httpserver"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "648f376a6c8edf95c21656be01be773fac2a82db3049aa42ee023ee44395718d",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "8080/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/648f376a6c8e",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "8d02f2ca1464fa81eee60ee682bc1036c16a1c1f1abb113ef3319ee043841fb4",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "cacca2ae5c5f3ad35ab3d6a997e8412368d1126791971749d84b9ea742bd0c27",
                    "EndpointID": "8d02f2ca1464fa81eee60ee682bc1036c16a1c1f1abb113ef3319ee043841fb4",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]

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
