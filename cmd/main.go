package main

import (
	"idrisgo/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	database.ConnectDb()
    app := fiber.New()

	app.Use(recover.New())

	setupRoutes(app)



    app.Listen(":8080")
}