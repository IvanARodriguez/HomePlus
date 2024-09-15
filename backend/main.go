package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func main() {

	godotenv.Load(".env")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(csrf.New())

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"message": "test endpoint is working"})
	})

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("Could not get PORT environment variable")
	}

	log.Fatal(app.Listen(":" + portString))
}
