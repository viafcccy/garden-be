FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /garden-be

# todo: 编译过程移到容器内
COPY ./garden-be-exe-file ./garden-be-exe-file
COPY ./config/config_pro.yaml ./manifest/config/config_pro.yaml

EXPOSE 8018

ENTRYPOINT [ "/bin/sh", "-c", "./garden-be-exe-file" ]
