package model

import "github.com/jinzhu/gorm"

//Category model
type Category struct {
	gorm.Model
	Title string
	Description string
	Games []Game
}