package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/gin-gonic/gin"
)

// InsertCar handles the POST /cars endpoint.
// @Summary Insert a new car
// @Description Insert a new car with the provided details
// @Tags Cars
// @Accept json
// @Produce json
// @Param carData body models.Car true "Car details"
// @Success 201 {object} gin.H "success"
// @Router /api/cars/add [post]
func (h *Handler) InsertCar(c *gin.Context) {
	var carData models.Car
	if c.BindJSON(&carData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	err := h.Service.InsertCar(carData)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New Car details registered"})

}
