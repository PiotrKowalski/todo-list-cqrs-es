package todolist

import (
	"context"
	dTodo_list "todo-list-cqrs-es/internal/domain/todo_list"
	"todo-list-cqrs-es/internal/ports/api"
)

type Create struct {
	TodoList api.CreateTodoListRequest
}

type CreateTodoListHandler interface {
	Handle(ctx context.Context, cmd Create) (*dTodo_list.TodoList, error)
}

type createTodoListHandler struct {
	TodoListRepo dTodo_list.TodoListRepo
	//GetTodoList  cTodolist.GetTodoListHandler
}

func (h createTodoListHandler) Handle(ctx context.Context, cmd Create) (*dTodo_list.TodoList, error) {

	todoList := dTodo_list.NewFromCreateTodoListRequest(cmd.TodoList)

	ret, err := h.TodoListRepo.Save(todoList)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func NewCreateTodoListHandler(todoListRepo dTodo_list.TodoListRepo) CreateTodoListHandler {
	return createTodoListHandler{TodoListRepo: todoListRepo}
}
