package main

import (
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Message string `json:"message"`
}

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var book []Book

func main() {
	app := fiber.New()

	book = append(book, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	book = append(book, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	book = append(book, Book{ID: 3, Title: "Book 3", Author: "Author 3"})
	book = append(book, Book{ID: 4, Title: "Book 4", Author: "Author 4"})
	book = append(book, Book{ID: 5, Title: "Book 5", Author: "Author 5"})

	app.Get("/", message)
	app.Get("/books", GetBooks)
	app.Get("/books/:id", getBook)
	app.Listen(":8000")
}

func message(c *fiber.Ctx) error {
	message := Message{
		Message: "Hello world!",
	}
	return c.JSON(message)
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(book)
}

func getBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	for _, b := range book {
		if b.ID == id {
			return c.JSON(b)
		}
	}

	return c.Status(404).SendString("Book not found")

}
