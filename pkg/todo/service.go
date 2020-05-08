package todo

import (
	"errors"
	"strings"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
	"github.com/zees-dev/go-twirp/pkg/storage"
)

type Service interface {
	CreateTodo(*pb.ToDo) (uint64, error)
	ReadTodo(uint64) (*pb.ToDo, error)
	UpdateTodo(*pb.ToDo) (*pb.ToDo, error)
	DeleteTodo(uint64) (uint64, error)
	ReadAll() ([]*pb.ToDo, error)
}

type service struct {
	todoR storage.Repository
}

func NewService(r storage.Repository) Service {
	return &service{r}
}

func (s *service) CreateTodo(todo *pb.ToDo) (uint64, error) {
	todoList, err := s.todoR.ReadAll()
	if err != nil {
		return 0, err
	}
	for _, td := range todoList {
		if strings.EqualFold(td.Title, todo.Title) {
			return 0, errors.New("todo with same title already exists in the database")
		}
	}
	id, err := s.todoR.AddTodo(todo)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) ReadTodo(id uint64) (*pb.ToDo, error) {
	return s.todoR.ReadTodo(id)
}

func (s *service) UpdateTodo(todo *pb.ToDo) (*pb.ToDo, error) {
	return s.todoR.UpdateTodo(todo)
}

func (s *service) DeleteTodo(id uint64) (uint64, error) {
	return s.todoR.DeleteTodo(id)
}

func (s *service) ReadAll() ([]*pb.ToDo, error) {
	return s.todoR.ReadAll()
}
