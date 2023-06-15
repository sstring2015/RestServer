package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) InsertCar(data models.Car) error {

	filter := bson.M{"brand": data.Brand, "model": data.Model}
	_, err := s.store.GetCarByModelBrand(filter)
	if err == nil {
		return serror.AlreadyInUse("Cars Details Already there")
	}

	err = s.store.InsertCar(data)
	if err != nil {
		return serror.BadRequestError("Problem adding a car")
	}

	return nil

}
