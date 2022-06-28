package repositories

import (
	"errors"

	"github.com/triagungtio07/kelas_golang/models"
	"github.com/triagungtio07/kelas_golang/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(id uint64) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindPaginated(paginator utils.Paginator) (utils.Paginator, error)
	Save(user *models.User) error
	Delete(user *models.User) error
}

type user struct {
	storage *gorm.DB
}

func NewUserRepository(storage *gorm.DB) UserRepository {
	storage.AutoMigrate(&models.User{})
	return user{storage: storage}
}

func (r user) Find(id uint64) (models.User, error) {
	user := models.User{}
	err := r.storage.First(&user, "id = ?", id).Error
	if user.ID == 0 {
		return user, errors.New("user not found")
	}
	return user, err
}

func (r user) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.storage.First(&user, "email = ?", email).Error

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, err
}

func (r user) FindPaginated(paginator utils.Paginator) (utils.Paginator, error) {
	users := []models.User{}
	result := r.storage.Scopes(paginator.Paginate(&users, r.storage)).Find(&users)

	return paginator, result.Error
}

func (r user) Save(user *models.User) error {
	return r.storage.Save(user).Error
}

func (r user) Delete(user *models.User) error {
	return r.storage.Delete(user).Error
}
