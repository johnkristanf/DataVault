package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	
)

var secretKey = []byte("jwt_secret")


func CreateJwtToken(email *string, user_id *string) (string, error){
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"email": email,
			"user_id": user_id,
			"exp" : time.Now().Add(time.Hour * 24).Unix(), 
	    })

	tokenString, err := token.SignedString(secretKey)
		if err != nil {
		  return "", err

		}
	
	return tokenString, nil
}

