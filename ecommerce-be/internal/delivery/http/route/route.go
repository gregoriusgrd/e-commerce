package route

import (
	"ecommerce/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
	CategoriesController *http.CategoriesController
}

func(c *RouteConfig) Setup(){
	c.SetupGuestRoute()
}

func(c *RouteConfig) SetupGuestRoute(){
	c.App.Get("/api/categories", c.CategoriesController.List)
}