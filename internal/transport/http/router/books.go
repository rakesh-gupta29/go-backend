package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rakesh-gupta29/portal/internal/repository/json"
	"github.com/rakesh-gupta29/portal/internal/service"
	handlers "github.com/rakesh-gupta29/portal/internal/transport/http/handlers/web"
)

func mountBookRoutes(app *fiber.App) *fiber.App {
	// Create repository and service instances
	bookRepo := json.NewBookRepository()
	bookService := service.NewBookService(bookRepo)

	// Create a single handler instance to reuse
	bookHandler := handlers.NewBookHandler(bookService)

	books := app.Group("/books")

	// List and create books
	books.Get("/", bookHandler.HandleListBooks)
	books.Get("/new", bookHandler.HandleNewBook)
	books.Post("/", bookHandler.HandleCreateBook)

	// Individual book operations
	books.Get("/:id", bookHandler.HandleGetBook)
	books.Put("/:id", bookHandler.HandleUpdateBook)
	books.Delete("/:id", bookHandler.HandleDeleteBook)

	// Book status changes
	books.Patch("/:id/publish", bookHandler.HandlePublishBook)
	books.Patch("/:id/draft", bookHandler.HandleDraftBook)

	return app
}
