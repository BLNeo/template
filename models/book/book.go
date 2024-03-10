package book

import (
	"gorm.io/gorm"
	"template/models"
)

func NewIBook() IBook {
	return &Book{
		db: models.GetDb(),
	}
}

type IBook interface {
	Create(in *models.Book) error
}

type Book struct {
	db *gorm.DB
}

func (b *Book) Create(in *models.Book) error {
	return b.db.Create(in).Error
}
