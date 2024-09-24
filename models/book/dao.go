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
	List(in *ListBookRequest) ([]*models.Book, int64, error)
	Get(in *GetBookInfoRequest) (*models.Book, error)
}

type Book struct {
	db *gorm.DB
}

func (b *Book) Get(in *GetBookInfoRequest) (*models.Book, error) {
	info := new(models.Book)
	err := b.db.Where("name=?", in.Name).First(in).Error
	return info, err
}

func (b *Book) List(in *ListBookRequest) ([]*models.Book, int64, error) {
	date := make([]*models.Book, 0)
	var count int64
	err := b.db.Model(&Book{}).Offset((in.PageNum - 1) * in.PageSize).Limit(in.PageSize).Find(&date).Count(&count).Error
	return date, count, err
}

func (b *Book) Create(in *models.Book) error {
	return b.db.Create(in).Error
}
