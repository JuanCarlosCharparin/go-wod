package handlers

import (
	"net/http"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/transformers"

	"github.com/gin-gonic/gin"
)

func GetUserPackDisciplines(c *gin.Context) {
	var user_packs_disciplines []models.UserPackDiscipline

	err := database.DB.
		Preload("UserPack.User").
		Preload("UserPack.Pack").
		Preload("UserPack.Gym").
		Preload("Discipline").
		Find(&user_packs_disciplines).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener datos"})
		return
	}

	var response []dto.UserPackDisciplineResponse
	for _, upd := range user_packs_disciplines {
		response = append(response, transformers.TransformUserPackDiscipline(upd))
	}

	c.JSON(http.StatusOK, response)
}


//devuelve los packs activos de los usuarios
func GetActiveUserPackDisciplines(c *gin.Context) {
	var user_packs_disciplines []models.UserPackDiscipline

	err := database.DB.
		Joins("JOIN users_packs ON users_packs.id = user_pack_disciplines.user_pack_id").
		Where("users_packs.status = ?", 1).
		Preload("UserPack.User").
		Preload("UserPack.Pack").
		Preload("UserPack.Gym").
		Preload("Discipline").
		Find(&user_packs_disciplines).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener packs activos"})
		return
	}

	var response []dto.UserPackDisciplineResponse
	for _, upd := range user_packs_disciplines {
		response = append(response, transformers.TransformUserPackDiscipline(upd))
	}

	c.JSON(http.StatusOK, response)
}


func CreateUserPackDisciplines(c *gin.Context) {
	var input struct {
		UserPackId    uint   `json:"user_pack_id"`
		DisciplineIds []uint `json:"discipline_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	for _, disciplineId := range input.DisciplineIds {
		upd := models.UserPackDiscipline{
			UserPackId:   input.UserPackId,
			DisciplineId: disciplineId,
		}

		// Evitar duplicados
		var existing models.UserPackDiscipline
		err := database.DB.
			Where("user_pack_id = ? AND discipline_id = ?", input.UserPackId, disciplineId).
			First(&existing).Error
		if err == nil {
			continue // ya existe, no crear
		}

		if err := database.DB.Create(&upd).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al asociar disciplina"})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Disciplinas asociadas correctamente al pack del usuario",
	})
}
