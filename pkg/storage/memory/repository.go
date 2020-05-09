package memory

import (
	"errors"
	"sync"

	pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"
)

type storage struct {
	todos       []*pb.ToDo
	todoCounter uint64
	mux         sync.RWMutex
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) AddTodo(todo *pb.ToDo) (uint64, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.todoCounter++
	s.todos = append(s.todos, &pb.ToDo{Id: s.todoCounter, Title: todo.Title, Description: todo.Description, Reminder: todo.Reminder})
	return s.todoCounter, nil
}

func (s *storage) ReadTodo(id uint64) (*pb.ToDo, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	for _, td := range s.todos {
		if td.Id == id {
			return td, nil
		}
	}
	return nil, errors.New("Todo Id does not exist")
}

func (s *storage) UpdateTodo(todo *pb.ToDo) (*pb.ToDo, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	index := -1
	for i, td := range s.todos {
		if td.Id == todo.Id {
			index = i
			break
		}
	}

	if index > -1 {
		todo.Id = s.todos[index].Id
		s.todos[index] = todo
		return todo, nil
	}

	return nil, errors.New("unable to find todo item with specified id")
}

func (s *storage) DeleteTodo(id uint64) (uint64, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	deleteIndex := -1
	for i, td := range s.todos {
		if td.Id == id {
			deleteIndex = i
			break
		}
	}
	if deleteIndex > -1 {
		deletedTodo := s.todos[deleteIndex]
		s.todos = append(s.todos[:deleteIndex], s.todos[deleteIndex+1:]...)
		return deletedTodo.Id, nil
	}

	return 0, errors.New("Todo item not found")
}

func (s *storage) ReadAll() ([]*pb.ToDo, error) {
	return s.todos, nil
}
