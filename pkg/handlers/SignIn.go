package handlers

import (
	"net/http"

	"github.com/RestServer/pkg/errorutil"
	"github.com/RestServer/pkg/models"
	"github.com/RestServer/pkg/service"
	"github.com/gin-gonic/gin"
)

type MyUser struct {
	models.User
}

type Handler struct {
	Service service.UserService
}

var userService service.UserService

func SetUserService(us service.UserService) {
	userService = us
}

func SignIn(c *gin.Context) {
	var userData models.UserInput
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	_, err := userService.SignIn(userData.Email)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Log in success"})
}

func SignUp(c *gin.Context) {
	var userData models.User
	if c.BindJSON(&userData) != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	err := userService.SignUp(userData)
	if err != nil {
		errorutil.HandleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New user account registered"})

}
