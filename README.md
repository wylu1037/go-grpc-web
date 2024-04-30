<h1 align="center">Go gRPC</h1>

## About
The go web build with gRPC.

## Feature
- [x] Logging.
- [x] API documentation with Swagger.
- [x] Generate go code use buf.
- [x] Support jwt authentication.
- [x] Secure transport use TLS.
- [x] Unit testing with testify
- [x] CI/CD with GitHub action

## Tech stack

- [Go](https://go.dev) Modern programming language
- [Docker](https://www.docker.com/) Containerization
- [GRPC](https://grpc.io/) A high performance, open source universal RPC framework Get started!
- [Fx](https://github.com/uber-go/fx) A dependency injection
- [Sqlite3](https://github.com/mattn/go-sqlite3) Embedded database
- [SqlC](https://sqlc.dev/) Compile SQL to type-safe code; catch failures before they happen.

## Environment
Before you start, make sure you have Git, Go, Docker, gRPC, sqlC, and Sqlite3 installed.

<h3>Buf</h5>

``` shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

<h3>gRPC-Gateway</h3>

``` shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

## Getting start

<h3>Generate proto</h3>

``` shell
buf generate proto
```

<h3>Generate database</h3>

```shell
sqlc generate
```
