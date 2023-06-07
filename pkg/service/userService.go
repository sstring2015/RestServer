package service

import (
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/store"
	"gopkg.in/mgo.v2/bson"
)

type UserService interface {
	GetUserByEmail(email string) (models.User, error)
	InsertUser(data models.User) error
}

type DefaultUserService struct {
	store *store.Store
}

func NewUserService(store *store.Store) *DefaultUserService {
	return &DefaultUserService{
		store: store,
	}
}

func (s *DefaultUserService) GetUserByEmail(email string) (models.User, error) {
	filter := bson.M{"email": email}
	return s.store.GetUserByEmail(filter)

}

func (s *DefaultUserService) InsertUser(data models.User) error {
	return s.store.InsertUser(data)
}
