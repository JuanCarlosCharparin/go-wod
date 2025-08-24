package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetDisciplines(c *gin.Context) {
	var disciplines []models.Discipline
	database.DB.Preload("Gym").Find(&disciplines)
	var response []dto.DisciplineResponse
	for _, cal := range disciplines {
		response = append(response, transformers.TransformDiscipline(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetDisciplineId(c *gin.Context) {

	id := c.Param("id")

	var discipline models.Discipline
	if err := database.DB.Preload("Gym").First(&discipline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Disciplina no encontrada"})
		return
	}

	database.DB.First(&discipline, id)
	response := transformers.TransformDiscipline(discipline)
	c.JSON(http.StatusOK, response)
}

//disciplinas por gym
func GetDisciplinesByGymId(c *gin.Context) {
	gymId := c.Param("id")

	var disciplines []models.Discipline
	if err := database.DB.Preload("Gym").Where("gym_id = ?", gymId).Find(&disciplines).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron disciplinas"})
		return
	}

	var response []dto.DisciplineResponse
	for _, discipline := range disciplines {
		response = append(response, transformers.TransformDiscipline(discipline))
	}
	c.JSON(http.StatusOK, response)
}

func CreateDiscipline(c *gin.Context) {
	var discipline models.Discipline
	if err := c.ShouldBindJSON(&discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Preload("Gym").Create(&discipline)
	response := transformers.TransformDiscipline(discipline)
	c.JSON(http.StatusOK, response)
}


func UpdatedDiscipline(c *gin.Context) {
	id := c.Param("id")

	var discipline models.Discipline
	if err := database.DB.First(&discipline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Disciplina no encontrada"})
		return
	}

	var updatedData models.Discipline
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&discipline).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}
	response := transformers.TransformDiscipline(discipline)
	c.JSON(http.StatusOK, response)
}

func DeleteDiscipline(c *gin.Context) {
	id := c.Param("id")

	var discipline models.Discipline
	if err := database.DB.First(&discipline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Disciplina no encontrado"})
		return
	}

	if err := database.DB.Delete(&discipline).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la disciplina"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Disciplina eliminado correctamente"})
}