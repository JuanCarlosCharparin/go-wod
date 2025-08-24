package handlers

import (
	"net/http"
	"strconv"
	"time"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/services"
	"wod-go/transformers"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []models.Class
	database.DB.Preload("Gym.Country").Preload("Discipline").Find(&classes)
	var response []dto.ClassResponse
	for _, cal := range classes {
		response = append(response, transformers.TransformClass(cal))
	}
	c.JSON(http.StatusOK, response)
}

func GetClassId(c *gin.Context) {

	id := c.Param("id")

	var class models.Class
	if err := database.DB.Preload("Gym").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	response := transformers.TransformClass(class)
	c.JSON(http.StatusOK, response)
}


func GetClassesByGymId(c *gin.Context) {
	gymID := c.Param("id")

	var classes []models.Class
	if err := database.DB.Preload("Gym").Preload("Discipline").Where("gym_id = ?", gymID).Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar gimnasios"})
		return
	}

	if len(classes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron clases para este gimnasio"})
		return
	}

	var response []dto.ClassResponseInfo
	for _, class := range classes {
		var count int64
		database.DB.Model(&models.Calendar{}).
			Where("class_id = ?", class.Id).
			Count(&count)

		response = append(response, transformers.TransformClassInfo(class, int(count)))
	}

	c.JSON(http.StatusOK, response)
}


func GetUpcomingClassesByGymId(c *gin.Context) {
	gymID := c.Param("id")

	var classes []models.Class
	err := database.DB.
		Preload("Gym").
		Preload("Discipline").
		Where("gym_id = ?", gymID).
		Find(&classes).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar clases"})
		return
	}

	var response []dto.ClassResponseInfo
	now := time.Now()

	for _, class := range classes {
		// Parse date y time
		dateParsed, errDate := time.Parse(time.RFC3339, class.Date)
		timeParsed, errTime := time.Parse("15:04:05", class.Time)

		if errDate != nil || errTime != nil {
			continue // ignorar clases con formato incorrecto
		}

		// Combinar date + time
		classDateTime := time.Date(
			dateParsed.Year(), dateParsed.Month(), dateParsed.Day(),
			timeParsed.Hour(), timeParsed.Minute(), timeParsed.Second(),
			0, dateParsed.Location(),
		)

		// Filtrar clases futuras
		if classDateTime.After(now) {
			var count int64
			database.DB.Model(&models.Calendar{}).
				Where("class_id = ?", class.Id).
				Count(&count)

			response = append(response, transformers.TransformClassInfo(class, int(count)))
		}
	}

	if len(response) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No hay clases futuras para este gimnasio"})
		return
	}

	c.JSON(http.StatusOK, response)
}


// GetClassesOnWeekByGymId obtiene todas las clases de la semana actual para un gimnasio
func GetClassesOnWeekByGymId(c *gin.Context) {
	gymID := c.Param("id")

	// offset opcional enviado desde el frontend (en semanas)
	offsetStr := c.Query("offset") // puede ser "-1", "0", "1"...
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // por defecto la semana actual
	}

	now := time.Now().AddDate(0, 0, offset*7) // desplazamos la fecha base N semanas

	// Calcular lunes y domingo
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	monday := now.AddDate(0, 0, -weekday+1).Truncate(24 * time.Hour)
	sunday := monday.AddDate(0, 0, 6).Truncate(24 * time.Hour)

	// Query a la BD
	var classes []models.Class
	if err := database.DB.Preload("Gym").Preload("Discipline").
		Where("gym_id = ? AND date BETWEEN ? AND ?", gymID, monday, sunday).
		Order("date ASC").
		Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las clases."})
		return
	}

	var transformedClasses []interface{}
	for _, class := range classes {
		var count int64
		database.DB.Model(&models.Calendar{}).
			Where("class_id = ? and status = ?", class.Id, "inscripto").
			Count(&count)
		transformedClasses = append(transformedClasses, transformers.TransformClassInfo(class, int(count)))
	}

	c.JSON(http.StatusOK, gin.H{
		"week_start": monday.Format("2006-01-02"),
		"week_end":   sunday.Format("2006-01-02"),
		"classes":    transformedClasses,
	})
}

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Validar que el gimnasio exista y no esté eliminado
	var gym models.Gym
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", class.GymId).
		First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}

	//Validar que la disciplina exista y no esté eliminada
	var discipline models.Discipline
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", class.DisciplineId).
		First(&discipline).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Disciplina no válida o eliminada"})
		return
	}
	
	database.DB.Create(&class)
	database.DB.Preload("Gym").Preload("Discipline").First(&class, class.Id)
	response := transformers.TransformClass(class)
	c.JSON(http.StatusOK, response)
}


func UpdatedClass(c *gin.Context) {
	id := c.Param("id")

	var class models.Class
	if err := database.DB.Preload("Gym").Preload("Discipline").First(&class, id).Error; err != nil {
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
	if err := database.DB.Preload("Gym").Preload("Discipline").First(&class, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cargar el gimnasio, país y disciplina"})
		return
	}

	response := transformers.TransformClass(class)
	c.JSON(http.StatusOK, response)
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
}


//Generar clases de manera automatica

func GenerateClassesHandler(c *gin.Context) {
	var body struct {
		GymID uint   `json:"gym_id"`
		From  string `json:"from"`
		To    string `json:"to"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	fromDate, err1 := time.Parse("2006-01-02", body.From)
	toDate, err2 := time.Parse("2006-01-02", body.To)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de fechas inválido"})
		return
	}

	if err := services.GenerateClassesFromTemplates(body.GymID, fromDate, toDate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando clases"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clases generadas exitosamente"})
}
