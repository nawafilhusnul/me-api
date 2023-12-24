package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormClient() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/invoice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to mySQL using gorm: %v", err)
	}

	return db
}
