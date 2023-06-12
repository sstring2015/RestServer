// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

package main

import (
	"github.com/RestServer/pkg/config"
	"github.com/RestServer/pkg/handlers"
	"github.com/RestServer/pkg/middlewares"
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
	h := handlers.New()
	public.POST("/signin", h.SignIn)
	public.POST("/signup", h.SignUp)
	protected := r.Group("/api/cars")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/add", h.InsertCar)
	protected.GET("", h.GetAllCars)
	protected.GET(":id", h.GetAllCars)
	protected.PUT(":id", h.UpdateCarByID)
	protected.DELETE(":id", h.DeleteByCarId)
	
	r.Run()
}
