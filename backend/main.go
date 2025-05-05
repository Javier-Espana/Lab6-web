// Punto de entrada principal de la aplicación
package main

import (
	// Importa los paquetes necesarios para la base de datos, controladores y middleware
	"series-tracker-backend/database"
	"series-tracker-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// Importa la documentación generada por Swagger
	_ "series-tracker-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Inicializa la conexión a la base de datos
	database.InitDB()

	// Crea una nueva instancia del servidor Gin
	r := gin.Default()

	// Configura la ruta para servir la documentación Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configura el middleware CORS para permitir solicitudes desde el frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"*"},
	}))

	// Configura las rutas para servir archivos estáticos del frontend
	r.StaticFile("/", "./frontend/index.html") // Sirve el archivo index.html para la ruta raíz
	r.Static("/static", "./frontend/static")   // Sirve archivos estáticos como imágenes, CSS, etc.

	// Define las rutas del API
	apiRoutes := r.Group("/api")
	{
		// Rutas para manejar las series
		apiRoutes.GET("/series", handlers.GetSeries)          // Obtiene la lista de series
		apiRoutes.GET("/series/:id", handlers.GetSerieByID)   // Obtiene una serie por su ID
		apiRoutes.POST("/series", handlers.CreateSerie)       // Crea una nueva serie
		apiRoutes.PUT("/series/:id", handlers.UpdateSerie)    // Actualiza una serie existente
		apiRoutes.DELETE("/series/:id", handlers.DeleteSerie) // Elimina una serie por su ID

		// Rutas para operaciones específicas de las series
		apiRoutes.PATCH("/series/:id/status", handlers.UpdateStatus)      // Actualiza el estado de una serie
		apiRoutes.PATCH("/series/:id/episode", handlers.IncrementEpisode) // Incrementa el número de episodios vistos
		apiRoutes.PATCH("/series/:id/upvote", handlers.Upvote)            // Incrementa el ranking de una serie
		apiRoutes.PATCH("/series/:id/downvote", handlers.Downvote)        // Decrementa el ranking de una serie
	}

	// Inicia el servidor en el puerto 8080
	r.Run(":8080")
}
