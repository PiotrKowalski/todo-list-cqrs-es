package app

import (
	"context"
	"todo-list-cqrs-es/internal/adapters/repo/mongodb"
	cTodolist "todo-list-cqrs-es/internal/app/commands/todolist"
	qTodolist "todo-list-cqrs-es/internal/app/queries/todolist"
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
	GetTodoListHandler qTodolist.GetTodoListHandler
}

type Config struct {
	MongoDBURI string
}

func NewApplication(config Config) *Application {
	todoListAdapter := mongodb.NewTodoListRepoAdapter(config.MongoDBURI)

	return &Application{
		Commands: commands{
			CreateTodoListHandler: cTodolist.NewCreateTodoListHandler(todoListAdapter),
		},
		Queries: queries{
			GetTodoListHandler: qTodolist.NewGetTodoListHandler(todoListAdapter),
		},
	}
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
