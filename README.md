# Go Twirp

[![Go Report Card](https://goreportcard.com/badge/github.com/zees-dev/go-twirp)](https://goreportcard.com/report/github.com/zees-dev/go-twirp)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/zees-dev/go-twirp)
[![Sourcegraph](https://sourcegraph.com/github.com/zees-dev/go-twirp/-/badge.svg)](https://sourcegraph.com/github.com/zees-dev/go-twirp?badge)

A lightweight rpc server that exposes HTTP 1.1 REST endpoints via use of [Twirp](https://twitchtv.github.io/twirp/).

**What is Twirp:**

[Twirp](https://twitchtv.github.io/twirp/), an rpc based framework developed by [twitch](https://www.twitch.tv/)

Think of it as taking the best parts of [gRPC](https://grpc.io/) and combining it with the simplicity of traditional REST APIs to give you best of both worlds - a high performance communication mechanism (RPC) that also supports REST clients.

For additional details, refer to the twitch blog post [here](https://twitchtv.github.io/twirp/)

**Why Twirp?**

- When you want high performance communication between microservices (internal infrastructure) and also want to use the same models for external clients (web based)
  - Single proto files (models) can be used as source of truth for all internal and external services!
- When [gRPC](https://grpc.io/) is too complex to setup/use for new developers and you need to hit the ground running - or those used to creating REST API's
- When you want to standardize your API endpoints - since theres a thousand ways to create REST API's (even though we have [OpenAPI](https://swagger.io/specification/))
- REST APIs are not an efficient/optimal means of communication between your internal microservices

**NOTE:** [gRPC-web](https://github.com/grpc/grpc-web) can be used to achieve similar results; however it is non-trivial to setup (compared to Twirp) and requires a [Envoy (gateway proxy)](https://www.envoyproxy.io/) to translate web-based REST calls to gRPC.

## Features

- RPC based web server that supports HTTP 1.1
- Hexagonal architecture
  - Loose coupling (everything is modular & testable)
- Support for both rest & protobuf based clients
- Docker support

## Development

### Pre-requisites

- [protoc compiler](https://github.com/protocolbuffers/protobuf/releases) - pre-compiled binary for your OS
- [protoc-gen-go](https://github.com/golang/protobuf/tree/master/protoc-gen-go) - protoc plugin for go
  - ```go get github.com/golang/protobuf/protoc-gen-go```
- [protoc-gen-twirp](https://github.com/twitchtv/twirp/tree/master/protoc-gen-twirp) - protoc plugin for twirp
  - ```go get -u github.com/twitchtv/twirp/protoc-gen-twirp```

- Note: All the dependencies must be available in your PATH

### Generate code from protobuf compiler

```sh
protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./api/proto/todo/service.proto
```

### Run Twirp based web-server

```sh
go run main.go
```

### Clients

Clients can also be generated from proto files.\
The magic of Twirp however is that it also supports HTTP 1.1 - hence allowing you to call RPC endpoints via curl (in addition to protobuf)

### cURL

#### GET

```sh
curl -i \
  --request "POST" \
  --location "http://localhost:8080/twirp/ToDoService/Read" \
  --header "Content-Type:application/json" \
  --data '{"id": 1}'
```

#### POST

```sh
curl -i \
  --request "POST" \
  --location "http://localhost:8080/twirp/ToDoService/Create" \
  --header "Content-Type:application/json" \
  --data '{"toDo":{"title":"Second","description":"inital item in todo list"}}'
```

#### UPDATE

```sh
curl -i \
  --request "POST" \
  --location "http://localhost:8080/twirp/ToDoService/Update" \
  --header "Content-Type:application/json" \
  --data '{"toDo":{"id": "1", "title":"Second","description":"inital item in todo list"}}'
```

#### DELETE

```sh
curl -i \
  --request "POST" \
  --location "http://localhost:8080/twirp/ToDoService/Delete" \
  --header "Content-Type:application/json" \
  --data '{"id": 1}'
```

#### GET (get all)

```sh
curl -i \
  --request "POST" \
  --location "http://localhost:8080/twirp/ToDoService/ReadAll" \
  --header "Content-Type:application/json" \
  --data '{}'
```

## Testing

```sh
go test ./... -v
```

## Benchmarking

A benchmark has been provided which compares REST client performance against Protobuf client using provided [benchmark file](./benchmark_test.go).\
In this benchmark an in-memory HTTP server is initialized which serves the twirp API; we benchmark client REST and Protobuf requests against this API.

According to the benchmarks, the protobuf requests appear to be 10-15% faster than REST based client requests against the API.

```sh
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/zees-dev/go-twirp
BenchmarkRestClient-8              13839             83063 ns/op
BenchmarkRPCClient-8               17268             69278 ns/op
PASS
ok      github.com/zees-dev/go-twirp    4.124s
```

### Run benchmark

```sh
go test -bench=.
```

## TODO

- [x] Create/generate todo struct(s)
- [x] Create DB (in-memory) to store db (this is mock)
- [x] Complete REST handlers
- [x] Implement interfaces and integrate with DB
- [x] Readme completion
- [x] Docker
- [x] Testing
- [x] Perform benchmarks to compare REST calls and native RPC calls on API
- [x] Makefile
- [ ] Github actions (CICD)
- [x] Goreport reference (badge)
- [ ] Complete documentation for public code (for godoc reference)
- [x] Godoc reference (badge)
- [x] Sourcegraph reference (badge)
- [ ] Build reference (badge)
- [ ] Code coverage reference (badge)

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
