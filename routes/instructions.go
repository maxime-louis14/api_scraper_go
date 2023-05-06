package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)



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

func findInstruction(id int, instruction *models.Instruction) error {
	database.Database.Db.Find(&instruction, "id = ?", id)
	if instruction.ID == 0 {
		return errors.New("instructions does not exist")
	}
	return nil
}

func GetInstruction(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var instruction models.Instruction
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findInstruction(id, &instruction); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseInstruction := CreateResponseInstruction(instruction)
	return c.Status(200).JSON(responseInstruction)
}
