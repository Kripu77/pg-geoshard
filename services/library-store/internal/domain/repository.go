package domain

// BookRepository defines the interface for book persistence operations
type BookRepository interface {
	Create(book *Book) error
	FindByID(id uint) (*Book, error)
	FindAll() ([]*Book, error)
	Update(book *Book) error
	Delete(id uint) error
	FindByISBN(isbn string) (*Book, error)
}
