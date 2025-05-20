package domain

import (
	"errors"
	"time"
)

// Book represents the core domain entity for a book in the library
type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	ISBN        string    `json:"isbn" validate:"required,isbn"`
	Publisher   string    `json:"publisher" validate:"required"`
	PublishedAt time.Time `json:"published_at" validate:"required"`
	Quantity    int       `json:"quantity" validate:"gte=0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewBook creates a new Book instance with validation
func NewBook(title, author, isbn, publisher string, publishedAt time.Time, quantity int) (*Book, error) {
	if title == "" || author == "" || isbn == "" || publisher == "" {
		return nil, errors.New("all fields are required")
	}

	if quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}

	return &Book{
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		Publisher:   publisher,
		PublishedAt: publishedAt,
		Quantity:    quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// UpdateQuantity updates the book quantity with validation
func (b *Book) UpdateQuantity(quantity int) error {
	if quantity < 0 {
		return errors.New("quantity cannot be negative")
	}
	b.Quantity = quantity
	b.UpdatedAt = time.Now()
	return nil
}
