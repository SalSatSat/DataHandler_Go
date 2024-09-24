package main

import (
	"datahandler_go/helpers"
	"datahandler_go/jobs"
	"datahandler_go/routes"
	"fmt"
	"log"
	"sync"

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

	routes.SetupRoutes(app)

	// WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Start jobs in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		jobs.Jobs() // Start running jobs
	}()

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
