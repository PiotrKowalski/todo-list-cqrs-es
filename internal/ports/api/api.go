package api

import (
	"context"
)

type API interface {
	CreateTodoList(ctx context.Context, list CreateTodoListRequest) (*TodoListResponse, error)
	GetTodoList(ctx context.Context, id string) (*TodoListResponse, error)
	DeleteTodoList(ctx context.Context, id string) error
}
