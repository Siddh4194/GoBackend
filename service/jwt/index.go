package jwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/siddhant/shetkariBasketGo/models"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	user *models.User
	jwt.RegisteredClaims
}

func GenerateJWT(user *models.User) (string, error) {
	exp := os.Getenv("JWT_SECRET_EXP")
	expDays, err := strconv.Atoi(exp)

	if err != nil {
		return "", err
	}

	expirationTime := time.Now().Add(time.Duration(expDays) * time.Hour)

	claims := &Claims{
		user:user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		// Check if the token is expired
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token has expired")
		}
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is not valid")
	}

	return claims, nil
}
