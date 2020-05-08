# Go Twirp

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

- todo

## TODO

- [x] Create/generate todo struct(s)
- [x] Create DB (in-memory) to store db (this is mock)
- [x] Complete REST handlers
- [x] Implement interfaces and integrate with DB
- [x] Readme completion
- [ ] Docker
- [ ] Testing
- [x] Makefile
- [ ] Github actions (CICD)
- [ ] Goreport reference (badge)
- [ ] Godoc reference (badge)
- [ ] Sourcegraph reference (badge)
- [ ] Build reference (badge)

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)