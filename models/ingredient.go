package models

import (
	"time"

	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	ID           int `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Ingredients  string  `json:"ingredients"`
	Photos       string  `json:"photos"`
	Instructions string  `json:"instructions"`
	RecetteID    int    // ID de la recette à laquelle appartient l'ingrédient
	Recette      Recette `gorm:"foreignKey:RecetteID"` // Relation Many-to-One avec Recette
	UpdatedAt    int     // Set to current unix seconds on updating or if it is zero on creating
	Updated      int64   `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	Deleted_at   gorm.DeletedAt
}
