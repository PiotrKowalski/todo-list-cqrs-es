package app

import (
	"context"
	"todo-list-cqrs-es/internal/adapters/repo/mock"
	"todo-list-cqrs-es/internal/adapters/repo/mongodb"
	cTodolist "todo-list-cqrs-es/internal/app/commands/todolist"
	qTodolist "todo-list-cqrs-es/internal/app/queries/todolist"
	qUser "todo-list-cqrs-es/internal/app/queries/user"
	dUser "todo-list-cqrs-es/internal/domain/user"
	"todo-list-cqrs-es/internal/ports/api"
)

type Application struct {
	Commands commands
	Queries  queries
}

type commands struct {
	CreateTodoListHandler cTodolist.CreateTodoListHandler
}

type queries struct {
	GetUserHandler     qUser.GetUserHandler
	GetTodoListHandler qTodolist.GetTodoListHandler
}

type Config struct {
	MongoDBURI string
}

func NewApplication(config Config) *Application {
	todoListAdapter := mongodb.NewTodoListRepoAdapter(config.MongoDBURI)
	userRepo := mock.NewUserRepoAdapter()

	return &Application{
		Commands: commands{
			CreateTodoListHandler: cTodolist.NewCreateTodoListHandler(todoListAdapter),
		},
		Queries: queries{
			GetUserHandler:     qUser.NewGetUserHandler(userRepo),
			GetTodoListHandler: qTodolist.NewGetTodoListHandler(todoListAdapter),
		},
	}
}

func (a Application) GetUser(ctx context.Context, id string) (*dUser.User, error) {
	q := qUser.Get{Id: id}

	user, err := a.Queries.GetUserHandler.Handle(ctx, q)
	if err != nil {
		return nil, nil
	}

	return user, nil
}

func (a Application) CreateTodoList(ctx context.Context, list api.CreateTodoListRequest) (*api.TodoListResponse, error) {

	cmd := cTodolist.Create{TodoList: list}

	ret, err := a.Commands.CreateTodoListHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return ret.ToTodoListResponse(), nil

}

func (a Application) GetTodoList(ctx context.Context, id string) (*api.TodoListResponse, error) {

	q := qTodolist.Get{Id: id}

	ret, err := a.Queries.GetTodoListHandler.Handle(ctx, q)
	if err != nil {
		return nil, err
	}

	return ret.ToTodoListResponse(), nil
}

func (a Application) DeleteTodoList(ctx context.Context, id string) error {
	panic("Not implemented")
}
