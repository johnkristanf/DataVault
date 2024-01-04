package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"server/internal"
	"github.com/rs/cors"
)


func main(){

	router := gin.Default()	

	routes.UserRoutes(router)
	
	fmt.Println("hi")


	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	RequestHandler := router

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:2000"},

		AllowedMethods: []string{
			http.MethodGet, 
			http.MethodPost, 
			http.MethodDelete, 
			http.MethodPut,
		},

		AllowCredentials: true,
		AllowedHeaders: []string{
			"Access-Control-Allow-Credentials", 
			"Access-Control-Allow-Headers",
			"Content-Type",
		},
	})

	handler := cors.Handler(RequestHandler)


	if err := http.ListenAndServe(":900", handler); err != nil {
		panic(err)
	}
}
