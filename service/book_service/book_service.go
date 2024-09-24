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
	List(in *book.ListBookRequest) ([]*book.ListBookRespond, int64, error)
}

type BookService struct{}

func (b *BookService) List(in *book.ListBookRequest) ([]*book.ListBookRespond, int64, error) {
	resp := make([]*book.ListBookRespond, 0)
	date, count, err := book.NewIBook().List(in)
	if err != nil {
		return resp, 0, err
	}
	for _, v := range date {
		resp = append(resp, &book.ListBookRespond{
			Name:     v.Name,
			StoreNum: v.StoreNum,
			Price:    v.Price,
		})
	}
	return resp, count, nil
}

func (b *BookService) Create(in *book.AddBookRequest) error {
	insertDate := &models.Book{
		Name:     in.Name,
		StoreNum: in.StoreNum,
		Price:    in.Price,
	}
	return book.NewIBook().Create(insertDate)
}
