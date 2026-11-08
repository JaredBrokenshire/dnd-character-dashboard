package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func Init() (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		os.Getenv("DB_PASSWORD"),
		host,
		port,
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to %v on port %v with username %v", host, port, user)

	db, err := gorm.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	sqlx.NewDb(db.DB(), driver)

	return db, nil
}
