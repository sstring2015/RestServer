package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateCarByID handles the PUT /cars/{id} endpoint.
// @Summary Update a car by ID
// @Description Update an existing car with the provided details
// @Tags Cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Param carData body models.Car true "Car details"
// @Success 200 {string} string "Updated successfully"
// @Router /api/cars/{id} [put]
func (h *Handler) UpdateCarByID(c *gin.Context) {
	id := c.Param("id")
	carid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}
	var carData models.Car
	if c.BindJSON(&carData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}
	cerr := h.Service.UpdateCarByID(carData, carid)
	if cerr != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}
