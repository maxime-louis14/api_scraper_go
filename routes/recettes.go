package routes

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)

type Recette struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	Page         string `json:"line"`
}

func CreateResponseRecette(recetteModel models.Recette) Recette {
	return Recette{ID: recetteModel.ID, Name: recetteModel.Name, Descriptions: recetteModel.Descriptions, Page: recetteModel.Page}
}

func CreateRecette(c *fiber.Ctx) error {
	var recette models.Recette
	if err := c.BodyParser(&recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&recette)
	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
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
	// Réponse HTTP avec un message de succès
	return c.SendString("Recettes ajoutées avec succès")
}

func GetRecettes(c *fiber.Ctx) error {
	recettes := []models.Recette{}
	database.Database.Db.Find(&recettes)
	responseRecettes := []Recette{}
	for _, recette := range recettes {
		responseRecette := CreateResponseRecette(recette)
		responseRecettes = append(responseRecettes, responseRecette)
	}
	return c.Status(200).JSON(responseRecettes)
}

func findRecette(id int, recette *models.Recette) error {
	database.Database.Db.Find(&recette, "id = ?", id)
	if recette.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func UpdateRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateRecette struct {
		Name         string `json:"name"`
		Descriptions string `json:"descriptions"`
		Ingredients  string `json:"ingredients"`
		Photos       string `json:"photos"`
		Instructions string `json:"instructions"`
		Page         string `json:"line"`
		SerialNumber string `json:"serial_number"`
	}
	var updateData UpdateRecette
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	recette.Name = updateData.Name
	recette.Page = updateData.Page

	database.Database.Db.Save(&recette)

	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func DeleteRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&recette).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Recette supprimée avec succès")
}
