package model

import "gorm.io/gorm"

// User represents the user for this application
// swagger:model
type User struct {
	gorm.Model
	// The ID for this user
	// required: true
	ID int `json:"id" form:"id"`

	// The name for this user
	// required: true
	Firstname string `json:"firstname" form:"firstname"`

	// The lastname for this user
	// required: true
	Lastname string `json:"lastname" form:"lastname"`

	// The unique email address for this user
	// required: true
	Email string `gorm:"unique" json:"email" form:"email"`

	// The password for this user
	// required: true
	Password string `json:"password" form:"password"`

	// The role for this user
	// required: false
	Role string `json:"role" form:"role"`
}

// swagger:model GenericError
type GenericError struct {
	// Response message
	// in: string
	// required: true
	Message string `json:"message"`
}

type Authentication struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
	Role        string `json:"role" form:"role"`
}
