package postgres

import (
	"datahandler_go/helpers"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	// Use postgres-db instead of localhost because this application is running within Docker
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		helpers.EnvVariable("POSTGRES_HOST"),
		helpers.EnvVariable("POSTGRES_DB_USER"),
		helpers.EnvVariable("POSTGRES_DB_PASSWORD"),
		helpers.EnvVariable("POSTGRES_DB_NAME"),
		helpers.EnvVariable("POSTGRES_PORT"),
		helpers.EnvVariable("POSTGRES_DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to Postgres database. \n", err)
	}

	DB = Dbinstance{
		Db: db,
	}
}

func DisconnectDb() {
	sqlDB, err := DB.Db.DB()
	if err != nil {
		log.Println("Error getting SQL database instance:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("Postgres database connection closed successfully.")
	}
}

func IsDbConnected() bool {
	sqlDB, err := DB.Db.DB()
	if err != nil {
		log.Println("Error getting SQL database instance:", err)
		return false
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Println("Postgres database connection is not healthy:", err)
		return false
	}

	return true
}
