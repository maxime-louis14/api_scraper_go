package models

import (
	"time"

	"gorm.io/gorm"
)

// Recette represents a recipe
//
// swagger:model
type Recette struct {
	// ID of the recette
	gorm.Model
	// Instructions of the recette
	//
	// required: true
	Instructions []Instruction
	// Name of the recette
	Name string `json:"name"`

	Page string `json:"page"`

	Image string `json:"image"`
	// Image of the recette
	Ingredients []Ingredient
	// Ingredients of the recette
	//
	// required: true
	CreatedAt time.Time
}
