package dao

import "template/dao/book"

type Dao struct{}

type IDao interface {
	Book() book.IBook
}

func NewDao() IDao {
	return &Dao{}
}

func (d *Dao) Book() book.IBook {
	return book.NewIBook()
}
