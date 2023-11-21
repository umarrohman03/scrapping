package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/umarrohman03/scrapping/internal/env"
)

type PostgresClient struct {
	DB *sql.DB
}

func NewPostgresDatabase(env *env.ENV) PostgresClient {
	fmt.Println("start postgres")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBPostgresHost
	dbPort := env.DBPostgresPort
	dbUser := env.DBPostgresUser
	dbPass := env.DBPostgresPass
	dbName := env.DBPostgresName

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	client, err := NewClientPostgres(connection)
	if err != nil {
		log.Fatal(err)
	}

	err = client.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func NewClientPostgres(connection string) (PostgresClient, error) {

	dsn := fmt.Sprintf("%s", connection)

	dbConn, err := sql.Open(`postgres`, dsn)
	//if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error conn")
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return PostgresClient{DB: dbConn}, err
}
