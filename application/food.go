package application

import (
	"food-app/domain/entity"
	"food-app/domain/repository"
)

type foodApp struct {
	fr repository.FoodRepository
}

type FoodAppInterface interface {
	SaveFood(*entity.Food) (*entity.Food, error)
	GetAllFood() ([]*entity.Food, error)
	GetFood(uint64) (*entity.Food, error)
	UpdateFood(*entity.Food) (*entity.Food, error)
	DeleteFood(uint64) error
}

var _ FoodAppInterface = foodApp{}

func (f foodApp) SaveFood(food *entity.Food) (*entity.Food, error) {
	return f.fr.SaveFood(food)
}

func (f foodApp) GetAllFood() ([]*entity.Food, error) {
	return f.fr.GetAllFood()
}

func (f foodApp) GetFood(id uint64) (*entity.Food, error) {
	return f.fr.GetFood(id)
}

func (f foodApp) UpdateFood(food *entity.Food) (*entity.Food, error) {
	return f.fr.UpdateFood(food)
}

func (f foodApp) DeleteFood(id uint64) error {
	return f.fr.DeleteFood(id)
}
