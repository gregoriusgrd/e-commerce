package main

import (
	"ecommerce/internal/config"
	"fmt"
	"log"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)


	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}