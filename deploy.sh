#!/bin/bash

# code build
echo "-------------------- code building --------------------"
go build -v -n -o ./docker/garden-be-exe-file

# prepare
echo "-------------------- removing --------------------"
docker stop garden_be
docker rm garden_be
docker rmi garden_be

# image build
echo "-------------------- image building --------------------"
# enter workdir
cd ./docker
docker build -t garden_be .

# depoly
echo "-------------------- deploying --------------------"
docker rm -f garden_be
docker-compose up

# display
echo "Checking health..."
curl 127.0.0.1:8018/health