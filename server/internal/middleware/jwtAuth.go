package middleware

import (
	"net/http"
    "server/internal/struct"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)



func AuthMiddleWare() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		tokenString, cookieErr := ctx.Cookie("auth_token")

		if cookieErr != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"ERROR": "INVALID COOKIE"})
			ctx.Abort()
			return
		}

		token, tokenErr := jwt.ParseWithClaims(tokenString, &user_struct.UserDataClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("jwt_secret"), nil
		})

		if tokenErr != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"ERROR": "UNAUTHORIZED"})
			ctx.Abort()
			return
		}

		user, ok := token.Claims.(*user_struct.UserDataClaims)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ERROR": "ERROR JWT STRUCT CLAIMS"})
			ctx.Abort()
			return
		}

		ctx.Set("UserData", user)
		ctx.Next()
	}
}
