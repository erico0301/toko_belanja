package database

import (
	"fmt"
	"log"
	"toko_belanja/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "1234"
	dbPort   = "5432"
	dbName   = "toko-belanja"

	// host     = os.Getenv("DB_HOST")
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// dbPort   = os.Getenv("DB_PORT")
	// dbName   = os.Getenv("DB_NAME")

	db  *gorm.DB
	err error
)

func StartDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error Connection to database : ", err)
	}

	fmt.Println("Succes Connection to Database")
	db.Debug().AutoMigrate(model.User{}, model.Product{}, model.Category{}, model.Transaction{})
}

func GetDB() *gorm.DB {
	return db
}
