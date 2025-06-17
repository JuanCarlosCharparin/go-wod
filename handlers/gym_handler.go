package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetGyms(c *gin.Context) {
	var gyms []models.Gym
	database.DB.Preload("Country").Find(&gyms)
	var response []dto.GymResponse
	for _, cal := range gyms {
		response = append(response, transformers.TransformGym(cal))
	}
	c.JSON(http.StatusOK, response)
}


func GetGymId(c *gin.Context) {

	id := c.Param("id")

	var gym models.Gym
	if err := database.DB.Preload("Country").First(&gym, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gimnasio no encontrado"})
		return
	}
	response := transformers.TransformGym(gym)
	c.JSON(http.StatusOK, response)
}


func GetGymsByCountryId(c *gin.Context) {
	countryID := c.Param("id")

	var gyms []models.Gym
	if err := database.DB.Preload("Country").Where("country_id = ?", countryID).Find(&gyms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar gimnasios"})
		return
	}

	if len(gyms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron gimnasios para este país"})
		return
	}

	var response []dto.GymResponse
	for _, gym := range gyms {
		response = append(response, transformers.TransformGym(gym))
	}

	c.JSON(http.StatusOK, response)
}


func CreateGym(c *gin.Context) {
	var gym models.Gym
	if err := c.ShouldBindJSON(&gym); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el país exista y no esté eliminado
	var country models.Country
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", gym.CountryId).
		First(&country).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "País no válido o eliminado"})
		return
	}
	database.DB.Create(&gym)
	database.DB.Preload("Country").First(&gym, gym.Id)
	response := transformers.TransformGym(gym)
	c.JSON(http.StatusOK, response)
}


func UpdatedGym(c *gin.Context) {
	id := c.Param("id")

	var gym models.Gym
	if err := database.DB.Preload("Country").First(&gym, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gimnasio no encontrado"})
		return
	}

	var updatedData models.Gym
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&gym).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el país actualizado
	if err := database.DB.Preload("Country").First(&gym, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar país"})
		return
	}

	response := transformers.TransformGym(gym)
	c.JSON(http.StatusOK, response)
}

func DeleteGym(c *gin.Context) {
	id := c.Param("id")

	var gym models.Gym
	if err := database.DB.First(&gym, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gimnasio no encontrado"})
		return
	}

	if err := database.DB.Delete(&gym).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el gimansio"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gimnasio eliminado correctamente"})
}