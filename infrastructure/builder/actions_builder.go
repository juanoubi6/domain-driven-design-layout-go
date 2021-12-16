package builder

import (
	"domain-driven-design-layout/domain/actions/users"
)

type Actions struct {
	CreateUser users.CreateUser
}

func CreateActions(repositories *Repositories) (*Actions, error) {
	createUser, err := users.NewCreateUserAction(repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	return &Actions{
		CreateUser: createUser,
	}, nil
}
