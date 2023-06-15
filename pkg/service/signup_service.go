package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) SignUp(data models.User) error {

	filter := bson.M{"email": data.Email}
	_, err := s.store.GetUserByEmail(filter)
	if err == nil {
		return serror.AlreadyInUse("Email is already in use")
	}

	err = s.store.InsertUser(data)
	if err != nil {
		return serror.BadRequestError("Problem creating an account")
	}

	return nil

}
