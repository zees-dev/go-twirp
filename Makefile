all: clean gen test build

gen:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./api/proto/todo/service.proto

build:
	go build

clean:
	rm -rf ./main ./go-twirp
	rm -rf ./pkg/http/rpc/*/*.pb.go ./pkg/http/rpc/*/*.twirp.go

test:
	go test ./... -v

bench:
	go test -bench=.

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
