package handlers

import (
	"net/http"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/transformers"

	"github.com/gin-gonic/gin"
)


func GetTemplatesByGymId(c *gin.Context) {
	gymID := c.Param("id")

	var templates []models.ScheduleTemplate
	err := database.DB.
		Where("gym_id = ?", gymID).
		Find(&templates).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener plantillas"})
		return
	}

	var response []dto.ScheduleTemplateResponse
	for _, template := range templates {
		response = append(response, transformers.TransformScheduleTemplate(template))
	}

	c.JSON(http.StatusOK, response)
}



func GetScheduleTemplatesByGymID(c *gin.Context) {
	gymID := c.Param("id")

	var templates []models.ScheduleTemplate
	err := database.DB.
		Where("gym_id = ?", gymID).
		Preload("Gym").
		Preload("Blocks.Discipline").
		Find(&templates).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener plantillas"})
		return
	}

	if len(templates) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron plantillas para este gimnasio"})
		return
	}

	var response []dto.ScheduleTemplateResponse
	for _, template := range templates {
		response = append(response, transformers.TransformScheduleTemplate(template))
	}

	c.JSON(http.StatusOK, response)
}


func CreateScheduleTemplate(c *gin.Context) {
	var template models.ScheduleTemplate
	if err := c.ShouldBindJSON(&template); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que el gimnasio exista
	var gym models.Gym
	if err := database.DB.First(&gym, template.GymID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido"})
		return
	}

	if err := database.DB.Create(&template).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear plantilla"})
		return
	}

	database.DB.Preload("Gym").Preload("Blocks.Discipline").First(&template, template.ID)

	response := transformers.TransformScheduleTemplate(template)
	c.JSON(http.StatusOK, response)
}



func UpdateScheduleTemplate(c *gin.Context) {
	id := c.Param("id")

	var template models.ScheduleTemplate
	if err := database.DB.First(&template, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plantilla no encontrada"})
		return
	}

	var input models.ScheduleTemplate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar gimnasio si se está actualizando
	if input.GymID != 0 && input.GymID != template.GymID {
		var gym models.Gym
		if err := database.DB.First(&gym, input.GymID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido"})
			return
		}
	}

	template.Day = input.Day
	template.GymID = input.GymID

	if err := database.DB.Save(&template).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar plantilla"})
		return
	}

	database.DB.Preload("Gym").Preload("Blocks.Discipline").First(&template, template.ID)
	response := transformers.TransformScheduleTemplate(template)
	c.JSON(http.StatusOK, response)
}


func DeleteScheduleTemplate(c *gin.Context) {
	id := c.Param("id")

	var template models.ScheduleTemplate
	if err := database.DB.First(&template, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Template no encontrado"})
		return
	}

	// Eliminar bloques asociados
	if err := database.DB.Where("template_id = ?", id).Delete(&models.ScheduleBlock{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar bloques asociados"})
		return
	}

	// Eliminar el template
	if err := database.DB.Delete(&template).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el template"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Template y bloques eliminados correctamente"})
}


