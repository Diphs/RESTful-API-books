package main

import (
	"errors"
	"sync"
)

// Book represents a book entity
type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

// CreateBookRequest represents the request body for creating a book
type CreateBookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

// UpdateBookRequest represents the request body for updating a book
type UpdateBookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
}

// BookStore manages book data in memory using Singleton pattern
type BookStore struct {
	books  map[int]*Book
	nextID int
	mutex  sync.RWMutex
}

var (
	bookStoreInstance *BookStore
	once              sync.Once
)

// GetBookStore returns the singleton instance of BookStore
func GetBookStore() *BookStore {
	once.Do(func() {
		bookStoreInstance = &BookStore{
			books:  make(map[int]*Book),
			nextID: 1,
		}
		// Add some sample data
		bookStoreInstance.initSampleData()
	})
	return bookStoreInstance
}

// initSampleData adds some initial books for testing
func (bs *BookStore) initSampleData() {
	sampleBooks := []*Book{
		{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", PublishedYear: 2015},
		{ID: 2, Title: "Clean Code", Author: "Robert Martin", PublishedYear: 2008},
	}

	for _, book := range sampleBooks {
		bs.books[book.ID] = book
		bs.nextID = book.ID + 1
	}
}

// GetAll returns all books
func (bs *BookStore) GetAll() []*Book {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	books := make([]*Book, 0, len(bs.books))
	for _, book := range bs.books {
		books = append(books, book)
	}
	return books
}

// GetByID returns a book by its ID
func (bs *BookStore) GetByID(id int) (*Book, error) {
	bs.mutex.RLock()
	defer bs.mutex.RUnlock()

	book, exists := bs.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

// Create adds a new book to the store
func (bs *BookStore) Create(req CreateBookRequest) *Book {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	book := &Book{
		ID:            bs.nextID,
		Title:         req.Title,
		Author:        req.Author,
		PublishedYear: req.PublishedYear,
	}

	bs.books[bs.nextID] = book
	bs.nextID++

	return book
}

// Update modifies an existing book
func (bs *BookStore) Update(id int, req UpdateBookRequest) (*Book, error) {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	book, exists := bs.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	book.Title = req.Title
	book.Author = req.Author
	book.PublishedYear = req.PublishedYear

	return book, nil
}

// Delete removes a book from the store
func (bs *BookStore) Delete(id int) error {
	bs.mutex.Lock()
	defer bs.mutex.Unlock()

	_, exists := bs.books[id]
	if !exists {
		return errors.New("book not found")
	}

	delete(bs.books, id)
	return nil
}
