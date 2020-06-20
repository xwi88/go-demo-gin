# build image
# image include: bash, git, golang, make, openssh, tzdata
# image ENV TZ=Asia/Shanghai
FROM v8fg/golang:1.14.4-upx AS builder

#GO Module ENV
ENV BUILD_TAGS="jsoniter"  LDFLAGS='-s -w' SRC_PATH=/go/src/github.com/xwi88/go-demo-gin
ENV GOPROXY=https://goproxy.cn,https://goproxy.io/,https://mirrors.aliyun.com/goproxy/,https://gocenter.io/,https://proxy.golang.org,direct

WORKDIR ${SRC_PATH}
COPY . ${SRC_PATH}

RUN make release-mod && upx ${SRC_PATH}/dist/app


#Local Run use alpine:3.12.1
# image include: tzdata
# image ENV TZ=Asia/Shanghai
FROM v8fg/alpine:3.12.1
LABEL maintainer=xwi88

ENV SRC_PATH=/go/src/github.com/xwi88/go-demo-gin APP_PATH=/data1/services/app

WORKDIR ${APP_PATH}
COPY --from=builder ${SRC_PATH}/dist .
RUN chmod +x app

EXPOSE 9990

ENTRYPOINT ["./entrypoint.sh"]
CMD ["app:start"]
