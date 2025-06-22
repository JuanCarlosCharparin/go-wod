package handlers

import (
	"net/http"
	"time"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/services"
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
		response = append(response, transformers.TransformCalendar(cal, nil))
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
	response := transformers.TransformCalendar(calendar, nil)
	c.JSON(http.StatusOK, response)
}


func GetUsersByClassId(c *gin.Context) {
	classID := c.Param("id")

	var calendars []models.Calendar
	if err := database.DB.Preload("User.Gym").Preload("User.Role").Where("class_id = ?", classID).Find(&calendars).Error; err != nil {
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


func GetClassesByUserId(c *gin.Context) {
	userID := c.Param("id")

	var calendars []models.Calendar
	if err := database.DB.
		Preload("Class.Gym").
		Preload("Class.Discipline").
		Where("user_id = ?", userID).
		Find(&calendars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar clases"})
		return
	}

	if len(calendars) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron clases para este usuario"})
		return
	}

	var response []dto.ClassResponseCapacity
	for _, calendar := range calendars {
		var count int64
		database.DB.Model(&models.Calendar{}).
			Where("class_id = ?", calendar.Class.Id). 
			Count(&count)

		response = append(response, transformers.TransformClassCapacity(calendar.Class, int(count)))
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

	// Verificar si ya está inscripto
	var existing models.Calendar
	if err := database.DB.
		Where("user_id = ? AND class_id = ?", calendar.UserId, calendar.ClassId).
		First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya está anotado en esta clase"})
		return
	}

	// Contar alumnos de una clase
	var count int64
	database.DB.Model(&models.Calendar{}).Where("class_id = ?", class.Id).Count(&count)

	if int(count) >= class.Capacity {
		c.JSON(http.StatusForbidden, gin.H{"error": "La clase ya está completa"})
		return
	}

	

	// ✅ Validar créditos disponibles
	today := time.Now()
	packUsage, err := services.GetPackUsage(calendar.UserId, class.GymId, class.DisciplineId, today)
	if err != nil || packUsage.Remaining <= 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "No hay créditos disponibles en el pack activo"})
		return
	}

	// Crear inscripción
	database.DB.Create(&calendar)
	database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").First(&calendar, calendar.Id)

	response := transformers.TransformCalendar(calendar/*, packUsage*/, nil)

	c.JSON(http.StatusOK, gin.H{
		"message": "Inscripción exitosa",
		"data":    response,
	})
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

	response := transformers.TransformCalendar(calendar, nil)
	c.JSON(http.StatusOK, response)
}


func CancelClassEnrollment(c *gin.Context) {
	var calendar models.Calendar

	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Buscar inscripción por user_id y class_id
	if err := database.DB.
		Where("user_id = ? AND class_id = ?", calendar.UserId, calendar.ClassId).
		First(&calendar).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró una inscripción para cancelar"})
		return
	}

	if err := database.DB.Delete(&calendar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo cancelar la inscripción"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscripción cancelada con éxito"})
}