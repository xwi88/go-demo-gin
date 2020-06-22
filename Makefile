# Makefile to build the command lines and tests in this project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.
SHELL := /bin/sh

existBash = $(shell cat /etc/shells|grep -w /bin/bash|grep -v grep)
ifneq (, $(strip ${existBash}))
	SHELL = /bin/bash
endif
$(info shell will use ${SHELL})

#BASEDIR = $(shell pwd)
BASEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# add following lines before go build!
versionDir = github.com/xwi88/version

gitBranch = $(shell git symbolic-ref --short -q HEAD)

ifeq ($(gitBranch),)
gitTag = $(shell git describe --always --tags --abbrev=0)
endif

buildTime = $(shell date "+%FT%T%z")
gitCommit = $(shell git rev-parse HEAD)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

# -ldflags flags accept a space-separated list of arguments to pass to an underlying tool during the build.
ldFlagsDebug="-X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
 -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w
#ldFlagsRelease="-s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
#  -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
#  -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w
# -a #force rebuilding of packages that are already up-to-date.
ldFlagsRelease="-installsuffix -s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
  -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
  -X ${versionDir}.gitTreeState=${gitTreeState}"

$(shell mkdir -p ${BASEDIR}/build/bin/conf)

#buildTags=""
buildTags="jsoniter"

export GOPROXY=https://goproxy.cn,https://goproxy.io/,https://mirrors.aliyun.com/goproxy/,https://gocenter.io/,https://proxy.golang.org,direct
export GOPRIVATE=*.xwi88.com

.PHONY: default run re-run app app-darwin app-linux docker docker-start docker-version \
	release-mod release-mod upx proto

default: app

all: app

clean:
	rm -r build/bin

app:
	go build -v -tags ${buildTags} -ldflags ${ldFlagsDebug} -o ${BASEDIR}/build/bin/app  ${BASEDIR}
	@echo "Done app built remain gdb info"

app-darwin:
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/build/bin/app-darwin ${BASEDIR}
	@echo "Done app built for darwin, remain gdb info "

app-linux:
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/build/bin/app-linux ${BASEDIR}
	@echo "Done app built for linux"

release:
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/dist/app ${BASEDIR}
	@echo "Done app release built"

release-mod:
	go build -v -mod=vendor -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/dist/app ${BASEDIR}
	@echo "Done app release built"

docker: docker-build-dev
docker-build-dev:
	bash ${BASEDIR}/docker/docker-build.sh dev
	@echo "Done docker built for dev"
docker-build-uat:
	bash ${BASEDIR}/docker/docker-build.sh uat
	@echo "Done docker built for uat"
docker-build-prod:
	bash ${BASEDIR}/docker/docker-build.sh prod
	@echo "Done docker built for prod"

docker-start:
	bash ${BASEDIR}/docker/docker-startup.sh
	@echo "Done docker start with default params"

docker-version:
	bash ${BASEDIR}/docker/docker-startup.sh version
	@echo "Done docker version"

docker-mod:
	go mod tidy && go mod vendor && \
	bash ${BASEDIR}/docker/docker-build.sh mod
	@echo "Done docker with go mod built for dev"

deploy:
	bash ${BASEDIR}/quick_startup.sh
	@echo "Done docker with go mod deploy only for dev test"

run:
	${BASEDIR}/build/bin/app start

re-run: clean app run

version:
	${BASEDIR}/build/bin/app version

upx: app-darwin app-linux
	upx ${BASEDIR}/build/bin/app-darwin
	upx ${BASEDIR}/build/bin/app-linux
	ls -lhr ${BASEDIR}/build/bin/*

proto:
	protoc -I proto/ --go_out=pb bus_log.proto
	@echo "Done proto built for golang"
