package routes

import (
	"car/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHealth()
	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-username/:username", h.UserByUserName)
	r.POST("/add/user", h.AddUser)
	r.POST("/bind/header", h.BindHeader)
	r.POST("/bind/query", h.BindQuery)
	r.POST("/bind/queryArray", h.BindQueryArray)
	r.POST("/bind/uri/:name", h.BindUri)
	r.POST("/bind/body", h.BindBody)
	r.POST("/bind/form", h.BindBForm)
	r.POST("/bind/file", h.BindFile)
}
