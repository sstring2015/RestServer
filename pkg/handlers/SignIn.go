package handlers

import (
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
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	user, err := userService.GetUserByEmail(userData.Email)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	if user.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Log in success"})
}

func SignUp(c *gin.Context) {
	var userData models.User
	if c.BindJSON(&userData) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	_, err := userService.GetUserByEmail(userData.Email)
	if err == nil {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err = userService.InsertUser(userData)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(201, gin.H{"message": "New user account registered"})
}
