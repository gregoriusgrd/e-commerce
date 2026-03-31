package usecase

import (
	"context"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoriesUsecase struct {
	DB *gorm.DB
	Log *logrus.Logger
	Validate *validator.Validate
	CategoriesRepository *repository.CategoriesRepository
	
}

func NewCategoriesUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, CategoriesRepository *repository.CategoriesRepository) *CategoriesUsecase{
	return &CategoriesUsecase{
		DB: db,
		Log: log,
		Validate: validate,
		CategoriesRepository: CategoriesRepository,
	}
}

func (c *CategoriesUsecase) List(ctx context.Context) (string, error){
	var responses []model.CategoryResponse

	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	//query to database
	//load parent categories
	categories, err := c.CategoriesRepository.FindAllParents(tx)
	if err != nil{
		c.Log.WithError(err).Error("Failed to load parent categories")
		return "", fiber.ErrInternalServerError
	}
	// iterate each parent
	for _, parent := range categories{
		response := model.CategoryResponse{
			ID: parent.ID,
			Name: parent.Name,
		}
		var childResponses []model.CategoryResponse
		for _, child := range parent.Children{
			responseChild := model.CategoryResponse{
				ID: child.ID,
				Name: child.Name,
			}
			childResponses = append(childResponses, responseChild)
		}
		response.Children = childResponses
		responses = append(responses, response)
	}

	jsonValue, err := json.Marshal(fiber.Map{
		"data" : responses,
	})
	if err != nil{
		c.Log.WithError(err).Error("Failed to load parent categories")
		return "", fiber.ErrInternalServerError
	}

	return string(jsonValue), nil
}