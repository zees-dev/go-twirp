package main

import (
	"log"
	"net/http"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	"github.com/zees-dev/go-twirp/pkg/storage/memory"
	"github.com/zees-dev/go-twirp/pkg/todo"
)

// Storage repo
const (
	// JSON will store data in JSON files saved on disk
	JSON int = iota
	// Memory will store data in memory
	Memory
)

func main() {
	// set up storage
	storageType := Memory // this could be a flag; hardcoded here for simplicity

	var todoSvc todo.Service

	switch storageType {
	case Memory:
		storage := memory.NewStorage()
		todoSvc = todo.NewService(storage)
	}

	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	twirpToDoHandler := pb.NewToDoServiceServer(todo.NewServer(todoSvc), nil)
	mux.Handle(twirpToDoHandler.PathPrefix(), twirpToDoHandler)

	log.Printf("server listening on port %v...", 8080)
	err := http.ListenAndServe(":8080", mux)
	log.Fatalln(err)
}
