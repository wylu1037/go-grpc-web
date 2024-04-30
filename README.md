<h1 align="center">Go gRPC</h1>

## About


## Feature

## Tech stack

- [Go](https://go.dev) modern programming language
- [Docker](https://www.docker.com/) containerization
- [GRPC]()
- [Fx](https://github.com/uber-go/fx)
- [Sqlite3]() embedded database
- [SqlC]()

## Environment

## Getting start
## buf
buf.gen.yaml
```yaml
version: v1
managed:
  enabled: true
  go_package_prefix:
    default: lattice-manager-grpc/gen
plugins:
  - plugin: go
    out: gen
    opt:
      - paths=source_relative
  - plugin: go-grpc
    out: gen
    opt:
      - paths=source_relative
```

用到了插件：
+ go: go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
+ go-grpc: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```shell
buf generate proto
```

gRPC-gateway
+ github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
+ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

