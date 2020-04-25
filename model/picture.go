package model

import "github.com/jinzhu/gorm"

//Picture model
type Picture struct {
	gorm.Model
	GameID int
	URL string
	Size string
}