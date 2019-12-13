package models

import "github.com/jinzhu/gorm"

type Board struct {
	gorm.Model
	UserId    uint   `json:"UserId"`
	Title     string `json:"Title"`
	DispOrder int    `json:"DispOrder"`
}
