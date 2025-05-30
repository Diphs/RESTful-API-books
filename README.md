# Book Management API

A RESTful API for managing books built with Golang and go-chi framework.

## Features

- Full CRUD operations for books
- In-memory storage with thread-safe operations
- Custom logger middleware
- Singleton pattern implementation
- DRY principle adherence
- JSON request/response handling
- Comprehensive error handling and validation

## API Endpoints

### Get All Books
```
GET /books
```

### Get Book by ID
```
GET /books/{id}
```

### Create New Book
```
POST /books
Content-Type: application/json

{
    "title": "Book Title",
    "author": "Author Name",
    "published_year": 2023
}
```

### Update Book
```
PUT /books/{id}
Content-Type: application/json

{
    "title": "Updated Title",
    "author": "Updated Author",
    "published_year": 2024
}
```

### Delete Book
```
DELETE /books/{id}
```

## Installation and Running

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run .
   ```

The server will start on `http://localhost:8080`

## Project Structure

```
.
├── main.go          # Application entry point and routing setup
├── models.go        # Data models and business logic
├── handlers.go      # HTTP request handlers
├── middleware.go    # Custom middleware implementations
├── go.mod          # Go module dependencies
├── start.txt       # Project start marker
├── end.txt         # Project completion marker
└── README.md       # Project documentation
```

## Testing

You can test the API using curl, Postman, or any HTTP client:

```bash
# Get all books
curl http://localhost:8080/books

# Get specific book
curl http://localhost:8080/books/1

# Create new book
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"New Book","author":"Author Name","published_year":2023}'

# Update book
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Book","author":"Updated Author","published_year":2024}'

# Delete book
curl -X DELETE http://localhost:8080/books/1
```

## Technical Implementation

- **Singleton Pattern**: BookStore uses singleton pattern to ensure single instance across the application
- **DRY Principle**: Common functionality like JSON response handling and validation are abstracted into reusable methods
- **Thread Safety**: All operations on the book store are protected with mutex locks
- **Middleware**: Custom logger middleware tracks all requests with method, path, and response time
- **Error Handling**: Comprehensive error handling with appropriate HTTP status codes
- **Validation**: Input validation for all create and update operations