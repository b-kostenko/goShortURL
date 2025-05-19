package auth

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `grom:"uniqueIndex"`
	Password string
}

type UserInputSchema struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
