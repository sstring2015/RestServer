package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Car struct {
	ID    primitive.ObjectID `json:"id"`
	Brand string             `json:"brand"`
	Model string             `json:"model"`
	DOM   time.Time          `json:"dom"`
	Price float64            `json:"price"`
}

type User struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"name"`
	Email string             `json:"email"`
	Age   string             `json:"age"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
