package models

import (
	"time"

	"gorm.io/gorm"
)

type Recette struct {
	ID           int `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string       `json:"name"`
	Descriptions string       `json:"descriptions"`
	Page         string       `json:"liens"`
	Ingredients  []Ingredient `gorm:"belongsTo:recette_ingredients"`
	UpdatedAt    int          // Set to current unix seconds on updating or if it is zero on creating
	Updated      int64        `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	Deleted_at   gorm.DeletedAt
}
