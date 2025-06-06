package database

import (
	"fmt"
	"log"
	"tourlink/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() error {
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	user := config.GetString("database.user")
	password := config.GetString("database.password")
	dbname := config.GetString("database.dbname")
	sslmode := config.GetString("database.sslmode")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	DB = db

	log.Println("database connected")
	return nil
}
