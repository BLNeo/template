package book

type AddBookRequest struct {
	Name     string  `json:"name" binding:"required"`
	StoreNum int     `json:"store_num" `
	Price    float64 `json:"price" binding:"required"`
}
