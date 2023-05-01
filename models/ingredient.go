package models

import (
	"time"
)

type Ingredient struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Ingredients  string  `json:"ingredients"`
	Photos       string  `json:"photos"`
	Instructions string  `json:"instructions"`
	RecetteRefer uint    `json:"recette_id"`
	Recette      Recette `gorm:"foreignKey:RecetteRefer"`
}