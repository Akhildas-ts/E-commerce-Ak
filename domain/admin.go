package domain

import "ak/models"

type TokenAdmin struct {
	Admin models.AdminDetailsResponse
	Token string
}

type Admin struct {
	ID        uint   `json:"id" gorm:"unique;not nul"`
	Firstname string `json:"firstname" gorm:"validate:required"`
	Lastname  string `json:"lastname" gorm:"validate:required"`
	Email     string `json:"email" gorm:"validate:required"`
	Password  string `json:"password" gorm:"validate:required"`
}
