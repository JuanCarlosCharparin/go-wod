package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"wod-go/dto"
	"wod-go/transformers"
	"github.com/gin-gonic/gin"
)

func GetCalendars(c *gin.Context) {
	var calendars []models.Calendar
	database.DB.
		Preload("User").
		Preload("Class.Gym").
		Preload("Class.Discipline").
		Find(&calendars)

	var response []dto.CalendarResponse
	for _, cal := range calendars {
		response = append(response, transformers.TransformCalendar(cal))
	}

	c.JSON(http.StatusOK, response)
}

func GetCalendarId(c *gin.Context) {

	id := c.Param("id")

	var calendar models.Calendar
	if err := database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").First(&calendar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Calendario no encontrado"})
		return
	}
	response := transformers.TransformCalendar(calendar)
	c.JSON(http.StatusOK, response)
}


func GetUsersByClassId(c *gin.Context) {
	classID := c.Param("id")

	var calendars []models.Calendar
	if err := database.DB.Preload("User").Where("class_id = ?", classID).Find(&calendars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuarios"})
		return
	}

	if len(calendars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron usuarios para esta clase"})
		return
	}

	var response []dto.UserResponse
	for _, calendar := range calendars {
		response = append(response, transformers.TransformUser(calendar.User))
	}

	c.JSON(http.StatusOK, response)
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
	database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").First(&calendar, calendar.Id)
	response := transformers.TransformCalendar(calendar)
	c.JSON(http.StatusOK, response)
}


func UpdatedCalendar(c *gin.Context) {
	id := c.Param("id")

	var calendar models.Calendar
	if err := database.DB.Preload("Gym").Preload("Discipline").First(&calendar, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gimnasio no encontrada"})
		return
	}

	var updatedData models.Calendar
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&calendar).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar"})
		return
	}

	// Recargamos con el gimnasio, país y disciplina actualizada
	if err := database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").First(&calendar, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio y disciplina"})
		return
	}

	response := transformers.TransformCalendar(calendar)
	c.JSON(http.StatusOK, response)
}


func DeleteCalendar(c *gin.Context) {
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
}