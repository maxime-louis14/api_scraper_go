package models

import (
	"time"

	"gorm.io/gorm"
)

type Instruction struct {
	gorm.Model
	Number      string `json:"number"`
	Description string `json:"description" gorm:"type:text"`
	RecetteID   uint
	CreatedAt   time.Time
}
