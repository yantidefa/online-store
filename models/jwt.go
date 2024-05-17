package models

import "gopkg.in/dgrijalva/jwt-go.v3"

type JWTClaim struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	jwt.StandardClaims
}

type GenerateJWT struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
