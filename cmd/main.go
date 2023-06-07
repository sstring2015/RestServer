package main

import (
	"github.com/RestServer/pkg/config"
	"github.com/RestServer/pkg/handlers"
	"github.com/RestServer/pkg/service"
	"github.com/RestServer/pkg/store"
	"github.com/gin-gonic/gin"
)

func init() {
	configFilePath := "../config/master.yaml"
	cfg, err := config.ReadConfig(configFilePath)
	if err != nil {
		panic(err)
	}

	storeInstance := store.GetStore(cfg.Database)
	userService := service.NewUserService(storeInstance)
	handlers.SetUserService(userService)
}

func main() {
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/signin", handlers.SignIn)
	public.POST("/signup", handlers.SignUp)
	r.Run()
}