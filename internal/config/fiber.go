package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var FiberConfig fiber.Config

func LoadFiberConfig() {
	// Create a new engine
	engine := html.New("web/templates", ".html")

	FiberConfig = fiber.Config{
		ReadTimeout: AppConfig.ReadTimeoutSecs,
		Views:       engine,
	}
}
