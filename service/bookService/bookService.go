package bookService

import (
	"template/dao"
	"template/models"
)

func NewBookService() IBookService {
	return &BookService{
		Dao: dao.NewDao(),
	}
}

type IBookService interface {
	Create(in *AddBookRequest) error
	List(in *ListBookRequest) ([]*ListBookRespond, int64, error)
}

type BookService struct {
	Dao dao.IDao // 数据库
	// 缓存
}

func (b *BookService) List(in *ListBookRequest) ([]*ListBookRespond, int64, error) {
	resp := make([]*ListBookRespond, 0)
	date, count, err := b.Dao.Book().List(in.PageNum, in.PageSize)
	if err != nil {
		return resp, 0, err
	}
	for _, v := range date {
		resp = append(resp, &ListBookRespond{
			Name:     v.Name,
			StoreNum: v.StoreNum,
			Price:    v.Price,
		})
	}
	return resp, count, nil
}

func (b *BookService) Create(in *AddBookRequest) error {
	insertDate := &models.TableBook{
		Name:     in.Name,
		StoreNum: in.StoreNum,
		Price:    in.Price,
	}
	return b.Dao.Book().Create(insertDate)
}
