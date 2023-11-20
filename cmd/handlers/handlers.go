package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/umarrohman03/scrapping/bootstrap"
	"github.com/umarrohman03/scrapping/controller"
)

type MyHandler struct {
	Application bootstrap.Application
}

func NewServiceInitial(app bootstrap.Application) MyHandler {
	return MyHandler{
		Application: app,
	}
}

type ServerInterfaceWrapper struct {
	ProductHandler controller.IProductController
}

func (w *ServerInterfaceWrapper) AddProduct(ctx echo.Context) error {
	return w.ProductHandler.AddProduct(ctx)
}
