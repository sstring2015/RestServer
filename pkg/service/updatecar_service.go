package service

import (
	"fmt"

	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) UpdateCarByID(data models.Car, id primitive.ObjectID) error {

	filter := bson.M{"id": id}

	// Build the update statement
	update := bson.M{"$set": bson.M{"brand": data.Brand}}

	err := s.store.UpdateCarByID(filter, update)
	if err != nil {
		return serror.InternalServerError(fmt.Sprintf("error calling store, %v", err))
	}
	return nil

}
