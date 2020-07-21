#!/usr/bin/env bash
go build -o teletraanBin *.go
docker build -t leo2n/ocrapp:0.91 .
docker run -d --name ocrapp -p 4001:4001 -v $PWD/imageStore:/usr/local/teletraan/imageStore leo2n/ocrapp:0.91