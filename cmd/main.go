package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/umarrohman03/scrapping/api"
	"github.com/umarrohman03/scrapping/bootstrap"
	"github.com/umarrohman03/scrapping/cmd/handlers"
)

func main() {
	fmt.Println("main page")
	e := echo.New()
	e.Static("/swagger", "cmd")

	// Start server
	app := bootstrap.NewInitializeBootstrap()

	// initial service dependencies
	serve := handlers.NewServiceInitial(app)
	productController := serve.ProductHandler()

	wrapper := &handlers.ServerInterfaceWrapper{
		ProductHandler: productController,
	}

	api.RegisterHandlers(e, wrapper)

	e.Logger.Fatal(e.Start(app.ENV.HTTPAddress))
}
