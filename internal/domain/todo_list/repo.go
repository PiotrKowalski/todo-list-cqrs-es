package todo_list

type TodoListRepo interface {
	Save(list *TodoList) (*TodoList, error)
	Load(id string) (*TodoList, error)
	Remove(id string) error
}
