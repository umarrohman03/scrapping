package usecase

import (
	"context"
	"fmt"
	models "github.com/umarrohman03/scrapping/model"
	"github.com/umarrohman03/scrapping/repositories"
	"log"
)

type ProductUseCase struct {
	ProductRepo repositories.IProductsRepository
}

type IProductUseCase interface {
	AddProduct() *models.Product
	CSVProduct() error
}

func NewProductUseCase(productRepo repositories.IProductsRepository) *ProductUseCase {
	return &ProductUseCase{
		ProductRepo: productRepo,
	}
}

// AddProduct ...
func (r *ProductUseCase) AddProduct() *models.Product {

	//get product data from tokopedia
	dataProduct, err := r.ProductRepo.GetProduct(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dataProduct)

	return &models.Product{}
}

func (r *ProductUseCase) CSVProduct() error {
	return nil
}
