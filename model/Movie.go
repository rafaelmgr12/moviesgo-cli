package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID    int    `gorm:"primaryKey"`
	Title string `json:"title" `
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}
