package models

import "time"

type Recette struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	Ingredients  string `json:"ingredients"`
	Photos       string `json:"photos"`
	Instructions string `json:"instructions"`
	Page         string `json:"line"`
	SerialNumber string `json:"serial_number"`
}
