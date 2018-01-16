#!/bin/bash
sudo rm -rf .docker/plugin/rootfs
mkdir .docker/plugin/rootfs

GOOS=linux GOARCH=amd64 go build
cp azurefile-dockervolumedriver .docker/plugin/

docker build -t rootfsimage .docker/plugin
id=`docker create rootfsimage true`

docker export $id |sudo tar -x -C .docker/plugin/rootfs

docker plugin rm jmaitrehenry/azurefile
sudo docker plugin create jmaitrehenry/azurefile .docker/plugin
docker plugin push jmaitrehenry/azurefile
