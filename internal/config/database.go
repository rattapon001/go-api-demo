package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"demo1/internal/entity"
)

func InitializeDatabase() *gorm.DB {

	dbHost := "host=" + os.Getenv("DB_HOST")
	dbUsername := "user=" + os.Getenv("DB_USERNAME")
	dbPassword := "password=" + os.Getenv("DB_PASSWORD")
	dbName := "dbname=" + os.Getenv("DB_DATABASE")
	dbPort := "port=" + os.Getenv("DB_PORT")
	dsn := dbHost + " " + dbUsername + " " + dbPassword + " " + dbName + " " + dbPort
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entity.User{}, entity.MediaObject{})
	return db
}
