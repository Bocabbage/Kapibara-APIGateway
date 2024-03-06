# Kapibara-APIGateway

## Overview
APIGateway service for self-service-system Kapibara.

## Build & Deploy
### 1. Modes
- [x] Local demo
- [ ] Kubernetes deploy configurations
### 2. How to start
#### Local-demo: docker-compose
```shell
bash build/localdemo.sh
```
#### gRPC protos
```shell
cd protobufs && \
protoc --go_out=../internal/grpc_utils \
--go_opt=paths=source_relative \
--go-grpc_out=../internal/grpc_utils \
--go-grpc_opt=paths=source_relative \
protobufs/mongodb_crud.proto
```

## Todo
### APIs
- [x] Auth-APIs
- Register
- Login
- [ ] WebCollector-related-APIs
### Features
- [ ] Swagger Documents
- [x] HTTPs support
- [ ] Traffic control