package rest

import (
	"net/http"

	"github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
)

type twirpMux struct {
	s todo.ToDoService
}

func NewTwirpMux(s todo.ToDoService) *twirpMux {
	return &twirpMux{s}
}

func (tm *twirpMux) Routes() http.Handler {
	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	twirpToDoHandler := pb.NewToDoServiceServer(tm.s, nil)
	mux.Handle(twirpToDoHandler.PathPrefix(), twirpToDoHandler)
	return mux
}
