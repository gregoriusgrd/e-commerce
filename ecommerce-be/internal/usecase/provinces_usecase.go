package usecase

import (
	"context"
	"ecommerce/internal/model"
	"ecommerce/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProvincesUsecase struct {
	DB *gorm.DB
	Log *logrus.Logger
	Validate *validator.Validate
	ProvincesRepository *repository.ProvincesRepository
}

func NewProvincesUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, ProvincesRepository *repository.ProvincesRepository) *ProvincesUsecase{
	return &ProvincesUsecase{
		DB: db,
		Log: log,
		Validate: validate,
		ProvincesRepository: ProvincesRepository,
	}
}

func (c *ProvincesUsecase) List(ctx context.Context)([]model.ProvincesResponse, error){
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	var responses []model.ProvincesResponse

	provinces, err := c.ProvincesRepository.FindAll(tx)
	if err != nil{
		c.Log.Warnf("Failed to load provinces : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	for _, province := range provinces{
		response := model.ProvincesResponse{
			ID: province.ID,
			Name: province.Name,
		}
		responses = append(responses, response)
	}

	return responses, nil
}