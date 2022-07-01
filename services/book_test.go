package services

import (
	"errors"
	"testing"

	mocks "github.com/triagungtio07/kelas_golang/mocks/repositories"
	"github.com/triagungtio07/kelas_golang/models"

	"github.com/triagungtio07/kelas_golang/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Book_Get_Paginated_Error(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.BookRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, err.Code)
}

func Test_Book_Get_Paginated(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.BookRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, nil).Once()

	service := Book{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_Book_Get_Not_Found(t *testing.T) {
	repository := mocks.BookRepository{}
	repository.On("Find", uint64(1)).Return(models.Book{}, errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.Get(uint64(1))

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_Book_Get_Found(t *testing.T) {
	book := models.Book{
		ID:     1,
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Find", book.ID).Return(book, nil).Once()

	service := Book{Repository: &repository}
	result, err := service.Get(book.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, book.Title, result.Title)
}

func Test_Book_Create_Error(t *testing.T) {
	form := models.CreateBook{
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_Book_Create_Duplicate(t *testing.T) {
	form := models.CreateBook{
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Save", mock.Anything).Return(errors.New("Title already exists")).Once()

	service := Book{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)
	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_Book_Create_Success(t *testing.T) {
	form := models.CreateBook{
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Save", mock.Anything).Return(nil).Once()

	service := Book{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_Book_Delete_Not_Found(t *testing.T) {
	repository := mocks.BookRepository{}
	repository.On("Find", uint64(1)).Return(models.Book{}, errors.New("")).Once()

	service := Book{Repository: &repository}
	err := service.Delete(uint64(1))

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_Book_Delete_Error(t *testing.T) {
	book := models.Book{
		ID:     1,
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Find", book.ID).Return(book, nil).Once()
	repository.On("Delete", &book).Return(errors.New("")).Once()

	service := Book{Repository: &repository}
	err := service.Delete(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_Book_Delete_Found(t *testing.T) {
	book := models.Book{
		ID:     1,
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}

	repository := mocks.BookRepository{}
	repository.On("Find", book.ID).Return(book, nil).Once()
	repository.On("Delete", &book).Return(nil).Once()

	service := Book{Repository: &repository}
	err := service.Delete(book.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_Book_Update_Not_Found(t *testing.T) {
	repository := mocks.BookRepository{}
	repository.On("Find", uint64(0)).Return(models.Book{}, errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.Update(models.UpdateBook{})

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrNotFound.Message)
}

func Test_Book_Update_Error(t *testing.T) {
	form := models.UpdateBook{
		ID:     1,
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}
	book := models.Book{
		ID:     form.ID,
		Title:  form.Title,
		Author: form.Author,
		Genre:  form.Genre,
	}

	repository := mocks.BookRepository{}
	repository.On("Find", form.ID).Return(book, nil).Once()
	repository.On("Save", &book).Return(errors.New("")).Once()

	service := Book{Repository: &repository}
	_, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_Book_Update_Success(t *testing.T) {
	form := models.UpdateBook{
		ID:     1,
		Title:  "buku 1",
		Author: "penulis 1",
		Genre:  "genre 1",
	}
	book := models.Book{
		ID:     form.ID,
		Title:  form.Title,
		Author: form.Author,
		Genre:  form.Genre,
	}

	repository := mocks.BookRepository{}
	repository.On("Find", form.ID).Return(book, nil).Once()
	repository.On("Save", &book).Return(nil).Once()

	service := Book{Repository: &repository}
	result, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, book.ID, result.ID)
	assert.Equal(t, book.Title, result.Title)
}
