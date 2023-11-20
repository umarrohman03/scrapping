package repositories

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"

	"github.com/umarrohman03/scrapping/commons"
	"github.com/umarrohman03/scrapping/internal/db"
	models "github.com/umarrohman03/scrapping/model"
)

type ProductsRepository struct {
	DB            db.PostgresClient
	ProductsModel models.IProduct
	Scrapper      *colly.Collector
}

type IProductsRepository interface {
	GetProduct(ctx context.Context) (string, error)
	InsertProduct(ctx context.Context, data models.Product) (int64, error)
}

func NewScrapsRepository(db db.PostgresClient, tableName models.IProduct, scrapper *colly.Collector) IProductsRepository {
	return &ProductsRepository{
		DB:            db,
		ProductsModel: tableName,
		Scrapper:      scrapper,
	}
}

func (r *ProductsRepository) InsertProduct(ctx context.Context, data models.Product) (int64, error) {
	var lastInsertID int64

	tx, err := r.DB.DB.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := r.DB.DB.PrepareContext(ctx, "INSERT INTO product (product_name, description, image_link, store_name, price, rating) VALUES ($1, $2, $3, $4, $5, $6)  RETURNING product_id")
	if err != nil {
		return 0, commons.ErrorInternalServer
	}

	err = stmt.QueryRowContext(ctx, data.Name, data.Description, data.ImageLink, data.StoreName, data.Price, data.Rating).Scan(&lastInsertID)
	if err != nil {
		// Rollback the transaction if there's an error
		tx.Rollback()

		return 0, commons.ErrorProductAlreadyExist
	}
	err = tx.Commit()
	if err != nil {
		return 0, commons.ErrorInternalServer
	}
	return lastInsertID, nil

}

// GetProduct ...
func (r *ProductsRepository) GetProduct(ctx context.Context) (string, error) {

	//todo get image still failed because the image exchange
	//todo move css to env
	//todo parse data
	r.Scrapper.OnHTML("a.css-54k5sq div.css-16vw0vn", func(e *colly.HTMLElement) {

		image := e.ChildAttr("div.css-79elbk div.css-377m5r div.css-1g5og91 img", "src")
		fmt.Println(image)

		product := e.ChildText("div.css-11s9vse")
		fmt.Println(product)

		price := e.ChildText("div.css-11s9vse span.css-o5uqvq")
		fmt.Println(price)

		rating := e.ChildAttrs("div.css-11s9vse img.css-177n1u3", "src")
		fmt.Println(len(rating))

	})

	return "", nil
}
