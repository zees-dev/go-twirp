# BUILD
FROM golang:1.14 as build
LABEL maintainer="github.com/zees-dev"

# Install proto compiler, go & twirp protoc plugins
RUN apt update && \
  apt install -y protobuf-compiler && \
  go get github.com/golang/protobuf/protoc-gen-go && \
  go get -u github.com/twitchtv/twirp/protoc-gen-twirp

WORKDIR /go/src/app

COPY . .
RUN go mod download

# generate & build
RUN protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./api/proto/todo/service.proto
RUN go build -o /go/bin/twirp

# MERGE
FROM gcr.io/distroless/base
COPY --from=build /go/bin/twirp /bin/
EXPOSE 8080
ENTRYPOINT [ "/bin/twirp" ]