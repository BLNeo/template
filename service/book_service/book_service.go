package book_service

import (
	"template/models"
	"template/models/book"
)

func NewBookService() IBookService {
	return &BookService{}
}

type IBookService interface {
	Create(in *book.AddBookRequest) error
}

type BookService struct{}

func (b *BookService) Create(in *book.AddBookRequest) error {
	insertDate := &models.Book{
		Name:     in.Name,
		StoreNum: in.StoreNum,
		Price:    in.Price,
	}
	return book.NewIBook().Create(insertDate)
}
