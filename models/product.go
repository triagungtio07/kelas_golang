package models

type Product struct {
	ProductId   int    `gorm:"type:int(11);notNull"`
	ProductName string `gorm:"type:varchar(255);notNull"`
	ProductType string `gorm:"type:varchar(255);notNull"`
}

func (Product) TableName() string {
	return "product"
}
