package main

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/", message)

	app.Listen(":8000")
}

func message(c *fiber.Ctx) error {
	message := Message{
		Message: "Hello world!",
	}
	return c.JSON(message)
}
