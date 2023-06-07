package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var userData models.User
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	err := h.Service.SignUp(userData)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New user account registered"})

}
