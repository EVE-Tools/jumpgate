#!/bin/bash
cd "$(dirname "$0")"

services=(esiMarkets marketStats staticData topStations)

go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/golang/protobuf/protoc-gen-go

for service in ${services[@]}; do
  mkdir -p ./element43/services/$service

  protoc -I/usr/local/include -I../element43/services/$service \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:./element43/services/$service \
  --grpc-gateway_out=logtostderr=true:./element43/services/$service \
  --swagger_out=logtostderr=true:./element43/services/$service \
  ../element43/services/$service/$service.proto
done