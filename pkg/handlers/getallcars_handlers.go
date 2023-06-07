package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

	cars, err := h.Service.GetAllCars()
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "cars": cars})

}
