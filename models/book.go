package models

type TableBook struct {
	Model
	Name     string  `json:"name" gorm:"column:name"`
	StoreNum int     `json:"store_num" gorm:"column:store_num;type:int4"`
	Price    float64 `json:"price" gorm:"column:price"`
}

func (b *TableBook) TableName() string {
	return "book"
}
