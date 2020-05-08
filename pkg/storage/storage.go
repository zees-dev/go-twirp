package storage

import pb "github.com/zees-dev/go-twirp/pkg/http/rpc/todo"

type Repository interface {
	AddTodo(*pb.ToDo) (uint64, error)
	ReadTodo(uint64) (*pb.ToDo, error)
	UpdateTodo(*pb.ToDo) (*pb.ToDo, error)
	DeleteTodo(uint64) (uint64, error)
	ReadAll() ([]*pb.ToDo, error)
}
