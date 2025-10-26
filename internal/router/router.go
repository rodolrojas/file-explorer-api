package router

import (
	"api/commands"
	"api/responses"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Global CORS middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		
		return c.Next()
	})

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
			TargetPath string `json:"targetPath"`
			Name string `json:"name"`
		}
		var body requestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := commands.MakeDir(body.TargetPath, body.Name)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(responses.SuccessResponse)
	})
	app.Post("/createFile", func(c *fiber.Ctx) error {
		type requestBody struct {
			TargetPath string `json:"targetPath"`
			Filename   string `json:"filename"`
		}
		var body requestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}
		err := commands.CreateFile(body.TargetPath, body.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(responses.SuccessResponse)
	})
}