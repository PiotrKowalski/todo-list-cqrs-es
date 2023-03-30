package api

type Status string

var (
	StatusCompleted   Status = "completed"
	StatusUnCompleted Status = "uncompleted"
)

type CreateTodoListRequest struct {
	Name string `json:"name"`
}

type GetTodoListRequest struct {
	Id string `json:"id"`
}

type TodoListResponse struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Entries []Entry
}

type Entry struct {
	Id          uint   `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
