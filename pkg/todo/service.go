package todo

import (
	"context"
	"errors"
	"strings"

	"github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	"github.com/zees-dev/go-twirp/pkg/storage"
)

type service struct {
	todoR storage.Repository
}

func NewService(r storage.Repository) todo.ToDoService {
	return &service{r}
}

func (s *service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	todoList, err := s.todoR.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, td := range todoList {
		if strings.EqualFold(td.Title, req.ToDo.Title) {
			return nil, errors.New("todo with same title already exists in the database")
		}
	}
	id, err := s.todoR.AddTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Id: id}, nil
}

func (s *service) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	todo, err := s.todoR.ReadTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.ReadResponse{ToDo: todo}, nil
}

func (s *service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	todo, err := s.todoR.UpdateTodo(req.ToDo)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{Updated: todo.Id}, nil
}

func (s *service) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id, err := s.todoR.DeleteTodo(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Deleted: id}, nil
}

func (s *service) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	todoList, err := s.todoR.ReadAll()
	if err != nil {
		return nil, err
	}
	return &pb.ReadAllResponse{ToDos: todoList}, nil
}
