package service

import (
	"context"
	"encoding/json"
	"github.com/abdukhashimov/go_api_mono_repo/generated/todo"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TodoService struct {
	todo.UnimplementedTodoServiceServer
	db repository.Store
}

//func NewTodoService(db repository.Store) *TodoService {
//	return &TodoService{
//		db: db,
//	}
//}

func (s *TodoService) CreateTodo(cts context.Context, req *todo.Todo) (*todo.Todo, error) {
	err := s.db.CreateTodo(cts, sqlc.CreateTodoParams{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}

	return &todo.Todo{
		Title:       req.Title,
		Description: req.Description,
	}, nil
}

func (s *TodoService) GetTodoById(cts context.Context, req *todo.Todo) (*todo.Todo, error) {
	to_do, err := s.db.GetTodoById(cts, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get todo: %s", err.Error())
	}
	response := todo.Todo{}
	if todoInBytes, err := json.Marshal(to_do); err == nil {
		if err := json.Unmarshal(todoInBytes, &response); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal todo: %s", err.Error())
		}
	} else {
		return nil, status.Errorf(codes.Internal, "failed to marshal todo: %s", err.Error())
	}
	return &response, nil
}

func (s *TodoService) ListTodos(cts context.Context, req *todo.Todo) (*todo.Todo, error) {
	todos, err := s.db.ListTodos(cts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list todos: %s", err.Error())
	}
	response := todo.Todo{}
	if todosInBytes, err := json.Marshal(todos); err == nil {
		if err := json.Unmarshal(todosInBytes, &response); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to unmarshal todos: %s", err.Error())
		}
	} else {
		return nil, status.Errorf(codes.Internal, "failed to marshal todos: %s", err.Error())
	}
	return &response, nil
}

func (s *TodoService) UpdateTodo(cts context.Context, req *todo.Todo) (*todo.Todo, error) {
	err := s.db.UpdateTodo(cts, sqlc.UpdateTodoParams{
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update todo: %s", err.Error())
	}
	return req, nil

}
