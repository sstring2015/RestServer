package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) DeleteByCarId(c *gin.Context) {
	id := c.Param("id")
	carid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}
	cerr := h.Service.DeleteByCarId(carid)
	if cerr != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
