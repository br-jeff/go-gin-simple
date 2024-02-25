package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string `json:"name"`
	Document string `json:"document"`
}
