package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// BookHandler handles HTTP requests for book operations
type BookHandler struct {
	store *BookStore
}

// NewBookHandler creates a new BookHandler instance
func NewBookHandler(store *BookStore) *BookHandler {
	return &BookHandler{store: store}
}

// GetAllBooks handles GET /books - returns all books
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.store.GetAll()
	h.sendJSONResponse(w, http.StatusOK, books)
}

// GetBookByID handles GET /books/{id} - returns a specific book
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseIDFromURL(r)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := h.store.GetByID(id)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}

	h.sendJSONResponse(w, http.StatusOK, book)
}

// CreateBook handles POST /books - creates a new book
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var req CreateBookRequest
	if err := h.parseJSONRequest(r, &req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.validateCreateBookRequest(req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	book := h.store.Create(req)
	h.sendJSONResponse(w, http.StatusCreated, book)
}

// UpdateBook handles PUT /books/{id} - updates an existing book
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseIDFromURL(r)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var req UpdateBookRequest
	if err := h.parseJSONRequest(r, &req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := h.validateUpdateBookRequest(req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	book, err := h.store.Update(id, req)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}

	h.sendJSONResponse(w, http.StatusOK, book)
}

// DeleteBook handles DELETE /books/{id} - deletes a book
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := h.parseIDFromURL(r)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.store.Delete(id); err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}

	h.sendJSONResponse(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}

// parseIDFromURL extracts and validates ID parameter from URL
func (h *BookHandler) parseIDFromURL(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	return strconv.Atoi(idStr)
}

// parseJSONRequest parses JSON request body into the provided struct
func (h *BookHandler) parseJSONRequest(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

// validateCreateBookRequest validates the create book request
func (h *BookHandler) validateCreateBookRequest(req CreateBookRequest) error {
	if req.Title == "" {
		return errors.New("title is required")
	}
	if req.Author == "" {
		return errors.New("author is required")
	}
	if req.PublishedYear <= 0 {
		return errors.New("published year must be positive")
	}
	return nil
}

// validateUpdateBookRequest validates the update book request
func (h *BookHandler) validateUpdateBookRequest(req UpdateBookRequest) error {
	if req.Title == "" {
		return errors.New("title is required")
	}
	if req.Author == "" {
		return errors.New("author is required")
	}
	if req.PublishedYear <= 0 {
		return errors.New("published year must be positive")
	}
	return nil
}

// sendJSONResponse sends a JSON response with the given status code and data
func (h *BookHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// sendErrorResponse sends an error response in JSON format
func (h *BookHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResp := map[string]string{"error": message}
	h.sendJSONResponse(w, statusCode, errorResp)
}
