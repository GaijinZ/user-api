package model

// User represents the user for this application
// swagger:model
type User struct {
	// The ID for this user
	// required: true
	ID int `json:"id" form:"id"`

	// The name for this user
	// required: true
	Firstname string `json:"firstname" form:"firstname"`

	// The lastname for this user
	// required: true
	Lastname string `json:"lastname" form:"lastname"`

	// The email address for this user
	// required: true
	Email string `json:"email" form:"email"`
}
