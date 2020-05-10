package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/zees-dev/go-twirp/pkg/http/rest"
	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	"github.com/zees-dev/go-twirp/pkg/storage/memory"
	"github.com/zees-dev/go-twirp/pkg/todo"
)

func getHTTPHandler() http.Handler {
	// Initialize storage with one item
	storage := memory.NewStorage()
	todoItem := &pb.ToDo{Title: "Initial Item", Description: "my first task"}
	storage.AddTodo(todoItem)

	todoSvc := todo.NewService(storage)
	handler := rest.NewTwirpMux(todoSvc)
	mux := handler.Routes()
	return mux
}

func BenchmarkRestClient(b *testing.B) {
	// Start HTTP server
	handler := getHTTPHandler()
	go func() { http.ListenAndServe(":8080", handler) }()

	var dataStr = []byte(`{"id": 1}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ := http.Post("http://localhost:8080/twirp/ToDoService/Read", "application/json", bytes.NewBuffer(dataStr))
		body, _ := ioutil.ReadAll(res.Body)
		var readResponse pb.ReadResponse
		json.Unmarshal(body, &readResponse)
		if readResponse.ToDo.Title != "Initial Item" {
			fmt.Println("incorrect proto response")
		}
	}
}

func BenchmarkRPCClient(b *testing.B) {
	// Start HTTP server
	handler := getHTTPHandler()
	go func() { http.ListenAndServe(":8080", handler) }()

	client := pb.NewToDoServiceProtobufClient("http://localhost:8080", &http.Client{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readResponse, _ := client.Read(context.Background(), &pb.ReadRequest{Id: 1})
		if readResponse.ToDo.Title != "Initial Item" {
			fmt.Println("incorrect proto response")
		}
	}
}
