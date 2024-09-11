package main

import (
	"datahandler_go/database/postgres"
	"datahandler_go/helpers"
	"datahandler_go/models/samples"
	"datahandler_go/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	port := helpers.EnvVariable("PORT")

	postgres.ConnectDb()
	postgres.DB.Db.AutoMigrate(&samples.Postgres_Sample{})

	if postgres.IsDbConnected() {
		fmt.Printf("Postgres successfully connected\n")
	} else {
		fmt.Printf("Postgres failed to connect\n")
	}

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

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
