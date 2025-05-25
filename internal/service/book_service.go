package service

import (
	"errors"

	domain "github.com/rakesh-gupta29/portal/internal/domain/book"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook(name string, status domain.Status) (*domain.Book, error) {
	if !domain.ValidStatus(status) {
		return nil, errors.New("invalid book status")
	}

	book := &domain.Book{
		Name:   name,
		Status: status,
	}

	if err := s.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBook(id string) (*domain.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) GetAllBooks() ([]*domain.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) UpdateBook(id string, name string, status domain.Status) (*domain.Book, error) {
	if !domain.ValidStatus(status) {
		return nil, errors.New("invalid book status")
	}

	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	book.Name = name
	book.Status = status

	if err := s.repo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}

func (s *BookService) PublishBook(id string) error {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	book.Status = domain.StatusPublished
	return s.repo.Update(book)
}

func (s *BookService) DraftBook(id string) error {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	book.Status = domain.StatusDraft
	return s.repo.Update(book)
}
