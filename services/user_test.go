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

func Test_User_Get_Paginated_Error(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.UserRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, fiber.StatusInternalServerError, err.Code)
}

func Test_User_Get_Paginated(t *testing.T) {
	paginator := utils.Paginator{}
	repository := mocks.UserRepository{}

	repository.On("FindPaginated", paginator).Return(paginator, nil).Once()

	service := User{Repository: &repository}
	_, err := service.GetPaginated(paginator)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Validate_Login_User_Not_Found(t *testing.T) {
	form := models.Login{}
	form.Email = "admin@email.com"
	form.Password = "12345"

	repository := mocks.UserRepository{}
	repository.On("FindByEmail", form.Email).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "user not found or password not match")
}

func Test_User_Validate_Login_Password_Not_Match(t *testing.T) {
	form := models.Login{}
	form.Email = "admin@email.com"
	form.Password = "12345"

	user := models.User{
		Email:    form.Email,
		Password: utils.EncodePassword("abcde"),
	}

	repository := mocks.UserRepository{}
	repository.On("FindByEmail", form.Email).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "user not found or password not match")
}

func Test_User_Validate_Login_Valid(t *testing.T) {
	form := models.Login{}
	form.Email = "admin@email.com"
	form.Password = "12345"

	user := models.User{
		Email:    form.Email,
		Password: utils.EncodePassword(form.Password),
	}

	repository := mocks.UserRepository{}
	repository.On("FindByEmail", form.Email).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.ValidateLogin(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Get_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", uint64(1)).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Get(uint64(1))

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Get_Found(t *testing.T) {
	user := models.User{
		ID:    1,
		Email: "admin@email.com",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()

	service := User{Repository: &repository}
	result, err := service.Get(user.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, result.Email)
}

func Test_User_Create_Error(t *testing.T) {
	form := models.CreateUser{
		Email:    "user",
		Password: "12345",
	}

	repository := mocks.UserRepository{}
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_User_Create_Success(t *testing.T) {
	form := models.CreateUser{
		Email:    "user",
		Password: "12345",
	}

	repository := mocks.UserRepository{}
	repository.On("Save", mock.Anything).Return(nil).Once()

	service := User{Repository: &repository}
	_, err := service.Create(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Delete_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", uint64(1)).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	err := service.Delete(uint64(1))

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusNotFound)
}

func Test_User_Delete_Error(t *testing.T) {
	user := models.User{
		ID:    1,
		Email: "admin",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()
	repository.On("Delete", &user).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	err := service.Delete(1)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Code, fiber.StatusInternalServerError)
}

func Test_User_Delete_Found(t *testing.T) {
	user := models.User{
		ID:    1,
		Email: "admin",
	}

	repository := mocks.UserRepository{}
	repository.On("Find", user.ID).Return(user, nil).Once()
	repository.On("Delete", &user).Return(nil).Once()

	service := User{Repository: &repository}
	err := service.Delete(user.ID)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}

func Test_User_Update_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", uint64(0)).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Update(models.UpdateUser{})

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrNotFound.Message)
}

func Test_User_Update_Error(t *testing.T) {
	form := models.UpdateUser{
		ID:   1,
		Name: "admin",
	}
	user := models.User{
		ID:   form.ID,
		Name: form.Name,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", &user).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_User_Update_Success(t *testing.T) {
	form := models.UpdateUser{
		ID:   1,
		Name: "admin",
	}
	user := models.User{
		ID:   form.ID,
		Name: form.Name,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", &user).Return(nil).Once()

	service := User{Repository: &repository}
	result, err := service.Update(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
}

func Test_User_Update_Password_Not_Found(t *testing.T) {
	repository := mocks.UserRepository{}
	repository.On("Find", uint64(0)).Return(models.User{}, errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(models.UpdatePassword{})

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrNotFound.Message)
}

func Test_User_Update_Password_Old_Password_Not_Match(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Name:     "admin",
		Password: form.Password,
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), "old password not match")
}

func Test_User_Update_Password_Error(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Name:     "admin",
		Password: utils.EncodePassword(form.OldPassword),
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", mock.Anything).Return(errors.New("")).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Equal(t, err.Error(), fiber.ErrInternalServerError.Message)
}

func Test_User_Update_Password_Success(t *testing.T) {
	form := models.UpdatePassword{
		ID:          1,
		OldPassword: "old",
		Password:    "new",
	}
	user := models.User{
		ID:       form.ID,
		Name:     "admin",
		Password: utils.EncodePassword(form.OldPassword),
	}

	repository := mocks.UserRepository{}
	repository.On("Find", form.ID).Return(user, nil).Once()
	repository.On("Save", mock.Anything).Return(nil).Once()

	service := User{Repository: &repository}
	_, err := service.UpdatePassword(form)

	repository.AssertExpectations(t)

	assert.Nil(t, err)
}
