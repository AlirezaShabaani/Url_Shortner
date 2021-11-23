package model

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Ourl string `json:"ourl"`
	Surl string `json:"surl" gorm:"primarykey"`
}
