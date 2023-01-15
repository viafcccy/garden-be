#!/bin/bash

# code build
echo "-------------------- code building --------------------"
# build in MacOS for Linux amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/garden-be-exe-file main.go
rm -f ./manifest/docker/garden-be-exe-file 
mv ./bin/garden-be-exe-file ./manifest/docker/garden-be-exe-file 

# prepare
echo "-------------------- removing --------------------"
docker stop garden_be
docker rm garden_be
docker rmi garden_be

# image build
echo "-------------------- image building --------------------"
# enter workdir
cd ./manifest/docker
docker build -t garden_be .

# depoly
echo "-------------------- deploying --------------------"
docker rm -f garden_be
docker run -d -p 8018:8018 --name garden_be --restart=no --link mysql garden_be

# display
echo "Checking health..."
sleep 3
curl localhost:8018/health