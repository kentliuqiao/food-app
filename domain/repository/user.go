package repository

import "food-app/domain/entity"

// UserRepository define user-related bussiness logic
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUsers() ([]*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, error)
}
