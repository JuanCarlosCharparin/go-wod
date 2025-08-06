package handlers

import (
	"math"
	"net/http"
	"strconv"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
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
	roleID := c.Param("rol")

	// Paginación
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	// Búsqueda y orden
	search := c.DefaultQuery("search", "")
	sort := c.DefaultQuery("sort", "id")
	order := c.DefaultQuery("order", "asc")

	db := database.DB.Model(&models.User{}).
		Preload("Gym").
		Preload("Role").
		Preload("UserPacks", "status = ?", 1).
		Where("gym_id = ? AND role_id = ?", gymID, roleID)

	if search != "" {
		searchTerm := "%" + search + "%"
		db = db.Where("name LIKE ? OR lastname LIKE ? OR email LIKE ? OR dni LIKE ?", searchTerm, searchTerm, searchTerm, searchTerm)
	}

	// Contar total antes de paginar
	var total int64
	db.Count(&total)

	// Obtener resultados paginados
	var users []models.User
	err := db.Order(sort + " " + order).Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	// Mapear a DTO
	var response []dto.UserResponseNoGym
	for _, user := range users {
		response = append(response, transformers.TransformUserNoGym(user))
	}

	// Devolver respuesta paginada
	c.JSON(http.StatusOK, dto.PaginatedUsersResponse{
		Data:       response,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	})
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
	// Establecer status en 1 (activo)
	user.Status = true
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

func DisableUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	user.Status = false
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo desactivar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario desactivado exitosamente"})
}

func EnableUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	user.Status = true
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo activar el usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario activado exitosamente"})
}
