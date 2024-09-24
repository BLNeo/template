package request

// Page .
type Page struct {
	PageNum  int `json:"page_num"  form:"page_num" binding:"gte=1"`
	PageSize int `json:"page_size" form:"page_size" binding:"gte=1,lte=100"`
}
