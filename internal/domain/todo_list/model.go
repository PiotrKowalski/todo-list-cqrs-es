package todo_list

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-list-cqrs-es/internal/ports/api"
)

type Status string

var (
	StatusCompleted   Status = "completed"
	StatusUnCompleted Status = "uncompleted"
)

type TodoList struct {
	Id   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `json:"name" bson:"name"`

	Entries []entry `json:"entries" bson:"entries"`
}

func (t *TodoList) GetId() string {
	return t.Id.Hex()
}

func (t *TodoList) GetName() string {
	return t.Name
}

func (t *TodoList) GetEntries() []entry {
	return t.Entries
}

func (t *TodoList) AddEntry(e entry) {
	t.Entries = append(t.Entries, e)
}

func (t *TodoList) ToTodoListResponse() *api.TodoListResponse {

	out := &api.TodoListResponse{}
	out.Id = t.Id.Hex()
	out.Name = t.Name

	for _, e := range t.Entries {
		out.Entries = append(out.Entries, api.Entry{
			Id:          e.Id,
			Description: e.Description,
			Status:      api.Status(e.Status),
		})
	}

	return out
}

type entry struct {
	Id          uint   `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	Status      Status `json:"status" bson:"status"`
}

func (e entry) GetId() uint {
	return e.Id
}

func (e entry) GetDescription() string {
	return e.Description
}

func (e entry) GetStatus() Status {
	return e.Status
}

func NewFromJSON(data []byte) (*TodoList, error) {
	var todolist TodoList

	if err := json.Unmarshal(data, &todolist); err != nil {
		return nil, err
	}

	return &todolist, nil
}

func New(id string) TodoList {
	return TodoList{}
}
