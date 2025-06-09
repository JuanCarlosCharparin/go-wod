package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCalendars(c *gin.Context) {
	var calendars []models.Calendar
	database.DB.Preload("User").Preload("Class").Find(&calendars)
	c.JSON(http.StatusOK, calendars)
}

func GetCalendarId(c *gin.Context) {

	id := c.Param("id")

	var calendar models.Calendar
	if err := database.DB.Preload("User").Preload("Class").First(&calendar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calendario no encontrado"})
		return
	}
	c.JSON(http.StatusOK, calendar)
}

func CreateCalendar(c *gin.Context) {
	var calendar models.Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el usuario exista y no esté eliminado
	var user models.User
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", calendar.UserId).
		First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no válido o eliminado"})
		return
	}

	//Validar que la clase exista y no esté eliminada
	var class models.Class
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", calendar.ClassId).
		First(&class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Clase no válida o eliminada"})
		return
	}
	
	database.DB.Create(&calendar)
	database.DB.Preload("User").Preload("Class").First(&calendar, calendar.ID)
	c.JSON(http.StatusOK, class)
}


/*func UpdatedClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}

	var updatedData models.Class
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&class).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio, país y disciplina actualizada
	if err := database.DB.Preload("Gym.Country").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio, país y disciplina"})
		return
	}

	c.JSON(http.StatusOK, class)
}


func DeleteClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}

	if err := database.DB.Delete(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la clase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clase eliminada correctamente"})
}*/