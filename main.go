package main

import (
	"context"
	"log"
	"os"

	"github.com/SoftwareArch-BackstreetBoys/club-service/application"
	http_server "github.com/SoftwareArch-BackstreetBoys/club-service/http"
	api_gen "github.com/SoftwareArch-BackstreetBoys/club-service/http/gen"
	"github.com/SoftwareArch-BackstreetBoys/club-service/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Enable CORS for all routes, allowing localhost:5555 as an origin
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTENT_ROUTE"), // Allow your frontend's origin
		AllowMethods: "GET,POST,PUT,DELETE",       // You can specify the methods allowed
	}))

	api_gen.RegisterHandlers(fiberApp, http)

	log.Fatal(fiberApp.Listen(":8080"))
}
