package routes

import (
	"server/handlers"
	"server/internal/middleware"

	"github.com/gin-gonic/gin"
)


func UserRoutes(router *gin.Engine){

	userRoutes := router.Group("/user")
		
	userRoutes.POST("/signup", handler.SignUp)

	userRoutes.POST("/login", handler.Login)


	userRoutes.GET("/data", middleware.AuthMiddleWare(), handler.UserData)

}


