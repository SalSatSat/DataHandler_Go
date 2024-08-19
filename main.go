package main

import (
	"datahandler_go/helpers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	port := helpers.EnvVariable("PORT")

	app := fiber.New()

	// Use middlewares
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${ua} | ${locals:requestid} | ${bytesSent}B\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}))
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// Create a map to hold the JSON data with a status code
		response := map[string]interface{}{
			"message":     "Hello, World!",
			"status_code": fiber.StatusOK, // Use Fiber's built-in status code
		}
		// Send the JSON response
		return c.JSON(response)
	})

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
