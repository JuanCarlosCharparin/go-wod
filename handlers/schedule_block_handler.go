package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateScheduleBlock(c *gin.Context) {
	var blocks []models.ScheduleBlock

	if err := c.ShouldBindJSON(&blocks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, block := range blocks {
		// Validar que exista la plantilla
		if err := database.DB.First(&models.ScheduleTemplate{}, block.TemplateID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Plantilla no encontrada"})
			return
		}
		// Validar que exista la disciplina
		if err := database.DB.First(&models.Discipline{}, block.DisciplineID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Disciplina no válida"})
			return
		}

		// Crear el bloque
		if err := database.DB.Create(&block).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear bloque de horario"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bloques creados correctamente"})
}
