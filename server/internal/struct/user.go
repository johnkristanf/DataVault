package user_struct

import (
	"github.com/golang-jwt/jwt/v5"
)


type UserDataClaims struct{
    Email   string `json:"email"`
	UserID   string `json:"user_id"`
	Exp      int64  `json:"exp"`
	jwt.RegisteredClaims
}
