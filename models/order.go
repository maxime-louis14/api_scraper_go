package models

import (
	"time"
)

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	RecetteRefer uint    `json:"Recette_id"`
	Recette      Recette `gorm:"foreignKey:RecetteRefer"`
	UserRefer    uint    `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
