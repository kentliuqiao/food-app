package persistence

import (
	"fmt"
	"food-app/domain/entity"
	"food-app/domain/repository"

	"github.com/jinzhu/gorm"
)

type Repositories struct {
	User repository.UserRepository
	Food repository.FoodRepository

	db *gorm.DB
}

func NewRepositories(dbDriver, dbUser, dbPassword, dbPort, dbHost, dbName string) (*Repositories, error) {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	db, err := gorm.Open(dbDriver, dbURL)
	if err != nil {
		return nil, err
	}
	db.Debug()

	return &Repositories{
		User: NewUserRepository(db),
		Food: NewFoodRepository(db),
		db:   db,
	}, nil
}

func (s *Repositories) Close() error {
	return s.db.Close()
}

func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.Food{}, &entity.User{}).Error
}
