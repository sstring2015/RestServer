package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) GetCarByID(id primitive.ObjectID) (models.Car, error) {
	filter := bson.M{"id": id}
	car, err := s.store.GetCarByID(filter)
	if err != nil {
		// Handle the error using the spearetaerror package
		return models.Car{}, serror.NotFoundError("Error in finding car")
	}
	return car, nil

}
