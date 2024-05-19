package utilities

import (
	"errors"
	"fmt"
	"log"
	"online-store/constants"
	"online-store/models"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

var jwtKey = []byte("4Q1S3CR3TK3Y")

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return
		}
		log.Print("ada error", err)
		return
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func ParseJwtToken(tokenStr string) (*models.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok || !token.Valid {
		return &models.JWTClaim{}, fmt.Errorf("failed to claim token: %v", err)
	}

	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(Auth *models.GenerateJWT) (tokenString string, expiredTime string, err error) {
	expirationTime := time.Now().Add(time.Duration(900) * time.Second)
	claims := &models.JWTClaim{
		ID:    Auth.UserId,
		Name:  Auth.Name,
		Email: Auth.Email,
		Role:  Auth.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	expiredTime = time.Now().Add(time.Duration(900) * time.Second).Format(constants.DateFromStd)
	return
}
