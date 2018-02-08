# jumpgate
[![Build Status](https://drone.element-43.com/api/badges/EVE-Tools/jumpgate/status.svg)](https://drone.element-43.com/EVE-Tools/jumpgate) [![Go Report Card](https://goreportcard.com/badge/github.com/EVE-Tools/jumpgate)](https://goreportcard.com/report/github.com/EVE-Tools/jumpgate) [![Docker Image](https://images.microbadger.com/badges/image/evetools/jumpgate.svg)](https://microbadger.com/images/evetools/jumpgate)

This service for [Element43](https://element-43.com) is based on [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and translates requests between HTTP and gRPC.

Issues can be filed [here](https://github.com/EVE-Tools/element43). Pull requests can be made in this repo.

## Installation
Either use the prebuilt Docker images and pass the appropriate env vars (see below), or:

* Install Go, clone this repo into your gopath
* Run `go get ./...` to fetch the service's dependencies
* Run `bash generateProto.sh` to generate the necessary gRPC-related code
* Run `go build` to build the service
* Run `./jumpgate` to start the service

## Deployment Info
Builds and releases are handled by Drone.

Environment Variable | Default | Description
--- | --- | ---
LOG_LEVEL | info | The service's log level
LISTEN_HOST | :8000 | The host/port jumpgate will listen on
ESI_MARKETS_HOST | esi-markets.element43.cluster.local:43000 | Location of the ESI markets service
MARKET_STATS_HOST | market-stats.element43.cluster.local:43000 | Location of the market statistics service
STATIC_DATA_HOST | static-data.element43.cluster.local:43000 | Location of the static data service
TOP_STATIONS_HOST | top-stations.element43.cluster.local:43000 | Location of the top stations service