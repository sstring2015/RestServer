package service

import (
	"fmt"

	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
)

func (s *Service) GetAllCars() ([]models.Car, error) {

	cars, err := s.store.GetAllCars()
	if err != nil {
		fmt.Print(err)
		return nil, serror.InternalServerError("Failed to Retrieve the cars")
	}

	return cars, nil

}
