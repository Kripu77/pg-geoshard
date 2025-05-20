package application

import (
	"errors"
	"time"

	"github.com/pg-geoshard/services/library-store/internal/domain"
)

type BookService struct {
	repo domain.BookRepository
}

func NewBookService(repo domain.BookRepository) *BookService {
	return &BookService{repo: repo}
}

type CreateBookDTO struct {
	Title       string    `json:"title" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	ISBN        string    `json:"isbn" validate:"required,isbn"`
	Publisher   string    `json:"publisher" validate:"required"`
	PublishedAt time.Time `json:"published_at" validate:"required"`
	Quantity    int       `json:"quantity" validate:"gte=0"`
}

func (s *BookService) CreateBook(dto CreateBookDTO) (*domain.Book, error) {
	existing, err := s.repo.FindByISBN(dto.ISBN)
	if err == nil && existing != nil {
		return nil, errors.New("book with this ISBN already exists")
	}

	book, err := domain.NewBook(
		dto.Title,
		dto.Author,
		dto.ISBN,
		dto.Publisher,
		dto.PublishedAt,
		dto.Quantity,
	)
	if err != nil {
		return nil, err
	}

	err = s.repo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBook(id uint) (*domain.Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) GetAllBooks() ([]*domain.Book, error) {
	return s.repo.FindAll()
}

func (s *BookService) UpdateBook(id uint, dto CreateBookDTO) (*domain.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	book.Title = dto.Title
	book.Author = dto.Author
	book.ISBN = dto.ISBN
	book.Publisher = dto.Publisher
	book.PublishedAt = dto.PublishedAt
	err = book.UpdateQuantity(dto.Quantity)
	if err != nil {
		return nil, err
	}

	err = s.repo.Update(book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) DeleteBook(id uint) error {
	return s.repo.Delete(id)
}

func (s *BookService) GetBookByISBN(isbn string) (*domain.Book, error) {
	return s.repo.FindByISBN(isbn)
}
