package routes

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)

type Ingredient struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Ingredients  string `json:"ingredients"`
	Photos       string `json:"photos"`
	Instructions string `json:"instructions"`
}

func CreateResponseIngredients(ingredientsModel models.Ingredient) Ingredient {
	return Ingredient{ID: ingredientsModel.ID, Ingredients: ingredientsModel.Ingredients, Photos: ingredientsModel.Photos, Instructions: ingredientsModel.Instructions}
}

func PostIngredients(c *fiber.Ctx) error {
	// Ouvrir le fichier data.json
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Décodez les données JSON dans une variable slice de recettes
	var ingredients []models.Ingredient
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ingredients)
	if err != nil {
		return err
	}

	// Insérer chaque recette dans la base de données MySQL
	for _, ingredients := range ingredients {

		// Utiliser la méthode Create de GORM pour insérer une recette dans la base de données
		// '&' est utilisé pour prendre l'adresse de la variable recette, car la méthode Create attend un pointeur de recette.
		result := database.Database.Db.Create(&ingredients)
		if result.Error != nil {
			return result.Error
		}
	}
	// Réponse HTTP avec un message de succès
	return c.SendString("Recettes ajoutées avec succès")
}
