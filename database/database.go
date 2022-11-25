package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"demo1/database/model"
)

func Initialize() *gorm.DB {

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

	db.AutoMigrate(&model.User{}, model.MediaObject{})
	return db
}
