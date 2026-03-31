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
	App *fiber.App
	Config *viper.Viper
	DB *gorm.DB
	Log *logrus.Logger
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	categoriesRepository := repository.NewCategoriesRepositories(config.Log)

	//setup use cases
	categoriesUseCase := usecase.NewCategoriesUseCase(config.DB, config.Log, config.Validate, categoriesRepository)
	//setup controller
	categoriesController := http.NewCategoriesController(categoriesUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App: config.App,
		CategoriesController: categoriesController,
	}

	routeConfig.Setup()
}