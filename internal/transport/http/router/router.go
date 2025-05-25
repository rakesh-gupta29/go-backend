package router

import (
	"github.com/gofiber/fiber/v2"
)

func MountAllRoutes(app *fiber.App) {

	app.Static("/static", "./static", fiber.Static{
		Browse: false,
		Index:  "index.html",
	})

	mountApiRoutes(app)
	mountBookRoutes(app)
	mountWebRoutes(app)
	mountPrivateRoutes(app)

}
