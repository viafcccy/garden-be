FROM golang:1.18

WORKDIR /app

# todo: 编译过程移到容器内
COPY ./garden-be-exe-file /app/garden-be-exe-file

ENTRYPOINT [ "/bin/sh", "-c", "/app/garden-be-exe-file" ]
