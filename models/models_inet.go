package models

import "gorm.io/gorm"

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
type Company01 struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,customName"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Line     string `json:"line"`
	Phone    string `json:"phone" validate:"required,max=10,min=10,numeric"`
	Business string `json:"business" validate:"required"`
	Website  string `json:"website" validate:"required,url,uri,customWeb"`
}

type Dogs struct {
	gorm.Model
	DogID int `json:"dog_id"`
	Name string `json:"name"`
}

// exercise 7.0
type DogsRes struct {
	Name string `json:"name"`
	DogID int `json:"dog_id"`
	Type string `json:"type"`
}

type ResultData struct {
	Count int `json:"count"`
	Data []DogsRes `json:"data"`
	Name string `json:"name"`
	Red int `json:"sum_red"`
	Green int `json:"sum_green"`
	Pink int `json:"sum_pink"`
	Nocolor int `json:"sum_nocolor"`
}

// exercise 7.0.1
type Company struct {
	gorm.Model
	Name string	`json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
	Phone string `json:"phone" validate:"required,max=10,min=10,numeric"`
	Website string `json:"website" validate:"required,url,uri"`
	Employee int `json:"employee"`
}