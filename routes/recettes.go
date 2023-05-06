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

type Recette struct {
	Name  string `json:"name"`
	Link  string `json:"liens"`
	Image string `json:"image"`
}

type Instruction struct {
	ID          uint   `gorm:"primaryKey"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

type Ingredient struct {
	Quantity string `json:"quantity"`
	Unit     string `json:"unit"`
}

func CreateResponseIngredient(ingredientModel models.Ingredient) Ingredient {
	return Ingredient{Quantity: ingredientModel.Quantity, Unit: ingredientModel.Unit}
}

func CreateResponseInstruction(instructionModel models.Instruction) Instruction {
	return Instruction{ID: instructionModel.ID, Number: instructionModel.Number, Description: instructionModel.Description}
}

func CreateResponseRecette(recetteModel models.Recette) Recette {
	return Recette{Name: recetteModel.Name, Link: recetteModel.Link, Image: recetteModel.Image}
}

func PostRecette(c *fiber.Ctx) error {
	// Ouvrir le fichier data.json
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Décodez les données JSON dans une variable slice de recettes
	var recettes []models.Recette
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&recettes)
	if err != nil {
		return err
	}

	// Insérer chaque recette dans la base de données MySQL
	for _, recette := range recettes {
		// Utiliser la méthode Create de GORM pour insérer une recette dans la base de données
		// '&' est utilisé pour prendre l'adresse de la variable recette, car la méthode Create attend un pointeur de recette.
		result := database.Database.Db.Create(&recette)
		if result.Error != nil {
			return result.Error
		}
	}
	fmt.Println(recettes)

	// Réponse HTTP avec un message de succès
	return c.SendString("Recettes ajoutées avec succès")
}

func GetRecettesDetails(c *fiber.Ctx) error {
	recettes := []models.Recette{}

	// récupérer toutes les recettes de la base de données
	database.Database.Db.Find(&recettes)

	responseRecettes := []struct {
		Name         string        `json:"name"`
		Link         string        `json:"liens"`
		Image        string        `json:"image"`
		Ingredients  []Ingredient  `json:"ingredients"`
		Instructions []Instruction `json:"instructions"`
	}{}

	// pour chaque recette, récupérer ses ingrédients et instructions
	for _, recette := range recettes {
		var ingredients []models.Ingredient
		var instructions []models.Instruction

		// récupérer les ingrédients de la recette
		database.Database.Db.Where("recette_id = ?", recette.ID).Find(&ingredients)

		// récupérer les instructions de la recette
		database.Database.Db.Where("recette_id = ?", recette.ID).Find(&instructions)

		// créer une réponse contenant toutes les informations de la recette
		responseRecette := struct {
			Name         string        `json:"name"`
			Link         string        `json:"liens"`
			Image        string        `json:"image"`
			Ingredients  []Ingredient  `json:"ingredients"`
			Instructions []Instruction `json:"instructions"`
		}{
			Name:         recette.Name,
			Link:         recette.Link,
			Image:        recette.Image,
			Ingredients:  make([]Ingredient, len(ingredients)),
			Instructions: make([]Instruction, len(instructions)),
		}

		// copier les informations des ingrédients et instructions dans la réponse
		for i, ingredient := range ingredients {
			responseRecette.Ingredients[i] = CreateResponseIngredient(ingredient)
		}
		for i, instruction := range instructions {
			responseRecette.Instructions[i] = CreateResponseInstruction(instruction)
		}

		responseRecettes = append(responseRecettes, responseRecette)
	}

	return c.Status(200).JSON(responseRecettes)
}

func findRecette(id int, recette *models.Recette) error {
	database.Database.Db.Find(&recette, "id = ?", id)
	if recette.ID == 0 {
		return errors.New("Recette does not exist")
	}
	return nil
}

func GetRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Veuillez vous assurer que !id est un interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func GetRecetteByName(c *fiber.Ctx) error {
	name := c.Params("name")

	recette := models.Recette{}
	err := database.Database.Db.Where("name = ?", name).First(&recette)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Recipe not found",
		})
	}

	responseRecette := struct {
		Recette      models.Recette       `json:"recette"`
		Ingredients  []models.Ingredient  `json:"ingredients"`
		Instructions []models.Instruction `json:"instructions"`
	}{
		Recette:      recette,
		Ingredients:  recette.Ingredients,
		Instructions: recette.Instructions,
	}

	return c.JSON(responseRecette)
}

func GetRecettesIngredient(c *fiber.Ctx) error {
	unit := c.Params("unit")

	// Requête pour récupérer toutes les recettes contenant l'unité spécifiée
	recettes := []models.Recette{}
	database.Database.Db.Where("id IN (SELECT recette_id FROM ingredients WHERE unit = ?)", unit).
		Order("name ASC").
		Find(&recettes)

	// Vérifier si des recettes ont été trouvées
	if len(recettes) == 0 {
		// Retourner une réponse vide avec le code 204 No Content
		return c.SendStatus(fiber.StatusNoContent)
	}

	responseRecettes := []struct {
		Recette      models.Recette       `json:"recette"`
		Ingredients  []models.Ingredient  `json:"ingredients"`
		Instructions []models.Instruction `json:"instructions"`
	}{}

	// Pour chaque recette trouvée, récupérer ses ingrédients et instructions associés
	for _, recette := range recettes {
		ingredients := []models.Ingredient{}
		database.Database.Db.Model(&recette).Association("Ingredients").Find(&ingredients)

		instructions := []models.Instruction{}
		database.Database.Db.Model(&recette).Association("Instructions").Find(&instructions)

		responseRecette := struct {
			Recette      models.Recette       `json:"recette"`
			Ingredients  []models.Ingredient  `json:"ingredients"`
			Instructions []models.Instruction `json:"instructions"`
		}{
			Recette:      recette,
			Ingredients:  ingredients,
			Instructions: instructions,
		}
		responseRecettes = append(responseRecettes, responseRecette)
	}

	return c.JSON(responseRecettes)

}
