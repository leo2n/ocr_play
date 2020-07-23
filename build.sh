#!/usr/bin/env bash
go build -o teletraanBin *.go
docker stop ocrapp && docker rm ocrapp;
docker rmi leo2n/ocrapp:test;
docker build -t leo2n/ocrapp:test . ;
docker run -d --name ocrapp -p 4001:4001 -v $HOME/teletraan/imageStore:/usr/local/teletraan/imageStore --network=ocr leo2n/ocrapp:test ;
