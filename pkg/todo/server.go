package todo

import (
	"context"
	"log"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
)

type server struct {
	svc Service
}

func NewServer(s Service) *server {
	return &server{s}
}

func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	log.Println(req.ToDo)
	id, err := s.svc.CreateTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (s *server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	todo, err := s.svc.ReadTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{ToDo: todo}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	todo, err := s.svc.UpdateTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Updated: todo.Id}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := s.svc.DeleteTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Deleted: id}, nil
}

func (s *server) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	todoList, err := s.svc.ReadAll()
	if err != nil {
		return nil, err
	}
	return &pb.ReadAllResponse{ToDos: todoList}, nil
}
