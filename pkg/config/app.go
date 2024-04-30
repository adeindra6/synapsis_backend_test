package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Can't load .env file")
	}

	db_url := os.Getenv("DATABASE_URL")
	db_port := os.Getenv("DATABASE_PORT")
	db_name := os.Getenv("DATABASE_NAME")
	db_username := os.Getenv("DATABASE_USERNAME")
	db_password := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		db_username,
		db_password,
		db_url,
		db_port,
		db_name)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
