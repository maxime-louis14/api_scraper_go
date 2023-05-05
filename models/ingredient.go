package models

import (
	"time"

	"gorm.io/gorm"
)

type Ingredient struct {
	gorm.Model
	Quantity  string `json:"quantity"`
	Unit      string `json:"unit"`
	RecetteID uint
	CreatedAt time.Time
}
