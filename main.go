package main

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Author model
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Category model
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Book model
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      Author    `json:"author"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
}

var (
	authors = []Author{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Jane"},
		{ID: 3, Name: "Joe"},
	}

	categories = []Category{
		{ID: 1, Name: "Fiction"},
		{ID: 2, Name: "Non-Fiction"},
		{ID: 3, Name: "Thriller"},
	}

	books = []Book{
		{ID: 1, Title: "Book 1", Author: authors[0], Category: categories[0], Description: "Description 1", PublishedAt: time.Now()},
		{ID: 2, Title: "Book 2", Author: authors[1], Category: categories[1], Description: "Description 2", PublishedAt: time.Now()},
		{ID: 3, Title: "Book 3", Author: authors[2], Category: categories[2], Description: "Description 3", PublishedAt: time.Now()},
		{ID: 4, Title: "Book 4", Author: authors[0], Category: categories[0], Description: "Description 4", PublishedAt: time.Now()},
		{ID: 5, Title: "Book 5", Author: authors[1], Category: categories[1], Description: "Description 5", PublishedAt: time.Now()},
	}
)

// API

// Get all authors
func getAuthors(c *fiber.Ctx) error {
	return c.JSON(authors)
}

// Get all categories
func getCategories(c *fiber.Ctx) error {
	return c.JSON(categories)
}

// Get all books
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}
	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
}

// Create a new book

func createBook(c *fiber.Ctx) error {
	var newBook Book
	if err := c.BodyParser(&newBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}
	newBook.ID = len(books) + 1
	books = append(books, newBook)
	return c.JSON(newBook)
}

// Update an existing book by ID
// Update an existing book by ID
func updateBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	var updatedBook Book
	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request body"})
	}

	// Find the index of the book with the given ID
	index := -1
	for i, book := range books {
		if book.ID == id {
			index = i
			break
		}
	}

	// If book with given ID not found, return not found error
	if index == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
	}

	// Retain the original ID of the book being updated
	updatedBook.ID = books[index].ID

	// Update the book in the slice
	books[index] = updatedBook

	return c.JSON(updatedBook)
}

// Delete a book by ID
func deleteBookByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid book ID"})
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(fiber.Map{"message": "Book deleted successfully"})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Book not found"})
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Routes
	app.Get("/authors", getAuthors)
	app.Get("/categories", getCategories)
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBookByID)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBookByID)
	app.Delete("/books/:id", deleteBookByID)

	// Start server
	app.Listen(":3000")
}
