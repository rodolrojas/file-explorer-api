package router

import (
	"api/commands"
	"api/responses"

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
	app.Post("/mkdir", func(c *fiber.Ctx) error {
		type requestBody struct {
			Name string `json:"name"`
		}
		var body requestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := commands.MakeDir(body.Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(responses.SuccessResponse())
	})
}