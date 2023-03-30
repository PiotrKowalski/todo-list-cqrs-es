package user

type Repo interface {
	GetUser(id string) (User, error)
}
