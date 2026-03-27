package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App *fiber.App
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
}