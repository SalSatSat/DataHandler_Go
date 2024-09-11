package routes

import (
	"datahandler_go/database/postgres"
	"datahandler_go/models/samples"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Create a map to hold the JSON data with a status code
		response := map[string]interface{}{
			"message":     "Hello, World!",
			"status_code": fiber.StatusOK, // Use Fiber's built-in status code
		}
		// Send the JSON response
		return c.JSON(response)
	})

	app.Get("/postgres_sample", func(c *fiber.Ctx) error {
		samples := []samples.Postgres_Sample{}
		postgres.DB.Db.Find(&samples)

		// Create a map to hold the JSON data with a status code
		response := map[string]interface{}{
			"data":        samples,
			"status_code": fiber.StatusOK, // Use Fiber's built-in status code
		}
		// Send the JSON response
		return c.JSON(response)
	})

	app.Post("/postgres_sample", func(c *fiber.Ctx) error {
		sample := new(samples.Postgres_Sample)
		if err := c.BodyParser(sample); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		postgres.DB.Db.Create(&sample)

		// Create a map to hold the JSON data with a status code
		response := map[string]interface{}{
			"data":        sample,
			"status_code": fiber.StatusOK, // Use Fiber's built-in status code
		}
		// Send the JSON response
		return c.JSON(response)
	})
}
