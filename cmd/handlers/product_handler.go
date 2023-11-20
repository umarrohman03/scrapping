package handlers

import (
	"github.com/umarrohman03/scrapping/controller"
	"github.com/umarrohman03/scrapping/model"
	"github.com/umarrohman03/scrapping/repositories"
	"github.com/umarrohman03/scrapping/usecase"
)

func (h *MyHandler) ProductHandler() controller.IProductController {
	table := model.NewProduct()
	productRepository := repositories.NewScrapsRepository(h.Application.Postgres, table, h.Application.Scrapper)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)
	return productController
}
