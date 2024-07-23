package server

import (
	"gin-gonic/controllers"
	"gin-gonic/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/auth")
	users.Use(middlewares.Logger())
	users.GET("/register", controllers.RegisterAccount)

}
