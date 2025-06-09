package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDisciplines(c *gin.Context) {
	var disciplines []models.Discipline
	database.DB.Find(&disciplines)
	c.JSON(http.StatusOK, disciplines)
}

func GetDisciplineId(c *gin.Context) {

	id := c.Param("id")

	var discipline models.Discipline
	if err := database.DB.First(&discipline, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Disciplina no encontrada"})
		return
	}

	database.DB.First(&discipline, id)
	c.JSON(http.StatusOK, discipline)
}

func CreateDiscipline(c *gin.Context) {
	var discipline models.Discipline
	if err := c.ShouldBindJSON(&discipline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&discipline)
	c.JSON(http.StatusOK, discipline)
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

	c.JSON(http.StatusOK, discipline)
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