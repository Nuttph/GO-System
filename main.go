package main

import (
	"strconv"

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

var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})

	app.Get("/", message)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Listen(":8000")
}

func message(c *fiber.Ctx) error {
	message := Message{
		Message: "Hello world!",
	}
	return c.JSON(message)
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		// return c.SendString(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, b := range books {
		if b.ID == bookId {
			return c.JSON(b)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(Message{
		Message: "Book not found!",
	})
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	// c.BodyParser(book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, b := range books {
		if b.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(Message{
		Message: "Book not found!",
	})
}

func deleteBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, b := range books {
		if b.ID == bookId {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}
