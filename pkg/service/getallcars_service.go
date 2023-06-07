package service

import (
	"fmt"

	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/serror"
	"github.com/RestServer/pkg/utils"
)

func (s *Service) GetAllCars(pagination utils.Pagination) ([]models.Car, int64, error) {

	cars, total, err := s.store.GetAllCars(pagination)
	if err != nil {
		fmt.Print(err)
		return nil, 0, serror.InternalServerError("Failed to Retrieve the cars")
	}

	return cars, total, nil

}
