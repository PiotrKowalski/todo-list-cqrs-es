package user

type User interface {
	GetId() string
}

type user struct {
	Id string `json:"id"`
}

func (u user) GetId() string {
	//TODO implement me
	return u.Id
}

func New(id string) User {
	return user{Id: id}
}
