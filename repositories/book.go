package repositories

import (
	"errors"

	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/utils"
	"gorm.io/gorm"
)

type BookRepository interface {
	Find(id uint64) (models.Book, error)
	FindByTitle(title string) (models.Book, error)
	FindPaginated(paginator utils.Paginator) (utils.Paginator, error)
	Save(book *models.Book) error
	Delete(book *models.Book) error
}

type book struct {
	storage *gorm.DB
}

func NewBookRepository(storage *gorm.DB) BookRepository {
	storage.AutoMigrate(&models.Book{})
	return book{storage: storage}
}

func (r book) Find(id uint64) (models.Book, error) {
	book := models.Book{}
	err := r.storage.First(&book, "id = ?", id).Error
	if book.ID == 0 {
		return book, errors.New("book not found")
	}
	return book, err
}

func (r book) FindByTitle(title string) (models.Book, error) {
	book := models.Book{}
	err := r.storage.First(&book, "title = ?", title).Error

	if book.ID == 0 {
		return book, errors.New("book not found")
	}

	return book, err
}

func (r book) FindPaginated(paginator utils.Paginator) (utils.Paginator, error) {
	books := []models.Book{}
	result := r.storage.Scopes(paginator.Paginate(&books, r.storage)).Find(&books)

	return paginator, result.Error
}

func (r book) Save(book *models.Book) error {
	return r.storage.Save(book).Error
}

func (r book) Delete(book *models.Book) error {
	return r.storage.Delete(book).Error
}
