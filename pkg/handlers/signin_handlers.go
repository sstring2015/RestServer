package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/utils"
	"github.com/gin-gonic/gin"
)

// SignIn godoc
// @Summary      Create an Item
// @Description  Sign in user with email and generate token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user body models.UserInput true "User input data"
// @Success 200 {string} string "User Logged In"
// @Router       /api/signin [post]
func (h *Handler) SignIn(c *gin.Context) {
	var userData models.UserInput
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	user, err := h.Service.SignIn(userData.Email)
	token, _err := utils.GenerateToken(user.ID.String())
	if _err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log in success", "token": token})
}
