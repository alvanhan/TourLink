package migrations

import (
	"log"
	"tourlink/internal/user/domain"
	"tourlink/pkg/database"
)

func RunAutoMigrations() {
	err := database.DB.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}
	log.Println("database migrated successfully")
}
