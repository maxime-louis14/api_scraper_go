package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)


func PostIngredients(c *fiber.Ctx) error {
	// Ouvrir le fichier data.json
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Décodez les données JSON dans une variable slice de ingredients
	var ingredients []models.Ingredient
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ingredients)
	if err != nil {
		return err
	}

	// Insérer chaque recette dans la base de données MySQL
	for _, ingredient := range ingredients {
		// '&' est utilisé pour prendre l'adresse de la variable recette, car la méthode Create attend un pointeur de recette.
		result := database.Database.Db.Create(&ingredient)
		if result.Error != nil {
			return result.Error
		}
	}

	fmt.Println(ingredients)
	// Réponse HTTP avec un message de succès
	return c.SendString("Recettes ajoutées avec succès")
}

func GetIngredients(c *fiber.Ctx) error {
	ingredients := []models.Ingredient{}
	database.Database.Db.Find(&ingredients)
	responseIngredients := []Ingredient{}
	for _, ingredient := range ingredients {
		responseIngredient := CreateResponseIngredient(ingredient)
		responseIngredients = append(responseIngredients, responseIngredient)
	}
	return c.Status(200).JSON(responseIngredients)
}

func findIngredients(id int, ingredient *models.Ingredient) error {
	database.Database.Db.Find(&ingredient, "id = ?", id)
	if ingredient.ID == 0 {
		return errors.New("ingredient does not exist")
	}
	return nil
}

func GetIngredient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var Ingredient models.Ingredient
	if err != nil {
		return c.Status(400).JSON("please ensure that !id is an interger")
	}
	if err := findIngredients(id, &Ingredient); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseIngredient := CreateResponseIngredient(Ingredient)
	return c.Status(200).JSON(responseIngredient)
}
