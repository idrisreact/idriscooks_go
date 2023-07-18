package main

import (
	"idrisgo/handlers"

	"github.com/gofiber/fiber/v2"
)


func setupRoutes(app *fiber.App)  {

	recipeGroupRoutes := app.Group("/api/v1")

	recipeGroupRoutes.Get("/recipe",handlers.GetAllRecipes)
	recipeGroupRoutes.Get("/recipe/search",handlers.SearchRecipeByName)
	recipeGroupRoutes.Get("/recipe/:id",handlers.GetRecipeById)
	recipeGroupRoutes.Put("/recipe/:id",handlers.UpdateRecipe)
	recipeGroupRoutes.Post("/recipe",handlers.CreateRecipes)

	app.Get("/", handlers.HomeHandler)
	
}