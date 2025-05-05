package database

import (
	"series-tracker-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Use hardcoded database connection details from docker-compose.yml
	dsn := "host=db user=user password=password dbname=seriesdb port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Auto-migrate the database schema
	DB.AutoMigrate(&models.Serie{})
}
