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
	Create(in *models.TableBook) error
	List(pageNum, pageSize int) ([]*models.TableBook, int64, error)
	Get(id int) (*models.TableBook, error)
}

type Book struct {
	db *gorm.DB
}

func (b *Book) Get(id int) (*models.TableBook, error) {
	info := new(models.TableBook)
	err := b.db.Where("id=?", id).First(info).Error
	return info, err
}

func (b *Book) List(pageNum, pageSize int) ([]*models.TableBook, int64, error) {
	date := make([]*models.TableBook, 0)
	var count int64
	err := b.db.Model(&Book{}).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&date).Count(&count).Error
	return date, count, err
}

func (b *Book) Create(in *models.TableBook) error {
	return b.db.Create(in).Error
}
