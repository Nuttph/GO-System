package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
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
	// app := fiber.New()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})

	app.Get("/", message)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)
	app.Delete("/books/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Get("test-html", testHTML)

	app.Listen(":8000")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File uploaded successfully!")
}

func testHTML(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello Nuttaphon",
	})
}
