package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize the book store (Singleton pattern)
	store := GetBookStore()
	
	// Initialize handlers with the store
	bookHandler := NewBookHandler(store)
	
	// Create router
	r := chi.NewRouter()
	
	// Add built-in middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	
	// Add custom logger middleware
	r.Use(LoggerMiddleware)
	
	// Setup routes
	setupRoutes(r, bookHandler)
	
	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// setupRoutes configures all API routes
func setupRoutes(r chi.Router, handler *BookHandler) {
	r.Route("/books", func(r chi.Router) {
		r.Get("/", handler.GetAllBooks)
		r.Post("/", handler.CreateBook)
		r.Get("/{id}", handler.GetBookByID)
		r.Put("/{id}", handler.UpdateBook)
		r.Delete("/{id}", handler.DeleteBook)
	})
}
