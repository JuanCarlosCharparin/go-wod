package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetUserPacks(c *gin.Context) {
	var user_packs []models.UserPack
	database.DB.Preload("Gym").Preload("User").Preload("Pack").Preload("Discipline").Find(&user_packs)
	var response []dto.UserPackResponse
	for _, cal := range user_packs {
		response = append(response, transformers.TransformUserPack(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetUserPackId(c *gin.Context) {

	id := c.Param("id")

	var user_pack models.UserPack
	if err := database.DB.Preload("Gym").Preload("User").Preload("Pack").Preload("Discipline").First(&user_pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pack no encontrado"})
		return
	}
	response := transformers.TransformUserPack(user_pack)
	c.JSON(http.StatusOK, response)
}


//Devuelve todos los packs de un usuario
func GetUserPackByUserId(c *gin.Context){

	user_id := c.Param("id")

	var user_packs []models.UserPack
	if err := database.DB.Preload("User").Where("user_id = ?", user_id).Find(&user_packs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuarios de este pack"})
		return
	}

	if len(user_packs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron packs para este usuario"})
		return
	}

	var response []dto.UserPackResponse
	for _, user_pack := range user_packs {
		response = append(response, transformers.TransformUserPack(user_pack))
	}

	c.JSON(http.StatusOK, response)
}




func CreateUserPack(c *gin.Context) {
	var user_pack models.UserPack
	if err := c.ShouldBindJSON(&user_pack); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", user_pack.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}
	//Validar que el usuario exista y no esté eliminado
	var user models.User
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", user_pack.UserId).
		First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no válido o eliminado"})
		return
	}
	//Validar que el pack exista y no esté eliminado
	var pack models.Pack
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", user_pack.PackId).
		First(&pack).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pack no válido o eliminado"})
		return
	}
	//Validar que la disciplina exista y no esté eliminada
	var discipline models.Discipline
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", user_pack.DisciplineId).
		First(&discipline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disciplina no válida o eliminada"})
		return
	}
	
	database.DB.Create(&user_pack)
	database.DB.Preload("Gym").Preload("User").Preload("Pack").Preload("Discipline").First(&user_pack, user_pack.Id)
	response := transformers.TransformUserPack(user_pack)
	c.JSON(http.StatusOK, response)
}


func UpdatedUserPack(c *gin.Context) {
	id := c.Param("id")

	var user_pack models.UserPack
	if err := database.DB.Preload("Gym").Preload("User").Preload("Pack").Preload("Discipline").First(&user_pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UserPack no encontrado"})
		return
	}

	var updatedData models.UserPack
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&user_pack).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio actualizado
	if err := database.DB.Preload("Gym").Preload("User").Preload("Pack").Preload("Discipline").First(&user_pack, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio"})
		return
	}

	response := transformers.TransformUserPack(user_pack)
	c.JSON(http.StatusOK, response)
}


func DeleteUserPack(c *gin.Context) {
	id := c.Param("id")

	var user_pack models.UserPack
	if err := database.DB.First(&user_pack, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UserPack no encontrado"})
		return
	}

	if err := database.DB.Delete(&user_pack).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el userpack"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "UserPack eliminado correctamente"})
}