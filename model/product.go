package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	StoreName   string `json:"store_name"`
	Price       int64  `json:"price"`
	Rating      int32  `json:"rating"`
}

//go:generate mockery --name ICompanies
type IProduct interface {
	GetTableName() string
}

func NewProduct() IProduct {
	return &Product{}
}

func (u *Product) GetTableName() string {
	return "product"
}
