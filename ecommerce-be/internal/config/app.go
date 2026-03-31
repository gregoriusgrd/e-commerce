package config

import (
	"ecommerce/internal/delivery/http"
	"ecommerce/internal/delivery/http/route"
	"ecommerce/internal/repository"
	"ecommerce/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	Config   *viper.Viper
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	categoriesRepository := repository.NewCategoriesRepositories(config.Log)
	provincesRepository := repository.NewProvincesRepositories(config.Log)
	citiesRepository := repository.NewCitiesRepository(config.Log)
	//setup use cases
	categoriesUseCase := usecase.NewCategoriesUseCase(config.DB, config.Log, config.Validate, categoriesRepository)
	provincesUseCase := usecase.NewProvincesUseCase(config.DB, config.Log, config.Validate, provincesRepository)
	citiesUseCase := usecase.NewCitiesUseCase(config.DB, config.Log, config.Validate, citiesRepository)
	//setup controller
	categoriesController := http.NewCategoriesController(categoriesUseCase, config.Log)
	provincesController := http.NewProvincesController(provincesUseCase, config.Log)
	citiesController := http.NewCitiesController(citiesUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:                  config.App,
		CategoriesController: categoriesController,
		ProvincesController:  provincesController,
		CitiesController:     citiesController,
	}

	routeConfig.Setup()
}
