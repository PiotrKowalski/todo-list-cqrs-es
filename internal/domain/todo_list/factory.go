package todo_list

import "todo-list-cqrs-es/internal/ports/api"

func NewFromCreateTodoListRequest(request api.CreateTodoListRequest) *TodoList {
	return &TodoList{
		Name: request.Name,
	}

}
