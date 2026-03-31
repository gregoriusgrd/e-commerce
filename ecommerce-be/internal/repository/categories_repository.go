package repository

import (
	"ecommerce/internal/entity"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoriesRepository struct {
	Repository[entity.Categories]
	Log *logrus.Logger
}

func NewCategoriesRepositories(log *logrus.Logger) *CategoriesRepository {
	return &CategoriesRepository{
		Log: log,
	}
}

func (r *CategoriesRepository) FindAllParents(tx *gorm.DB) ([]entity.Categories, error) {
	var categories []entity.Categories
	err := tx.Where("parent_id is null").Preload("Children").Find(&categories).Error
	if err != nil {
		fmt.Println("GORM ERROR:", err) //
		return nil, err
	}
	return categories, nil
}
