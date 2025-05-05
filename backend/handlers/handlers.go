package handlers

import (
	"series-tracker-backend/database"
	"series-tracker-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Función para actualizar una serie existente en la base de datos
// Recibe el ID de la serie como parámetro y los datos actualizados en el cuerpo de la solicitud
func UpdateSerie(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie

	if err := database.DB.First(&serie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&serie); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&serie).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar la serie"})
		return
	}

	c.JSON(200, serie)
}

// Función para obtener una lista de series con filtros opcionales (estado, búsqueda y orden)
func GetSeries(c *gin.Context) {
	var series []models.Serie
	db := database.DB

	status := c.Query("status")
	search := c.Query("search")

	if status != "" {
		db = db.Where("status = ?", status)
	}

	if search != "" {
		db = db.Where("title ILIKE ?", "%"+search+"%")
	}

	sortOrder := c.Query("sort")
	if sortOrder == "asc" {
		db = db.Order("ranking asc")
	} else {
		db = db.Order("ranking desc")
	}

	if err := db.Find(&series).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener series"})
		return
	}

	c.JSON(200, series)
}

// Función para crear una nueva serie en la base de datos
// Valida que el campo 'title' sea obligatorio
func CreateSerie(c *gin.Context) {
	var newSerie models.Serie

	if err := c.ShouldBindJSON(&newSerie); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if newSerie.Title == "" {
		c.JSON(400, gin.H{"error": "El campo 'title' es obligatorio"})
		return
	}

	result := database.DB.Create(&newSerie)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al crear la serie: " + result.Error.Error()})
		return
	}

	c.JSON(201, newSerie)
}

// Función para incrementar el número de episodios vistos de una serie
// Valida que no se exceda el total de episodios disponibles
func IncrementEpisode(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie

	if err := database.DB.First(&serie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	if serie.LastEpisodeWatched >= serie.TotalEpisodes {
		c.JSON(400, gin.H{"error": "Ya has visto todos los episodios"})
		return
	}

	serie.LastEpisodeWatched++

	if err := database.DB.Save(&serie).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el episodio: " + err.Error()})
		return
	}

	c.JSON(200, serie)
}

// Función para obtener una serie específica por su ID
func GetSerieByID(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie
	result := database.DB.First(&serie, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, serie)
}

// Función para eliminar una serie de la base de datos por su ID
func DeleteSerie(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Serie{}, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar la serie"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Serie eliminada correctamente"})
}

// Función para actualizar el estado de una serie
// Recibe el nuevo estado en el cuerpo de la solicitud
func UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("status", request.Status)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el estado"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Estado actualizado correctamente"})
}

// Función para incrementar el ranking de una serie (upvote)
func Upvote(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("ranking", gorm.Expr("ranking + 1"))

	handleVoteResult(c, result)
}

// Función para decrementar el ranking de una serie (downvote)
func Downvote(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("ranking", gorm.Expr("ranking - 1"))

	handleVoteResult(c, result)
}

// Función auxiliar para manejar el resultado de las operaciones de votación (upvote/downvote)
func handleVoteResult(c *gin.Context, result *gorm.DB) {
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el ranking"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Ranking actualizado correctamente"})
}
