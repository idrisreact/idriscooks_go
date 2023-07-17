package models

import (
	"time"

	"gorm.io/gorm"
)


type Recipe struct {
	gorm.Model
	Name string `json:"name" gorm:text;not null;default:null`
	Duration int `json:"duration"`
	Serves int `json:"serves"`
	Instructions string `json:"instructions" gorm:text;not null;default:null`
	CreatedAt time.Time `json:"createdAt"`
}

type RecipeResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Duration     int       `json:"duration"`
	Serves       int       `json:"serves"`
	Instructions string    `json:"instructions"`
	CreatedAt    time.Time `json:"createdAt"`
}