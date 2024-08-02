package models

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isactive" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
}

// exercise 6
type Company struct {
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name" validate:"required,customName"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Line string `json:"line"`
	Phone string `json:"phone" validate:"required,max=10,min=10,numeric"`
	Business string `json:"business" validate:"required"`
	Website string `json:"website" validate:"required,url,uri,customWeb"`
}