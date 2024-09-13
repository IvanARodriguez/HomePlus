package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	app.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"message": "test endpoint is working"})
	})

	log.Fatal(app.Listen(":5000"))
}
