package domain

type Status string

const (
	StatusPublished Status = "published"
	StatusDraft     Status = "draft"
)

type Book struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

func ValidStatus(s Status) bool {
	return s == StatusPublished || s == StatusDraft
}

func IsBookPublished(status Status) bool {
	return status == StatusPublished
}

type BookRepository interface {
	Create(book *Book) error
	GetByID(id string) (*Book, error)
	GetAll() ([]*Book, error)
	Update(book *Book) error
	Delete(id string) error
}
