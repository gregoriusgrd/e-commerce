package http

import (
	"ecommerce/internal/model"
	"ecommerce/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CitiesController struct {
	UseCase *usecase.CitiesUseCase
	Log     *logrus.Logger
}

func NewCitiesController(usecase *usecase.CitiesUseCase, log *logrus.Logger) *CitiesController {
	return &CitiesController{
		UseCase: usecase,
		Log:     log,
	}
}

func (c *CitiesController) FindByProvinceId(ctx *fiber.Ctx) error {
	request := new(model.GetCityRequest)
	if err := ctx.QueryParser(request); err != nil {
		c.Log.Warnf("Failed to parse request : %+v", err)
		return fiber.ErrBadRequest
	}
	responses, err := c.UseCase.FindByProvinceId(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get cities")
		return err
	}

	return ctx.JSON(fiber.Map{
		"data": responses,
	})
}
