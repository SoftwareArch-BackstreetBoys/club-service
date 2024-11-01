package auth_util

import (
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

func GetUserId(jwtToken string) (string, error) {
	parsedToken, _ := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if !parsedToken.Valid {
		// comment this "if" if you want to test with expired token
		return "", errors.New("invalid token")
	}

	userClaims := parsedToken.Claims.(*UserClaims)

	return userClaims.Id, nil
}

func GetUserIdFromFiberContext(c *fiber.Ctx) (string, error) {
	jwtToken := c.Cookies("jwt")
	if jwtToken == "" {
		return "", errors.New("jwt token not found")
	}

	return GetUserId(jwtToken)
}
