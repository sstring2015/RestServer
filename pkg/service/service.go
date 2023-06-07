package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/store"
)

type UserService interface {
	SignIn(email string) (models.User, error)
	SignUp(data models.User) error
	InsertCar(data models.Car) error
}

type Service struct {
	store *store.Store
}

func NewUserService(store *store.Store) *Service {
	return &Service{
		store: store,
	}
}
