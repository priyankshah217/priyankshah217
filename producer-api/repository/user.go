package repository

import (
	"github.com/priyankshah217/model"
)

type UserRepository struct {
	Users map[string]*model.User
}

func (u *UserRepository) GetUsers() []model.User {
	var response []model.User
	for _, user := range u.Users {
		response = append(response, *user)
	}
	return response
}

func (u *UserRepository) ByUserName(userName string) (*model.User, error) {
	if user, ok := u.Users[userName]; ok {
		return user, nil
	}
	return nil, model.ErrNotFound
}

func (u *UserRepository) ByID(ID int) (*model.User, error) {
	for _, user := range u.Users {
		if user.ID == ID {
			return user, nil
		}
	}
	return nil, model.ErrNotFound
}
