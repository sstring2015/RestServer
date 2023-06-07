package models

import "time"

type Cars struct {
	ID    string    `json:"id"`
	Brand string    `json:"brand"`
	Model string    `json:"model"`
	DOM   time.Time `json:"dom"`
	Price float64   `json:"price"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   string `json:"age"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
