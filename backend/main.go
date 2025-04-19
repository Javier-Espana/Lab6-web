package main

import (
	"series-tracker-backend/database"
	"series-tracker-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "series-tracker-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	database.InitDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost"}, 
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},                
		AllowHeaders: []string{"*"},               
	}))

	apiRoutes := r.Group("/api")
	{
		apiRoutes.GET("/series", handlers.GetSeries)        
		apiRoutes.GET("/series/:id", handlers.GetSerieByID)   
		apiRoutes.POST("/series", handlers.CreateSerie)       
		apiRoutes.PUT("/series/:id", handlers.UpdateSerie)   
		apiRoutes.DELETE("/series/:id", handlers.DeleteSerie)

		apiRoutes.PATCH("/series/:id/status", handlers.UpdateStatus)      
		apiRoutes.PATCH("/series/:id/episode", handlers.IncrementEpisode) 
		apiRoutes.PATCH("/series/:id/upvote", handlers.Upvote)           
		apiRoutes.PATCH("/series/:id/downvote", handlers.Downvote)        
	}

	r.Run(":8080")
}
