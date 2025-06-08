package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Gym").Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUserId(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	if err := database.DB.Preload("Gym").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User no encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", user.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}
	
	//Validar que el usuario no se encuentre registrado
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El email ya está registrado"})
		return
	}
	database.DB.Create(&user)
	database.DB.Preload("Gym").First(&user, user.ID)
	c.JSON(http.StatusOK, user)
}