package user

import (
	"context"
	dUser "todo-list-cqrs-es/internal/domain/user"
)

type Get struct {
	Id string
}

type GetUserHandler interface {
	Handle(ctx context.Context, query Get) (*dUser.User, error)
}

type getUserHandler struct {
	UserRepo dUser.Repo
}

func (g getUserHandler) Handle(ctx context.Context, query Get) (*dUser.User, error) {
	user, err := g.UserRepo.GetUser(query.Id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewGetUserHandler(userRepo dUser.Repo) GetUserHandler {
	return getUserHandler{UserRepo: userRepo}
}
