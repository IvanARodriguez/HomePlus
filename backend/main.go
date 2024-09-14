package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"message": "test endpoint is working"})
	})

	log.Fatal(app.Listen(":5000"))
}
