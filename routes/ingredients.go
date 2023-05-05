package routes

import (
	"encoding/json"
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
