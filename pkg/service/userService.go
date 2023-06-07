package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"

	"github.com/RestServer/pkg/store"
	"gopkg.in/mgo.v2/bson"
)

type UserService interface {
	SignIn(email string) (models.User, error)
	SignUp(data models.User) error
}

type DefaultUserService struct {
	store *store.Store
}

func NewUserService(store *store.Store) *DefaultUserService {
	return &DefaultUserService{
		store: store,
	}
}

func (s *DefaultUserService) SignIn(email string) (models.User, error) {
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

func (s *DefaultUserService) SignUp(data models.User) error {

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
