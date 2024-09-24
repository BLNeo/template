package book

import "template/middleware/request"

type AddBookRequest struct {
	Name     string  `json:"name" binding:"required"`
	StoreNum int     `json:"store_num" `
	Price    float64 `json:"price" binding:"required"`
}

type GetBookInfoRequest struct {
	Name string `json:"name" binding:"required"`
}

type ListBookRequest struct {
	request.Page `form:",inline" json:",inline"`
}

type ListBookRespond struct {
	Name     string  `json:"name" `
	StoreNum int     `json:"store_num"`
	Price    float64 `json:"price" `
}
