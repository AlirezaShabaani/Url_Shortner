package model

import "gorm.io/gorm"

type Data struct {
	gorm.DB
	Ourl string `json:"ourl"`
	Surl string `json:"surl"`
}
