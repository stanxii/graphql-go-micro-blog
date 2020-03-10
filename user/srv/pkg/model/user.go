package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Account  string `json:"account"`
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   int    `json:"gender"`
	Password string `json:"password"`
}