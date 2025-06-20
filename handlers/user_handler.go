package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Gym").Preload("Role").Find(&users)
	var response []dto.UserResponse
	for _, cal := range users {
		response = append(response, transformers.TransformUser(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetUserId(c *gin.Context) {

	id := c.Param("id")

	var user models.User
	if err := database.DB.Preload("Gym").Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}
	response := transformers.TransformUser(user)
	c.JSON(http.StatusOK, response)
}

func GetUsersByGymId(c *gin.Context) {
	gymID := c.Param("id")

	var users []models.User
	if err := database.DB.Preload("Gym").Preload("Role").Where("gym_id = ?", gymID).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuarios"})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron usuarios para este gimnasio"})
		return
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, transformers.TransformUser(user))
	}

	c.JSON(http.StatusOK, response)
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
	database.DB.Preload("Gym").Preload("Role").First(&user, user.Id)
	response := transformers.TransformUser(user)
	c.JSON(http.StatusOK, response)
}


func UpdatedUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.Preload("Gym").Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&user).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio y país actualizado
	if err := database.DB.Preload("Gym.Country").Preload("Role").First(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio y el país"})
		return
	}

	response := transformers.TransformUser(user)
	c.JSON(http.StatusOK, response)
}


func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}