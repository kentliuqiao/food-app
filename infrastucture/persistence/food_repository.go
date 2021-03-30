package persistence

import (
	"food-app/domain/entity"
	"food-app/domain/repository"
	"os"

	"github.com/jinzhu/gorm"
)

type foodRepo struct {
	db *gorm.DB
}

var _ repository.FoodRepository = foodRepo{}

func NewFoodRepository(db *gorm.DB) repository.FoodRepository {
	return foodRepo{db: db}
}

func (f foodRepo) SaveFood(food *entity.Food) (*entity.Food, error) {
	food.FoodImage = os.Getenv("DO_SPACES_URL") + food.FoodImage

	err := f.db.Debug().Create(&food).Error
	if err != nil {
		return nil, err
	}

	return food, nil
}

func (f foodRepo) GetFood(id uint64) (*entity.Food, error) {
	food := &entity.Food{}
	if err := f.db.Where("id = ?", id).Take(food).Error; err != nil {
		return nil, err
	}

	return food, nil
}

func (f foodRepo) GetAllFood() ([]*entity.Food, error) {
	foods := []*entity.Food{}
	if err := f.db.Debug().Limit(100).Order("created_at desc").Find(foods).Error; err != nil {
		return nil, err
	}

	return foods, nil
}

func (f foodRepo) UpdateFood(food *entity.Food) (*entity.Food, error) {
	if err := f.db.Debug().Save(food).Error; err != nil {
		return nil, err
	}

	return food, nil
}

func (f foodRepo) DeleteFood(id uint64) error {
	food := &entity.Food{}
	if err := f.db.Where("id = ?", id).Delete(food).Error; err != nil {
		return err
	}

	return nil
}
