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

	Id       string `json:"id"`
	FullName string `json:"fullName"`
}

var JWT_SECRET string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWT_SECRET = os.Getenv("JWT_SECRET")
}

func GetUserFromJWTToken(jwtToken string) (*UserClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		// comment this "if" if you want to test with expired token
		return nil, errors.New("invalid token")
	}

	userClaims := parsedToken.Claims.(*UserClaims)

	return userClaims, nil
}

func GetUserFromFiberContext(c *fiber.Ctx) (*UserClaims, error) {
	jwtToken := c.Cookies("jwt")
	if jwtToken == "" {
		return nil, errors.New("jwt token not found")
	}

	return GetUserFromJWTToken(jwtToken)
}
