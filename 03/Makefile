release := $(shell pwd)
arch := $(shell go env GOARCH)

major := $(shell cat $(release)/version.go | grep "MajorVersion" | grep -oh "\d\+")
middle :=  $(shell cat $(release)/version.go | grep "MiddleVersion" | grep -oh "\d\+")
minor :=  $(shell cat $(release)/version.go | grep "MinorVersion" | grep -oh "\d\+")
build :=  $(shell cat $(release)/version.go | grep "BuildNo" | grep -oh "\d\+")
tmp := $(shell expr $(build) + 1)

.PHONY: build.mac
build.mac:
	@make build.version tmp=$(tmp)
	@go build -o httpserver main.go

.PHONY:	build.linux
build.linux:
	@make build.version tmp=$(tmp)
	@GOOS=linux GOARCH=$(arch) go build -o httpserver main.go

.PHONY: build.docker
build.docker:
	@make build.linux tmp=$(tmp)
	@docker build -t httpserver:v$(major).$(middle).$(minor).$(tmp) .

.PHONY: build.version
build.version:
	# shellcheck disable=SC2084
    $(shell sed -in-place -e "s/"BuildNo\ =\ $(build)"/"BuildNo\ =\ $(tmp)"/g" $(release)/version.go)
    $(shell rm $(release)/version.gon-place)