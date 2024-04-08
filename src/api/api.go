package api

import (
	"car/api/middlewares"
	"car/api/routes"
	"car/api/validations"
	"car/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			return
		}
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddlewares())
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		routes.Health(v1.Group("/health"))
		routes.TestRouter(v1.Group("/test"))

	}
	err := r.Run(fmt.Sprintf(":%v", cfg.Server.Port))
	if err != nil {
		return
	}
}
