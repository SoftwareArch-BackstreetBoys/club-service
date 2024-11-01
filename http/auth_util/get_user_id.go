package auth_util

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type UserClaims struct {
	jwt.StandardClaims

	Id string `json:"id"`
}

var JWT_SECRET string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
}

func GetUserId(jwtToken string) string {
	parsedToken, _ := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	userClaims := parsedToken.Claims.(*UserClaims)

	return userClaims.Id
}
