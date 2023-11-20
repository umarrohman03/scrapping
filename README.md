
## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.21.3
2. [Docker](https://docs.docker.com/get-docker/) version 20
3. [Docker Compose](https://docs.docker.com/compose/install/) version 1.29
4. [GNU Make](https://www.gnu.org/software/make/)
5. [oapi-codegen](https://github.com/deepmap/oapi-codegen)

   Install the latest version with:
    ```
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    ```
6. [mock](https://github.com/golang/mock)

   Install the latest version with:
    ```
    go install github.com/golang/mock/mockgen@latest
    ```

## Initiate The Project

To start working, execute

```
make init
```

## Running

To run the project, run the following command:

```
docker-compose up --build
```

You should be able to access the API at http://localhost:8080

If you change `database.sql` file, you need to reinitate the database by running:

```
docker-compose down --volumes
```

## Testing

To run test, run the following command:

```
make test
```


## Detail Update Code Test Interview Backend
Aplicant

| Name        | Email                  |
|-------------|------------------------|
| Umar rohman | umarrohman03@gmail.com |



### Explain Structure file
| Folder     | Description                                    | Example |
|------------|------------------------------------------------|---------|
| Api        | Generator swagger api contract                 |         |
| bootstrap  | initial all depend injection for service       |         |
| cmd        | main application for project                   |         |
| commons    | code common function used                      |         |
| controller | verify all request server http                 |         |
| db         | for engine migration database                  |         |
| internal   | all related dependencies liblary using service |         |
| model      | define field for table / document              |         |
| pkg        | define http to extnal service                  |         |
| repository | get data from client or database redis etc     |         |
| usecase    | for put logic api                              |         |

### Explain New Schema Table
| Field       | Data Type |       value |
|-------------|:---------:|------------:|
| id          |  serial   | PRIMARY KEY |
| name        |  varchar  |         255 |
| description |  varchar  |         255 |
| image_link  |  varchar  |         255 |
| price       |   float   |             |
| rating      |  integer  |             |














