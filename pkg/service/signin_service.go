package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"gopkg.in/mgo.v2/bson"
)

func (s *Service) SignIn(email string) (models.User, error) {
	filter := bson.M{"email": email}
	user, err := s.store.GetUserByEmail(filter)
	if err != nil {
		// Handle the error using the spearetaerror package
		return models.User{}, serror.NotFoundError("Error in finding user")
	}
	if user.Email == "" {
		return models.User{}, serror.NotFoundError("User not found")
	}
	return user, nil

}
