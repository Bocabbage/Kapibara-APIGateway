# Kapibara-APIGateway

## Overview
APIGateway service for self-service-system Kapibara.

## Build & Deploy
### 1. Modes
- Local demo
- Kubernetes deploy configurations
### 2. How to start
#### Local-demo: docker-compose
```shell
bash build/localdemo.sh
```
#### gRPC protos
```shell
# GTPC_GATEWAY_THIRDPARTY_PATH="${HOME}/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis/"
cd protobufs/ && protoc -I ${GTPC_GATEWAY_THIRDPARTY_PATH} -I . \
  --go_out ../internal/grpc/mikanani \
  --go_opt paths=source_relative \
  --go-grpc_out ../internal/grpc/mikanani \
  --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ../internal/grpc/mikanani \
  --grpc-gateway_opt paths=source_relative \
  mikanani_grpc.proto
```

## Todo
- [x] Use `grpc-gateway` to replace hard-code http->grpc (mikanani)
- [x] Use `GORM` to replace direct mysql access
- [x] Extract code for Redis cache
- [ ] Timeout setting
- [ ] Traffic management
- [ ] Params Validator
- [ ] Configuration injected with yaml
- [ ] API Document