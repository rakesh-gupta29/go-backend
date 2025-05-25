package router

import (
	"github.com/gofiber/fiber/v2"
	handlers "github.com/rakesh-gupta29/portal/internal/transport/http/handlers/web"
)

func mountWebRoutes(app *fiber.App) *fiber.App {
	web := app.Group("/")
	web.Get("/", handlers.HandleRenderHomepage)
	web.Get("/about", handlers.HandleRenderAbout)

	return app
}
