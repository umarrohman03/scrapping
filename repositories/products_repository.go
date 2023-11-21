package repositories

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"time"

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
	fmt.Println("start get product repo")

	// =============== multi thread ================
	r.Scrapper.Limit(&colly.LimitRule{
		Parallelism: 5,
		Delay:       5 * time.Second,
	})

	//todo mode csv process to usecase
	// ============== CREATE CSV ====================
	fName := "tokopedia-list-product.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write CSV header
	writer.Write([]string{"Image", "Name", "Price", "Rating"})

	//todo get image still failed because the image exchange
	//todo move css to env
	//todo parse data
	var row int
	r.Scrapper.OnHTML("a.css-54k5sq div.css-16vw0vn", func(e *colly.HTMLElement) {
		row++
		rating := e.ChildAttrs("div.css-11s9vse img.css-177n1u3", "src")
		image := e.ChildAttr("div.css-79elbk div.css-377m5r div.css-1g5og91 img", "src")
		if image == "" {
			image = "empty image url"
		}
		price := e.ChildText("div.css-11s9vse span.css-o5uqvq")
		if row < 100 {
			writer.Write([]string{
				image,
				e.ChildText("div.css-11s9vse"),
				price,
				strconv.Itoa(len(rating)),
			})
		}

	})

	// Start scraping in five threads on www.tokopedia.com
	for i := 0; i < 5; i++ {
		if i == 0 {
			r.Scrapper.Visit(fmt.Sprintf("%s?page=%d", "https://www.tokopedia.com/p/handphone-tablet/handphone", i))
		}
		if i == 1 {
			r.Scrapper.Visit(fmt.Sprintf("%s?page=%d", "https://www.tokopedia.com/p/handphone-tablet/handphone", i))
		}
		if i == 2 {
			r.Scrapper.Visit(fmt.Sprintf("%s?page=%d", "https://www.tokopedia.com/p/handphone-tablet/handphone", i))
		}
		if i == 3 {
			r.Scrapper.Visit(fmt.Sprintf("%s?page=%d", "https://www.tokopedia.com/p/handphone-tablet/handphone", i))
		}
		if i == 4 {
			r.Scrapper.Visit(fmt.Sprintf("%s?page=%d", "https://www.tokopedia.com/p/handphone-tablet/handphone", i))
		}
	}
	// Wait until threads are finished
	r.Scrapper.Wait()

	return "", nil
}
