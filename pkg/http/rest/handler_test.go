package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
	handler := NewTwirpMux(todoSvc)
	mux := handler.Routes()
	return mux
}

func TestGetTodo(t *testing.T) {
	handler := getHTTPHandler()

	var dataStr = []byte(`{"id": 1}`)
	req, err := http.NewRequest("POST", "/twirp/ToDoService/Read", bytes.NewBuffer(dataStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"toDo":{"id":"1","title":"Initial Item","description":"my first task"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostTodo(t *testing.T) {
	handler := getHTTPHandler()

	var dataStr = []byte(`{"toDo":{"title":"Second","description":"inital item in todo list"}}`)

	req, err := http.NewRequest("POST", "/twirp/ToDoService/Create", bytes.NewBuffer(dataStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"2"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMultiplePostTodo(t *testing.T) {
	// table-driven tests (with output): https://gobyexample.com/testing

	handler := getHTTPHandler()

	cases := []struct {
		dataStr, responseStr string
	}{
		{`{"toDo":{"title":"Second","description":"inital item in todo list"}}`, `{"id":"2"}`},
		{`{"toDo":{"title":"third","description":"third item"}}`, `{"id":"3"}`},
		{`{"toDo":{"title":"4th item","description":"fourth"}}`, `{"id":"4"}`},
	}

	for i, c := range cases {
		testname := fmt.Sprintf("%d/%d", i+1, len(cases))
		t.Run(testname, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/twirp/ToDoService/Create", bytes.NewBuffer([]byte(c.dataStr)))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			// Check response status code
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body is what we expect.
			expected := c.responseStr
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	handler := getHTTPHandler()

	var dataStr = []byte(`{"toDo":{"id": "1", "title":"First","description":"changed item in todo list"}}`)

	req, err := http.NewRequest("POST", "/twirp/ToDoService/Update", bytes.NewBuffer(dataStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"updated":"1"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteTodo(t *testing.T) {
	handler := getHTTPHandler()

	var dataStr = []byte(`{"id": 1}`)

	req, err := http.NewRequest("POST", "/twirp/ToDoService/Delete", bytes.NewBuffer(dataStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"deleted":"1"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllTodos(t *testing.T) {
	handler := getHTTPHandler()

	var dataStr = []byte(`{}`)

	req, err := http.NewRequest("POST", "/twirp/ToDoService/ReadAll", bytes.NewBuffer(dataStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"toDos":[{"id":"1","title":"Initial Item","description":"my first task"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
