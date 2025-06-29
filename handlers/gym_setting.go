package handlers

import (
	"wod-go/database"
	"wod-go/models"
	//"wod-go/services"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
	//"time"
)

func GetSettings(c *gin.Context) {
	var settings []models.GymSetting
	database.DB.Preload("Gym").Find(&settings)
	var response []dto.GymSettingResponse
	for _, cal := range settings {
		response = append(response, transformers.TransformGymSetting(cal))
	}
	c.JSON(http.StatusOK, response)
}