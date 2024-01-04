package routes

import (
	"server/handlers"

	"github.com/gin-gonic/gin"
)


func UserRoutes(router *gin.Engine){

	userRoutes := router.Group("/user")
		
	userRoutes.POST("/signup", handler.SignUp)

	userRoutes.POST("/login", handler.Login)

}


