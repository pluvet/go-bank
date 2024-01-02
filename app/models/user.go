package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	id       int     `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Account  Account `gorm:"foreignKey:id"`
}

type Account struct {
	gorm.Model
	id      int     `json:"id" gorm:"primary_key"`
	user_id int     `gorm:"foreignKey:id"`
	Balance float32 `json:"balance"`
}
