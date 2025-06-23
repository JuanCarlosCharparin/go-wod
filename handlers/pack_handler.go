package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetPacks(c *gin.Context) {
	var packs []models.Pack
	database.DB.Preload("Gym").Find(&packs)
	var response []dto.PackResponse
	for _, cal := range packs {
		response = append(response, transformers.TransformPack(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetPackId(c *gin.Context) {

	id := c.Param("id")

	var pack models.Pack
	if err := database.DB.Preload("Gym").First(&pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack no encontrado"})
		return
	}
	response := transformers.TransformPack(pack)
	c.JSON(http.StatusOK, response)
}


func GetPackByGymId(c *gin.Context) {

	gymID := c.Param("id")

	var packs []models.Pack
	if err := database.DB.Preload("Gym").Where("gym_id = ?", gymID).Find(&packs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar packs del gimnasio"})
		return
	}

	if len(packs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron packs para este gimnasio"})
		return
	}

	var response []dto.PackResponse
	for _, pack := range packs {
		response = append(response, transformers.TransformPack(pack))
	}

	c.JSON(http.StatusOK, response)
}

func CreatePack(c *gin.Context) {
	var pack models.Pack
	if err := c.ShouldBindJSON(&pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", pack.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}
	
	database.DB.Create(&pack)
	database.DB.Preload("Gym").First(&pack, pack.Id)
	response := transformers.TransformPack(pack)
	c.JSON(http.StatusOK, response)
}


func UpdatedPack(c *gin.Context) {
	id := c.Param("id")

	var pack models.Pack
	if err := database.DB.Preload("Gym").First(&pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack no encontrado"})
		return
	}

	var updatedData models.Pack
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&pack).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio actualizado
	if err := database.DB.Preload("Gym").First(&pack, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio"})
		return
	}

	response := transformers.TransformPack(pack)
	c.JSON(http.StatusOK, response)
}


func DeletePack(c *gin.Context) {
	id := c.Param("id")

	var pack models.Pack
	if err := database.DB.First(&pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack no encontrado"})
		return
	}

	if err := database.DB.Delete(&pack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el pack"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pack eliminado correctamente"})
}