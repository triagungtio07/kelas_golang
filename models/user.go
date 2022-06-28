package models

import "time"

type (
	User struct {
		ID        uint64 `gorm:"primaryKey" json:"id"`
		Name      string `gorm:"type:varchar(255);notNull"`
		Email     string `gorm:"type:varchar(255);unique;notNull"`
		Password  string `gorm:"type:varchar(255);notNull"`
		Phone     string `gorm:"type:varchar(255)"`
		Address   string `gorm:"type:varchar(255)"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CreateUser struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"-" validate:"required"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	UpdateUser struct {
		ID      uint64
		Name    string `json:"name" validate:"required"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
	}

	UpdatePassword struct {
		ID          uint64
		OldPassword string `json:"old_password" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	Login struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func (User) TableName() string {
	return "users"
}
