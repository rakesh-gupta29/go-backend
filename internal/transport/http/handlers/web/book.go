package handlers

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/rakesh-gupta29/portal/internal/domain/book"
	"github.com/rakesh-gupta29/portal/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) HandleListBooks(c *fiber.Ctx) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}

func (h *BookHandler) HandleGetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.service.GetBook(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Render("partials/book-form", fiber.Map{
		"Book": book,
	})
}

func (h *BookHandler) HandleNewBook(c *fiber.Ctx) error {
	return c.Render("partials/book-form", fiber.Map{
		"Book": &domain.Book{},
	})
}

func (h *BookHandler) HandleCreateBook(c *fiber.Ctx) error {
	name := c.FormValue("name")
	statusStr := c.FormValue("status")

	var status domain.Status
	switch statusStr {
	case "published":
		status = domain.StatusPublished
	case "draft":
		status = domain.StatusDraft
	default:
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	if _, err := h.service.CreateBook(name, status); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create book",
		})
	}

	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}

func (h *BookHandler) HandleUpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	name := c.FormValue("name")
	statusStr := c.FormValue("status")

	var status domain.Status
	switch statusStr {
	case "published":
		status = domain.StatusPublished
	case "draft":
		status = domain.StatusDraft
	default:
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	if _, err := h.service.UpdateBook(id, name, status); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update book",
		})
	}

	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}

func (h *BookHandler) HandleDeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteBook(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	// Return the updated list of books
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}

func (h *BookHandler) HandlePublishBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.PublishBook(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	// Return the updated list of books
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}

func (h *BookHandler) HandleDraftBook(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DraftBook(id); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	// Return the updated list of books
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch books",
		})
	}

	return c.Render("pages/books", fiber.Map{
		"Title": "Books",
		"Books": books,
	}, "layouts/base")
}
