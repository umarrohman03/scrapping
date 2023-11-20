package bootstrap

import (
	"github.com/gocolly/colly"
	"github.com/umarrohman03/scrapping/internal/db"
	"github.com/umarrohman03/scrapping/internal/env"
	"github.com/umarrohman03/scrapping/internal/scrapper"
)

type Application struct {
	Postgres db.PostgresClient
	ENV      *env.ENV
	Scrapper *colly.Collector
}

func NewInitializeBootstrap() Application {
	app := Application{}
	app.ENV = env.NewENV()
	app.Postgres = db.NewPostgresDatabase(app.ENV)
	app.Scrapper = scrapper.NewScrapperClient(app.ENV)
	return app
}
