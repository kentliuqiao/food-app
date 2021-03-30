package persistence

import (
	"food-app/domain/entity"
	"food-app/domain/repository"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepo{db: db}
}

func (u userRepo) SaveUser(user *entity.User) (*entity.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepo) GetUser(id uint64) (*entity.User, error) {
	user := &entity.User{}
	if err := u.db.Where("id = ?", id).Take(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepo) GetUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u userRepo) GetUserByEmailAndPassword(user *entity.User) (*entity.User, error) {
	res := &entity.User{}
	if err := u.db.Where("email = ?", user.Email).Take(res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
