package models

import (
	"time"
)

type Recette struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string       `json:"name"`
	Descriptions string       `json:"descriptions"`
	Page         string       `json:"line"`
}
