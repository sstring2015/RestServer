package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/store"
	"github.com/RestServer/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	SignIn(email string) (models.User, error)
	SignUp(data models.User) error
	InsertCar(data models.Car) error
	GetAllCars(data utils.Pagination) ([]models.Car, int64, error)
	GetCarByID(id primitive.ObjectID) (car models.Car, err error)
	UpdateCarByID(data models.Car, id primitive.ObjectID) error
}

type Service struct {
	store *store.Store
}

func NewUserService(store *store.Store) *Service {
	return &Service{
		store: store,
	}
}
