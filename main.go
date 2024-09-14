package main

import (
	"context"
	"log"

	"github.com/SoftwareArch-BackstreetBoys/club-service/application"
	http_server "github.com/SoftwareArch-BackstreetBoys/club-service/http"
	api_gen "github.com/SoftwareArch-BackstreetBoys/club-service/http/gen"
	"github.com/SoftwareArch-BackstreetBoys/club-service/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository := repository.NewRepository(context.Background())

	app := application.NewApplication(repository)

	http := http_server.NewHttp(app)

	fiberApp := fiber.New()

	api_gen.RegisterHandlers(fiberApp, http)

	log.Fatal(fiberApp.Listen(":8080"))
}
