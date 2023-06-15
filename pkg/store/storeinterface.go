package store

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type Store interface {
	GetUserByEmail(filter bson.M) (models.User, error)
	InsertUser(data models.User) error
	GetCarByModelBrand(filter bson.M) (models.Car, error)
	GetCarByID(filter bson.M) (models.Car, error)
	InsertCar(data models.Car) error
	GetAllCars(pagination utils.Pagination) ([]models.Car, int64, error)
	UpdateCarByID(filter bson.M, updater bson.M) error
	DeleteByCarID(filter bson.M) error
}
