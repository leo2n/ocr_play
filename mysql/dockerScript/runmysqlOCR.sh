#!/bin/bash
# mkdir -p $HOME/docker/mysql/conf.d $HOME/docker/mysql/data
docker stop ocrmysql && docker rm ocrmysql
docker rmi ocrmysql:test
docker build -t ocrmysql:test .
docker network create ocr # 创建这个应用的专属 network bridge ocr

docker run -d --name ocrmysql -p 127.0.0.1:3310:3306 --restart=always -v $HOME/teletraan/mysql/conf.d:/etc/mysql/conf.d -v $HOME/teletraan/mysql/data:/var/lib/mysql -v $PWD/initScripts:/docker-entrypoint-initdb.d --network=ocr --net-alias=ocrmysql leo2n/ocrmysql:test
# 在ocr网络中, 找到数据库的对应IP地址, 填写到应用的配置文件中
