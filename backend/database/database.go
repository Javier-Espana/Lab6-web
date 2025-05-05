package database

import (
	"series-tracker-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es la conexión global a la base de datos
var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	// Usa los detalles de conexión a la base de datos definidos en docker-compose.yml
	dsn := "host=db user=user password=password dbname=seriesdb port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar con la base de datos: " + err.Error())
	}

	// Realiza la migración automática del esquema de la base de datos
	DB.AutoMigrate(&models.Serie{})
}
