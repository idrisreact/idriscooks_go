package handlers

import (
	"errors"
	"idrisgo/database"
	"idrisgo/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)



func HomeHandler(c *fiber.Ctx) error {
	
	return c.SendString("Hello, Welcome to IdrisCooks API")

}

func GetAllRecipes(c *fiber.Ctx) error {

 	var recipes []models.Recipe
	database.Db.DB.Find(&recipes)

	recipeResponses := make([]models.RecipeResponse, len(recipes))

	for i, recipe := range recipes {
		recipeResponses[i] = models.RecipeResponse{
			ID:           recipe.ID,
			Name:         recipe.Name,
			Duration:     recipe.Duration,
			Serves:       recipe.Serves,
			Instructions: recipe.Instructions,
			CreatedAt:    recipe.CreatedAt,
		}
	}



	return c.JSON(recipeResponses)

}

func GetRecipeById(c *fiber.Ctx) (error) {

	recipeId,err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"invalid Id ",
		})
	}

	var recipe models.Recipe

	result := database.Db.DB.First(&recipe,recipeId)

	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"recipe not found",
		})
	}else if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"failed to retrieve recipe",
		})
	}



	return c.JSON(recipe)
}

func CreateRecipes(c *fiber.Ctx) error {

	recipe := new(models.Recipe)

	if err := c.BodyParser(recipe); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	database.Db.DB.Create(recipe)

	return c.Status(fiber.StatusOK).JSON(recipe)
}


func UpdateRecipe(c *fiber.Ctx) error {

	recipeId,err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"invalid id",
		})
	}

	var recipe models.Recipe

	result := database.Db.DB.First(&recipe,recipeId)

	if errors.Is(result.Error,gorm.ErrRecordNotFound){
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"recipe not found with that ID",
		})
	}else if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"failed to retrieve recipe",
		})
	}

	UpdateRecipe :=new(models.Recipe)

	if err := c.BodyParser(UpdateRecipe); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Invalid request body",
		})
	}

	recipe.Name = UpdateRecipe.Name
	recipe.Serves = UpdateRecipe.Serves
	recipe.Instructions = UpdateRecipe.Instructions
	recipe.Duration = UpdateRecipe.Duration

	result = database.Db.DB.Save(&recipe)

	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"failed to update recipe",
		})
	}
return c.JSON(recipe)

	return c.SendString("update recipe")
}