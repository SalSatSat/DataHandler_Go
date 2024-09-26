# DataHandler_Go

DataHandler_Go is an app implemented in [Go](https://go.dev/) to allow you to fetch, process and save data to a MongoDB database and serve it via API.

## Requirements
- [Go](https://go.dev/dl/) v1.23.0
- [Docker](https://www.docker.com/) v27.2.0, build 3ab4256
- [Docker Compose](https://docs.docker.com/compose/) v2.29.1-desktop.1

## Features
- Job scheduler to perform ETL process on raw data and store transformed data to a MongoDB database as a collection
- REST API to serve transformed data from the MongoDB collection where the name of the collection is the API endpoint
- Sample models and jobs where ETL process is performed on raw data from Postgres and MongoDB databases

## File Structure and Descriptions
### Root Directory
- **main.go**: This file acts as the main entry point for the web application built using the Fiber framework

### `database` Directory
- `mongo` Directory
  - **mongo.go**: This file is a helper script to handle connection management with MongoDB database
- `postgres` Directory
  - **postgres.go**: This file is a helper script to handle connection management with Postgres database

### `helpers` Directory
- **env.go**: This file loads environment variables from a .env file using the godotenv package and returns the value of a specified key

### `jobs` Directory
- **jobs.go**: This file schedules and manages recurring jobs using the `cron` package
- `samples` Director
  - **mongo_sample.go**: This file defines the job to perform ETL process on sample data and store them to a MongoDB database as a collection
  - **postgres_sample.go**: This file defines the job to perform ETL process on raw data from a Postgres database and store them to a MongoDB database as a collection

### `models` Directory
- **models.go**: This file schedules and manages recurring jobs using the `cron` package
- `samples` Director
  - **mongo_sample.go**: This file defines the schema of a record stored in a MongoDB collection
  - **postgres_sample.go**: This file defines the schema of a record stored in a Postgres table

### `routes` Directory
- **routes.go**: This file defines routes for a Fiber web application to query a MongoDB database and return data as CSV or JSON

## Getting Started
1. Fork this repository to your own GitHub account.
2. Clone the forked repository to your local machine.
    ```bash 
    git clone https://github.com/<YOUR_USERNAME>/DataHandler_Go.git
    ```
3. In the root folder, create a `.env` file and populate it with the necessary environment variables. Below is an example:
    ```
    PORT=42423

    # Postgres
    POSTGRES_HOST=postgres-db
    POSTGRES_PORT=5432
    POSTGRES_DB_USER=user
    POSTGRES_DB_PASSWORD=password123
    POSTGRES_DB_NAME=test-db
    POSTGRES_DB_TIMEZONE=Asia/Singapore

    # Mongo
    MONGO_HOST=mongo-db
    MONGO_PORT=27017
    MONGO_DB_USER=user
    MONGO_DB_PASSWORD=password123
    MONGO_DB_NAME=test-db
    ```

## Running the App
1. Open the terminal and change directory to the project folder
2. Build and run the app with Docker Compose
   ```bash
   docker compose up
   ```
3. You can test the app by calling the REST API (e.g. `localhost:42423/mongo_sample`)

## Additional Resources
- [github.com/gofiber/fiber/v2 v2.52.5](https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.52.5)
- [github.com/joho/godotenv v1.5.1](https://pkg.go.dev/github.com/joho/godotenv@v1.5.1)
- [github.com/robfig/cron/v3 v3.0.0](https://pkg.go.dev/github.com/robfig/cron/v3@v3.0.0)
- [go.mongodb.org/mongo-driver v1.17.0](https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.17.0)
- [gorm.io/driver/postgres v1.5.9](https://pkg.go.dev/gorm.io/driver/postgres@v1.5.9)
- [gorm.io/gorm v1.25.12](https://pkg.go.dev/gorm.io/gorm@v1.25.12)

## Team
- Muhammad Salihin Bin Zaol-kefli: salsatsat@gmail.com

## License
This project is distributed under the [MIT license](https://en.wikipedia.org/wiki/MIT_License) (see the [LICENSE](./LICENSE.md) file)

## Support
For any issues or support requests, please create an issue on the repository