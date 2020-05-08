package main

import (
	"log"
	"net/http"

	"github.com/zees-dev/go-twirp/pkg/http/rest"
	"github.com/zees-dev/go-twirp/pkg/storage"
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
	storageType := Memory // this could be a flag; hardcoded here for simplicity

	var storage storage.Repository

	switch storageType {
	case Memory:
		storage = memory.NewStorage()
	}

	todoSvc := todo.NewService(storage)
	todoServer := todo.NewServer(todoSvc)
	handler := rest.NewTwirpMux(todoServer)
	mux := handler.Routes()

	log.Printf("server listening on port %v...", 8080)
	err := http.ListenAndServe(":8080", mux)
	log.Fatalln(err)
}
