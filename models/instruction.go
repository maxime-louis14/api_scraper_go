package models

import (
	"time"

	"gorm.io/gorm"
)

type Instruction struct {
	gorm.Model
	Number      string `json:"number"`
	Description string `json:"description"`
	RecetteID   uint
	CreatedAt   time.Time
}
