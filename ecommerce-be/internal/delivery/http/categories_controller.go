package http

import (
	"ecommerce/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CategoriesController struct {
	UseCase *usecase.CategoriesUsecase
	Log *logrus.Logger
}

func NewCategoriesController(useCase *usecase.CategoriesUsecase, log *logrus.Logger) *CategoriesController{
	return &CategoriesController{
		UseCase: useCase,
		Log: log,
	}
}

func (c *CategoriesController) List(ctx *fiber.Ctx) error{
	responses, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to list Categories")
		return err
	}

	ctx.Set("content-Type", "application/json")
	return ctx.Send([]byte(responses))
}