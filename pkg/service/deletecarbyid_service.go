package service

import (
	"github.com/RestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func (s *Service) DeleteByCarId(id primitive.ObjectID) error {
	filter := bson.M{"id": id}
	err := s.store.DeleteByCarId(filter)
	if err != nil {
		// Handle the error using the spearetaerror package
		return serror.NotFoundError("Error in finding car")
	}
	return nil

}
