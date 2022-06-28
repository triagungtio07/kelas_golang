package models

import "time"

type (
	Book struct {
		ID        uint64 `gorm:"primaryKey" json:"id"`
		Title     string `gorm:"type:varchar(255);unique;notNull" json:"title"`
		Author    string `gorm:"type:varchar(255);notNull" json:"author"`
		Genre     string `gorm:"type:varchar(255);notNull" json:"genre"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	CreateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Genre  string `json:"genre"`
	}

	UpdateBook struct {
		ID     uint64 `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
		Genre  string `json:"genre"`
	}
)

func (Book) TableName() string {
	return "books"
}
