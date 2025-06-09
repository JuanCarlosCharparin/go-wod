package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWods(c *gin.Context) {
	var wods []models.Wod
	database.DB.Preload("Gym.Country").Find(&wods)
	c.JSON(http.StatusOK, wods)
}

func GetWodId(c *gin.Context) {

	id := c.Param("id")

	var wod models.Wod
	if err := database.DB.Preload("Gym.Country").First(&wod, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wod no encontrado"})
		return
	}
	c.JSON(http.StatusOK, wod)
}

func CreateWod(c *gin.Context) {
	var wod models.Wod
	if err := c.ShouldBindJSON(&wod); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", wod.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}
	
	database.DB.Create(&wod)
	database.DB.Preload("Gym.Country").First(&wod, wod.ID)
	c.JSON(http.StatusOK, wod)
}


func UpdatedWod(c *gin.Context) {
	id := c.Param("id")

	var wod models.Wod
	if err := database.DB.Preload("Gym").First(&wod, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wod no encontrado"})
		return
	}

	var updatedData models.Wod
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&wod).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio y país actualizado
	if err := database.DB.Preload("Gym.Country").First(&wod, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio y el país"})
		return
	}

	c.JSON(http.StatusOK, wod)
}


func DeleteWod(c *gin.Context) {
	id := c.Param("id")

	var wod models.Wod
	if err := database.DB.First(&wod, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wod no encontrado"})
		return
	}

	if err := database.DB.Delete(&wod).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el wod"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Wod eliminado correctamente"})
}