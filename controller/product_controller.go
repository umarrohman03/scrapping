package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/umarrohman03/scrapping/usecase"
)

type ProductController struct {
	ProductUseCase usecase.IProductUseCase
}

//go:generate mockery --name IproductController
type IProductController interface {
	AddProduct(ctx echo.Context) error
}

func NewProductController(productUseCase usecase.IProductUseCase) IProductController {
	return &ProductController{
		ProductUseCase: productUseCase,
	}
}

// AddProduct ...
func (r ProductController) AddProduct(ctx echo.Context) error {
	fmt.Println("start product controller")
	r.ProductUseCase.AddProduct()
	return nil
}
