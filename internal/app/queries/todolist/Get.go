package todolist

import (
	"context"
	dTodo_list "todo-list-cqrs-es/internal/domain/todo_list"
)

type Get struct {
	Id string
}

type GetTodoListHandler interface {
	Handle(ctx context.Context, cmd Get) (*dTodo_list.TodoList, error)
}

type getTodoListHandler struct {
	TodoListRepo dTodo_list.TodoListRepo
}

func (h getTodoListHandler) Handle(ctx context.Context, cmd Get) (*dTodo_list.TodoList, error) {
	todoList, err := h.TodoListRepo.Load(cmd.Id)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func NewGetTodoListHandler(todoListRepo dTodo_list.TodoListRepo) GetTodoListHandler {
	return getTodoListHandler{TodoListRepo: todoListRepo}
}
