package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	var countries []models.Country
	database.DB.Find(&countries)
	var response []dto.CountryResponse
	for _, cal := range countries {
		response = append(response, transformers.TransformCountry(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetCountryId(c *gin.Context) {

	id := c.Param("id")

	var country models.Country
	if err := database.DB.First(&country, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "País no encontrado"})
		return
	}

	database.DB.First(&country, id)
	response := transformers.TransformCountry(country)
	c.JSON(http.StatusOK, response)
}

func CreateCountry(c *gin.Context) {
	var country models.Country
	if err := c.ShouldBindJSON(&country); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&country)
	response := transformers.TransformCountry(country)
	c.JSON(http.StatusOK, response)
}


func UpdatedCountry(c *gin.Context) {
	id := c.Param("id")

	var country models.Country
	if err := database.DB.First(&country, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "País no encontrado"})
		return
	}

	var updatedData models.Country
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&country).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	response := transformers.TransformCountry(country)
	c.JSON(http.StatusOK, response)
}

func DeleteCountry(c *gin.Context) {
	id := c.Param("id")

	var country models.Country
	if err := database.DB.First(&country, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "País no encontrado"})
		return
	}

	if err := database.DB.Delete(&country).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el país"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "País eliminado correctamente"})
}