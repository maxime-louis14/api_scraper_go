package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)

type Instruction struct {
	ID          uint   `gorm:"primaryKey"`
	Number      string `json:"number"`
	Description string `json:"description"`
}

func CreateResponseInstruction(instructionModel models.Instruction) Instruction {
	return Instruction{ID: instructionModel.ID, Number: instructionModel.Number, Description: instructionModel.Description}
}

func PostInstructions(c *fiber.Ctx) error {
	// Ouvrir le fichier data.json
	file, err := os.Open("data.json")
	if err != nil {
		return err
	}
	defer file.Close()

	// Décodez les données JSON dans une variable slice d'instructions
	var instructions []models.Instruction
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&instructions)
	if err != nil {
		return err
	}

	// Parcourir toutes les instructions
	for _, instruction := range instructions {
		// Faire quelque chose avec les données d'instruction
		result := database.Database.Db.Create(&instruction)
		if result.Error != nil {
			return result.Error
		}
	}

	log.Println(instructions)
	fmt.Println(instructions)
	// Réponse HTTP avec un message de succès
	return c.SendString("Recettes ajoutées avec succès")

}

func GetInstructions(c *fiber.Ctx) error {
	instructions := []models.Instruction{}
	database.Database.Db.Find(&instructions)
	responseInstructions := []Instruction{}
	for _, instruction := range instructions {
		responseInstruction := CreateResponseInstruction(instruction)
		responseInstructions = append(responseInstructions, responseInstruction)
	}
	return c.Status(200).JSON(responseInstructions)
}
