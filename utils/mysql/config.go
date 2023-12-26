package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormClient() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger(),
	})
	if err != nil {
		log.Fatalf("failed to connect to mySQL using gorm: %v", err)
	}

	autoMigrate(db)

	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.
		Set("gorm:table_options", "ENGINE=InnoDB").
		Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(
			domain.Project{},
		)
	if err != nil {
		log.Fatalf("auto migrate failed, err: %v", err)
	}
}

func gormLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			Colorful:                  false,                  // Disable color
		},
	)

	return newLogger
}
