all: clean test gen build

gen:
	protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./api/proto/todo/service.proto

build:
	go build

clean:
	rm -rf ./main ./go-twirp
	rm -rf ./pkg/http/rpc/*/*.pb.go ./pkg/http/rpc/*/*.twirp.go

test:
	echo "todo testing..."
	# go test ./… -v

docker-up:
	echo "todo docker-compose up --build"

docker-down:
	echo "todo docker-compose down"
