package router

import (
	"chant/middlewares"
	"chant/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/api/v1/login", service.Login) // User login
	r.POST("/api/v1/send/code", service.SendCode)

	auth := r.Group("/api/v1/u", middlewares.AuthCheck()) // Token check
	auth.GET("/user/detail", service.UserDetail)          // User detail

	return r
}
