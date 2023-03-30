package api

import (
	"context"
	"todo-list-cqrs-es/internal/domain/user"
)

type API interface {
	GetUser(ctx context.Context, id string) (*user.User, error)

	CreateTodoList(ctx context.Context, list CreateTodoListRequest) (*TodoListResponse, error)
	GetTodoList(ctx context.Context, id string) (*TodoListResponse, error)
	DeleteTodoList(ctx context.Context, id string) error
}
