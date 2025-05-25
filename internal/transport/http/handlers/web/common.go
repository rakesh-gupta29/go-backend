package handlers

import "github.com/gofiber/fiber/v2"

// TODO: how can we use a function of this type as a handler?
func HandleRenderHomepage(c *fiber.Ctx) error {
	return c.Render("pages/home", fiber.Map{
		"Title": "Home",
	})
}

func HandleRenderAbout(c *fiber.Ctx) error {
	return c.Render("pages/about", fiber.Map{
		"Title": "About",
	})
}
