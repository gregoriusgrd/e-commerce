package http

import (
	"ecommerce/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ProvincesController struct {
	UseCase *usecase.ProvincesUsecase
	Log     *logrus.Logger
}

func NewProvincesController(useCase *usecase.ProvincesUsecase, log *logrus.Logger) *ProvincesController {
	return &ProvincesController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *ProvincesController) List(ctx *fiber.Ctx) error {
	responses, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.WithError(err).Error("failed to list provinces")
		return err
	}

	return ctx.JSON(fiber.Map{
		"data": responses,
	})
}
