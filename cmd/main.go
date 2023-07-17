package main

import (
	"idrisgo/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
    app := fiber.New()

	setupRoutes(app)



    app.Listen(":8080")
}