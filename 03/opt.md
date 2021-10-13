# 操作步骤
```shell
make build.docker
```
执行结果
```shell
# shellcheck disable=SC2084
 03> docker build -t httpserver .
Sending build context to Docker daemon  23.91MB
Step 1/13 : FROM golang:1.17.2-stretch AS builder
 ---> f4af25eac23b
Step 2/13 : WORKDIR /go/src
 ---> Running in f0fb39852a32
Removing intermediate container f0fb39852a32
 ---> 30c9e2a2cb18
Step 3/13 : COPY go.mod .
 ---> fb5cfc10ab9d
Step 4/13 : COPY main.go .
 ---> 537b5d154148
Step 5/13 : RUN go mod vendor
 ---> Running in 83686aeb3778
go: downloading github.com/sirupsen/logrus v1.8.1
go: downloading golang.org/x/sys v0.0.0-20191026070338-33540a1f6037
Removing intermediate container 83686aeb3778
 ---> b3a5adb02dee
Step 6/13 : ENV CGO_ENABLED=0
 ---> Running in 9c85c8261dd5
Removing intermediate container 9c85c8261dd5
 ---> 211823e6f9ef
Step 7/13 : RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/httpserver main.go
 ---> Running in c16816c99f27
Removing intermediate container c16816c99f27
 ---> 70913b96ac31
Step 8/13 : FROM buildpack-deps:stretch-scm
 ---> dcbd17451a03
Step 9/13 : LABEL name="httpserver"
 ---> Running in 7f2ba23439a5
Removing intermediate container 7f2ba23439a5
 ---> 4a9849d0113e
Step 10/13 : WORKDIR /opt/httpserver
 ---> Running in a882d264915b
Removing intermediate container a882d264915b
 ---> f1a67ce0b28a
Step 11/13 : COPY httpserver /opt/httpserver/
 ---> 2b5315c97c47
Step 12/13 : EXPOSE 8080
 ---> Running in 936b36df1259
Removing intermediate container 936b36df1259
 ---> 3f2d5d1a644d
Step 13/13 : ENTRYPOINT  ["/opt/httpserver/httpserver"]
 ---> Running in 08f73bea3dca
Removing intermediate container 08f73bea3dca
 ---> e3704d18d9f7
Successfully built e3704d18d9f7
Successfully tagged httpserver:latest
```
运行容器
```shell
 docker run -dit httpserver:latest
```
执行结果
```shell
e56234aa054b4c5e23509bcc9f4973e701377440200dd77d3dd919851017bd1d
```
推送到docker仓库
```shell
docker login
docker tag httpserver:latest sawyer523/httpserver:latest
docker push sawyer523/httpserver:latest

The push refers to repository [docker.io/sawyer523/httpserver]
27a5bd1040d3: Pushed 
bc27313a33a4: Pushed 
e2eb06d8af82: Pushed 
v1.0.0.4: digest: sha256:2bc4fe6669d18fb248172a2c5fdd526fa0b6de4351995676a72b55a0b30c9b5e size: 946
```

进入容器
```shell
03> docker inspect f608a6ae36ef
[
    {
        "Id": "f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072",
        "Created": "2021-10-13T14:31:08.212410715Z",
        "Path": "/opt/httpserver/httpserver",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 178935,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2021-10-13T14:31:08.529901697Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:ecd18bfcca7a8bc87164737cd4f460f5041272ad8f488cbac061840672db6294",
        "ResolvConfPath": "/var/lib/docker/containers/f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072/hostname",
        "HostsPath": "/var/lib/docker/containers/f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072/hosts",
        "LogPath": "/var/lib/docker/containers/f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072/f608a6ae36efdcaf9879c58120aa1dabe4f33e07e310170c438df4bf417f3072-json.log",
        "Name": "/hopeful_hawking",
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
                "LowerDir": "/var/lib/docker/overlay2/b8ae49ebb8bbd09fb0439ed3febdf35d20d0cd5e09dbe33b92727e56c0483f3c-init/diff:/var/lib/docker/overlay2/b3e41b27207968249e3e39607466ec9232fe7ac71f5db0b9c0a81947976c9b60/diff:/var/lib/docker/overlay2/92e374e934845cb52874203f88f584820dba839802babf2dadd7f2e564961db3/diff:/var/lib/docker/overlay2/d14f5f19144cafa7fb179bc3a8bab465fac76891e231984f1a0c24c0f938b1f5/diff:/var/lib/docker/overlay2/4d29a63f9e06190b34b64cae4a338052a83ce81fa38e96e7980d8cdb6e2dc54b/diff:/var/lib/docker/overlay2/567cdcd4a4a631404dc2977307ba2137b1a4bc06f9eaf785d3927873e002513c/diff:/var/lib/docker/overlay2/13837b98278442c98d188f087df1dd9fbe60d007e686eb32d9ed1ba711e30216/diff",
                "MergedDir": "/var/lib/docker/overlay2/b8ae49ebb8bbd09fb0439ed3febdf35d20d0cd5e09dbe33b92727e56c0483f3c/merged",
                "UpperDir": "/var/lib/docker/overlay2/b8ae49ebb8bbd09fb0439ed3febdf35d20d0cd5e09dbe33b92727e56c0483f3c/diff",
                "WorkDir": "/var/lib/docker/overlay2/b8ae49ebb8bbd09fb0439ed3febdf35d20d0cd5e09dbe33b92727e56c0483f3c/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "f608a6ae36ef",
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
            "SandboxID": "dc3a08cbde46419b9dd466be1851e14cdf4af7b181996950871e4c2df9c584c0",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {
                "8080/tcp": null
            },
            "SandboxKey": "/var/run/docker/netns/dc3a08cbde46",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "09a1f20591b016b470d89c32614e2ca667fe9bafa630abbf4a08cc666c785e5f",
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
                    "EndpointID": "09a1f20591b016b470d89c32614e2ca667fe9bafa630abbf4a08cc666c785e5f",
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

root@k8s:~/go/src/geektime-cloud-pratices/03# nsenter --target 178935 --mount --uts --ipc --net --pid
root@f608a6ae36ef:/# ls
bin  boot  dev	etc  home  lib	lib64  media  mnt  opt	proc  root  run  sbin  srv  sys  tmp  usr  var
```