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


//agregar solo un bloque
func AddScheduleTemplateBlock(c *gin.Context) {
	var input models.ScheduleBlock
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Bloque creado correctamente"})
}




func UpdateScheduleBlock(c *gin.Context) {
	id := c.Param("id")

	var block models.ScheduleBlock
	if err := database.DB.First(&block, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bloque no encontrado"})
		return
	}

	var input models.ScheduleBlock
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validaciones si cambian los campos clave
	if input.TemplateID != 0 && input.TemplateID != block.TemplateID {
		if err := database.DB.First(&models.ScheduleTemplate{}, input.TemplateID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Plantilla no válida"})
			return
		}
	}

	if input.DisciplineID != 0 && input.DisciplineID != block.DisciplineID {
		if err := database.DB.First(&models.Discipline{}, input.DisciplineID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Disciplina no válida"})
			return
		}
	}

	// Actualizar campos
	block.StartTime = input.StartTime
	block.EndTime = input.EndTime
	block.Capacity = input.Capacity
	block.TemplateID = input.TemplateID
	block.DisciplineID = input.DisciplineID

	if err := database.DB.Save(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar bloque"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bloque actualizado correctamente"})
}


func DeleteScheduleBlock(c *gin.Context) {
	// Obtener el ID del bloque desde la URL
	id := c.Param("id")

	// Buscar el bloque
	var block models.ScheduleBlock
	if err := database.DB.First(&block, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bloque no encontrado"})
		return
	}

	// Eliminar el bloque
	if err := database.DB.Delete(&block).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el bloque"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bloque eliminado correctamente"})
}

