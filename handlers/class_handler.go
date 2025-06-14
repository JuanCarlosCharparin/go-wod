package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []models.Class
	database.DB.Preload("Gym.Country").Preload("Discipline").Find(&classes)
	c.JSON(http.StatusOK, classes)
}

func GetClassId(c *gin.Context) {

	id := c.Param("id")

	var class models.Class
	if err := database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	c.JSON(http.StatusOK, class)
}

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", class.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}

	//Validar que la disciplina exista y no esté eliminada
	var discipline models.Discipline
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", class.DisciplineId).
		First(&discipline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disciplina no válida o eliminada"})
		return
	}
	
	database.DB.Create(&class)
	database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, class.ID)
	c.JSON(http.StatusOK, class)
}


func UpdatedClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}

	var updatedData models.Class
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&class).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio, país y disciplina actualizada
	if err := database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio, país y disciplina"})
		return
	}

	c.JSON(http.StatusOK, class)
}


func DeleteClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}

	if err := database.DB.Delete(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la clase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clase eliminada correctamente"})
}