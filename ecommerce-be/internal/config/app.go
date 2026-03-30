package config

import (
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
}

func Bootstrap(config *BootstrapConfig) {
}