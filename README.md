# go-demo-gin

go demo with gin, support docker and docker-compose

## Env

- go1.14.4+
- platform: Linux|Unix|Mac OS

>GO Proxy 配置

```bash
export GOPROXY=https://goproxy.cn,https://goproxy.io/,https://mirrors.aliyun.com/goproxy/,https://gocenter.io/,https://proxy.golang.org,direct

# 私有仓库，公司仓库
export GOPRIVATE=*.example.com
```
## Compile & Run

- `make run` or `go build`

>具体命令请参考: `Makefile` 文件

## Docker build & start

- `make docker`        创建镜像
- `make docker-start`  镜像启动
- `make docker-version` 二进制版本 version
- `bash docker/docker-startup.sh` 镜像启动
- `docker-compose up` docker-compose 方式启动

## Base docker images

- go源码编译镜像 [v8fg/golang:1.14.4-upx](https://hub.docker.com/repository/docker/v8fg/golang)
- go二进制运行镜像 [v8fg/alpine](https://hub.docker.com/repository/docker/v8fg/alpine)

### Tip

- *docker-startup.sh* 启动脚本中输入参数，务必使用 $PARAMS 而非 "$PARAMS"!
- $PARAMS 为空，则执行默认命令; 输入 "$PARAMS" 可能会被当做 " " 参数，会造成未知影响!
- 使用 `upx` 压缩二进制可执行程序，减小 `image` 或可执行程序大小(30%-50%)，`brew install upx` 安装，执行 `make upx`
- 测试服务器测试，借助vendor目录加速构建 `make docker-mod`

## Deploy

- **创建镜像** `make docker-build-dev`|`make docker-build-uat`|`make docker-build-prod`
- **镜像推送** 推送生成的镜像到 *corp* 镜像仓库
- **镜像使用** 选择要发布的镜像版本，发布即可，默认启动
- **开发测试部署** `make deploy` 借助 `docker`, `docker-compose`
