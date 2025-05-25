package json

import (
	"errors"
	"strconv"
	"sync"

	domain "github.com/rakesh-gupta29/portal/internal/domain/book"
)

type bookRepository struct {
	books  []*domain.Book
	mu     sync.RWMutex
	nextID int
}

// NewBookRepository creates a new instance of bookRepository
func NewBookRepository() domain.BookRepository {
	return &bookRepository{
		books:  make([]*domain.Book, 0),
		nextID: 1,
	}
}

// Create adds a new book to the repository
func (r *bookRepository) Create(book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !domain.ValidStatus(book.Status) {
		return errors.New("invalid book status")
	}

	book.ID = r.nextID
	r.nextID++
	r.books = append(r.books, book)
	return nil
}

// GetByID retrieves a book by its ID
func (r *bookRepository) GetByID(id string) (*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid book ID")
	}

	for _, book := range r.books {
		if book.ID == idInt {
			return book, nil
		}
	}
	return nil, errors.New("book not found")
}

// GetAll retrieves all books
func (r *bookRepository) GetAll() ([]*domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	books := make([]*domain.Book, len(r.books))
	copy(books, r.books)
	return books, nil
}

// Update modifies an existing book
func (r *bookRepository) Update(book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !domain.ValidStatus(book.Status) {
		return errors.New("invalid book status")
	}

	for i, b := range r.books {
		if b.ID == book.ID {
			r.books[i] = book
			return nil
		}
	}
	return errors.New("book not found")
}

// Delete removes a book by its ID
func (r *bookRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid book ID")
	}

	for i, book := range r.books {
		if book.ID == idInt {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}
