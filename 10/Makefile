release := $(shell pwd)
arch := $(shell go env GOARCH)

major := $(shell cat $(release)/version.go | grep "MajorVersion" | grep -oh "\d\+")
middle :=  $(shell cat $(release)/version.go | grep "MiddleVersion" | grep -oh "\d\+")
minor :=  $(shell cat $(release)/version.go | grep "MinorVersion" | grep -oh "\d\+")
build :=  $(shell cat $(release)/version.go | grep "BuildNo" | grep -oh "\d\+")
tmp := $(shell expr $(build) + 1)

.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build
.PHONY: mac
mac: build.version ## build for mac
	@make build.version tmp=$(tmp)
	@go build -o httpserver main.go

.PHONY:	linux
linux: build.version  ## build for linux
	@GOOS=linux GOARCH=$(arch) go build -o httpserver main.go

.PHONY: docker
docker:  ## build for docker
	@docker build -t httpserver:v$(major).$(middle).$(minor).$(tmp) .
	@docker tag httpserver:v$(major).$(middle).$(minor).$(tmp) sawyer523/httpserver:v$(major).$(middle).$(minor).$(tmp)
	@docker push sawyer523/httpserver:v$(major).$(middle).$(minor).$(tmp)


build.version:
	@# shellcheck disable=SC2084
	@sed -in-place -e "s/"BuildNo\ =\ $(build)"/"BuildNo\ =\ $(tmp)"/g" $(release)/version.go
	@rm $(release)/version.gon-place

