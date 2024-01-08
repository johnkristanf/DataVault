package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"server/internal"
	"github.com/rs/cors"
	"fmt"
)


func main(){

	router := gin.Default()	

	routes.UserRoutes(router)

	fmt.Println("jake bag o ni")
	

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	RequestHandler := router

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:2000"},

		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},

		AllowCredentials: true,

		AllowedHeaders: []string{
			"Access-Control-Allow-Credentials", 
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Origin",
			"Cookie",
		},
	})

	handler := cors.Handler(RequestHandler)


	if err := http.ListenAndServe(":900", handler); err != nil {
		panic(err)
	}
}
