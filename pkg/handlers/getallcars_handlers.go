package handlers

import (
	"net/http"
	"strconv"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAllCarsResponse struct {
    Cars       []models.Car `json:"cars"`
    TotalPages int64        `json:"totalPages"`
}

// GetAllCars retrieves all cars with pagination.
// @Summary Get all cars with pagination
// @Description Retrieve all cars with pagination
// @Tags Cars
// @Accept json
// @Produce json
// @Param _start query integer false "Start index for pagination"
// @Param _end query integer false "End index for pagination"
// @Success 200 {object} GetAllCarsResponse "success"
// @Router /api/cars [get]
func (h *Handler) GetAllCars(c *gin.Context) {

	idParam := c.Query("id")
	if idParam != "" {
		// Retrieve a specific car by ID
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
			return
		}

		car, err := h.Service.GetCarByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve car"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"car": car})
		return
	}
	start, err := strconv.Atoi(c.DefaultQuery("_start", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start value"})
		return
	}

	end, err := strconv.Atoi(c.DefaultQuery("_end", "3"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end value"})
		return
	}
	pagination := utils.Pagination{
		PageNumber: (start / end) + 1,
		PageSize:   end - start,
	}

	cars, total, err := h.Service.GetAllCars(pagination)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "cars": cars, "totalPages": total})

}
