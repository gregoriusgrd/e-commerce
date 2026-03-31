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

type CitiesUseCase struct {
	DB               *gorm.DB
	Log              *logrus.Logger
	Validate         *validator.Validate
	CitiesRepository *repository.CitiesRepository
}

func NewCitiesUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, citiesRepository *repository.CitiesRepository) *CitiesUseCase {
	return &CitiesUseCase{
		DB:               db,
		Log:              log,
		Validate:         validate,
		CitiesRepository: citiesRepository,
	}
}

func (c *CitiesUseCase) FindByProvinceId(ctx context.Context, request *model.GetCityRequest) ([]model.CitiesResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	var responses []model.CitiesResponse

	cities, err := c.CitiesRepository.FindByProvinceId(tx, request.ProvinceId)
	if err != nil{
		c.Log.Warnf("failed to load cities : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	for _, city := range cities{
		response := model.CitiesResponse{
			ID: city.ID,
			Name: city.Name,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
