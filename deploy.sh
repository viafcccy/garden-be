#!/bin/bash

# code build
echo "-------------------- code building --------------------"
go build -o ./bin/garden-be-exe-file
rm -f ./docker/garden-be-exe-file 
mv ./bin/garden-be-exe-file ./docker/garden-be-exe-file 

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
docker run -d -p 8018:8018 --name garden_be --restart=always garden_be

# display
echo "Checking health..."
curl localhost:8018/health