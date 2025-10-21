package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetWaitListByClassId(c *gin.Context) {
	class_id := c.Param("id")
	var waitlist []models.Waitlist
	database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").Where("class_id = ?", class_id).Find(&waitlist)

	if len(waitlist) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Lista de espera vacía"})
		return
	}

	response := make([]dto.WaitlistResponse, len(waitlist))
	for i, wl := range waitlist {
		response[i] = transformers.TransformWaitlist(wl)
	}

	c.JSON(http.StatusOK, response)
}