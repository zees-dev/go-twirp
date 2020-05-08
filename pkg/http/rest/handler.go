package rest

import (
	"net/http"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	"github.com/zees-dev/go-twirp/pkg/todo"
)

type twirpMux struct {
	s *todo.Server
}

func NewTwirpMux(s *todo.Server) *twirpMux {
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
