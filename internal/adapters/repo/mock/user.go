package mock

import dUser "todo-list-cqrs-es/internal/domain/user"

type UserRepoAdapter struct {
}

func (a UserRepoAdapter) GetUser(id string) (dUser.User, error) {

	return dUser.New(id), nil
}

func NewUserRepoAdapter() *UserRepoAdapter {
	return &UserRepoAdapter{}
}
