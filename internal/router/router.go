package router

import (
	"api/commands"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/list", func(c *fiber.Ctx) error {
		path := c.Query("path", "/")
		files, err := commands.ListFiles(path, false)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(files)
	})
	app.Get("/listDirs", func(c *fiber.Ctx) error {
		path := c.Query("path", "/")
		files, err := commands.ListFiles(path, true)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(files)
	})
}