package models

import (
	"gorm.io/gorm"
)

type User struct {
	id_u     int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func init() {
	confiig.Connect()
}
