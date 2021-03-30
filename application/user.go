package application

import (
	"food-app/domain/entity"
	"food-app/domain/repository"
)

type userApp struct {
	ur repository.UserRepository
}

var _ UserAppInterface = userApp{}

type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}

func (u userApp) SaveUser(user *entity.User) (*entity.User, error) {
	return u.ur.SaveUser(user)
}

func (u userApp) GetUsers() ([]*entity.User, error) {
	return u.ur.GetUsers()
}

func (u userApp) GetUser(id uint64) (*entity.User, error) {
	return u.ur.GetUser(id)
}

func (u userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, error) {
	return u.ur.GetUserByEmailAndPassword(user)
}
