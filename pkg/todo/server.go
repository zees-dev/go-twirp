package todo

import (
	"context"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
)

type Server struct {
	svc Service
}

func NewServer(s Service) *Server {
	return &Server{s}
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := s.svc.CreateTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (s *Server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	todo, err := s.svc.ReadTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{ToDo: todo}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	todo, err := s.svc.UpdateTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Updated: todo.Id}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := s.svc.DeleteTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Deleted: id}, nil
}

func (s *Server) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	todoList, err := s.svc.ReadAll()
	if err != nil {
		return nil, err
	}
	return &pb.ReadAllResponse{ToDos: todoList}, nil
}
