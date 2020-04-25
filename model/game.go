package model

import "github.com/jinzhu/gorm"

//Game model
type Game struct {
	gorm.Model
	Question string
	Answer string
	CategoryID int
	Pictures []Picture
}