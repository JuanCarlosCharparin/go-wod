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
		response = append(response, transformers.TransformCalendar(cal, nil, cal.CreatedAt, nil))
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
	response := transformers.TransformCalendar(calendar, nil, calendar.CreatedAt, nil)
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


func GetInfoByClassId(c *gin.Context) {
    classID := c.Param("id")

    // Subconsulta: obtener el id del último registro por usuario para esta clase
    sub := database.DB.
        Model(&models.Calendar{}).
        Select("MAX(id)").
        Where("class_id = ? AND (status = ? OR status = ?)", classID, "inscripto", "ausente").
        Group("user_id")

    var calendars []models.Calendar
    if err := database.DB.
        Preload("User.Gym").
        Preload("Class.Discipline").
        Where("id IN (?)", sub).
        Find(&calendars).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar información de la clase"})
        return
    }

    if len(calendars) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró información para esta clase"})
        return
    }

    var response []dto.CalendarResponse
    for _, cal := range calendars {
        response = append(response, transformers.TransformCalendar(cal, nil, cal.CreatedAt, &cal.UpdatedAt))
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

	var response []dto.ClassWithStatusResponse
	for _, calendar := range calendars {
		response = append(response, transformers.TransformClassWithStatus(calendar.Class, calendar.Status, calendar.CreatedAt))
	}

	c.JSON(http.StatusOK, response)
}





func CreateCalendar(c *gin.Context) {
	var calendar models.Calendar

	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que el usuario exista
	var user models.User
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", calendar.UserId).
		First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no válido o eliminado"})
		return
	}

	// Validar que la clase exista
	var class models.Class
	if err := database.DB.
		Where("id = ? AND deleted_at IS NULL", calendar.ClassId).
		First(&class).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Clase no válida o eliminada"})
		return
	}

	// Verificar si ya está inscripto en esta clase con estado "inscripto"
	var existing models.Calendar
	if err := database.DB.
		Where("user_id = ? AND class_id = ? AND status = ?", calendar.UserId, calendar.ClassId, "inscripto").
		First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "El usuario ya está anotado en esta clase"})
		return
	}

	// Validar créditos disponibles
	today := time.Now()
	   packUsage, err := services.GetPackUsage(calendar.UserId, class.GymId, []uint{class.DisciplineId}, today)
	if err != nil || packUsage.Remaining <= 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "No hay créditos disponibles en el pack activo"})
		return
	}

	// Verificar si el usuario tiene al menos una disciplina que coincida con la disciplina de la clase
	var userPacks []models.UserPack
	err = database.DB.
		Preload("Disciplines").
		Where("user_id = ? AND status = 1 AND ? BETWEEN start_date AND expiration_date", calendar.UserId, today).
		Find(&userPacks).Error

	if err != nil || len(userPacks) == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "No tenés packs activos con disciplinas válidas"})
		return
	}

	// Buscar si alguna disciplina del pack coincide con la de la clase
	disciplinaValida := false
	for _, pack := range userPacks {
		for _, d := range pack.Disciplines {
			if d.DisciplineId == class.DisciplineId {
				disciplinaValida = true
				break
			}
		}
		if disciplinaValida {
			break
		}
	}

	if !disciplinaValida {
		c.JSON(http.StatusForbidden, gin.H{"error": "Tu pack no incluye la disciplina de esta clase"})
		return
	}

	// Contar inscriptos actuales
	var count int64
	database.DB.Model(&models.Calendar{}).
		Where("class_id = ? AND status = ?", class.Id, "inscripto").
		Count(&count)

	if int(count) >= class.Capacity {
		// Clase llena → inscribir en waitlist
		wait := models.Waitlist{
			UserId:  calendar.UserId,
			ClassId: calendar.ClassId,
		}

		// Verificar si ya está en la lista
		var existingWait models.Waitlist
		if err := database.DB.
			Where("user_id = ? AND class_id = ?", wait.UserId, wait.ClassId).
			First(&existingWait).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Ya estás en la lista de espera de esta clase"})
			return
		}

		database.DB.Create(&wait)

		c.JSON(http.StatusOK, gin.H{
			"message": "Clase llena. Te has anotado en la lista de espera.",
			"waitlist": true,
		})
		return
	}

	// Establecer el estado de inscripción como "inscripto"
	calendar.Status = "inscripto"

	// Guardar inscripción
	if err := database.DB.Create(&calendar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar la inscripción"})
		return
	}

	// Obtener datos completos para el response
	database.DB.Preload("User").Preload("Class.Gym").Preload("Class.Discipline").First(&calendar, calendar.Id)
	response := transformers.TransformCalendar(calendar, packUsage, calendar.CreatedAt, nil)

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

	response := transformers.TransformCalendar(calendar, nil, calendar.CreatedAt, nil)
	c.JSON(http.StatusOK, response)
}


func CancelClassEnrollment(c *gin.Context) {
	var calendar models.Calendar

	// 1. Bind JSON de la request al modelo Calendar
	if err := c.ShouldBindJSON(&calendar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// 2. Buscar inscripción que coincida con user_id y class_id
	// Además, precargar la clase y gimnasio para obtener detalles
	if err := database.DB.
		Preload("Class").
		Preload("Class.Gym").
		Where("user_id = ? AND class_id = ? and status = 'inscripto'", calendar.UserId, calendar.ClassId).
		First(&calendar).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró una inscripción para cancelar"})
		return
	}

	// 4. Parsear la fecha (string con zona horaria)
	const layoutDateTime = "2006-01-02T15:04:05-07:00"
	dateParsed, err := time.Parse(layoutDateTime, calendar.Class.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parseando fecha de clase"})
		return
	}

	// 5. Parsear la hora (formato HH:MM:SS)
	timeParsed, err := time.Parse("15:04:05", calendar.Class.Time)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parseando hora de clase"})
		return
	}

	// 6. Combinar fecha y hora en un solo time.Time
	classTime := time.Date(
		dateParsed.Year(), dateParsed.Month(), dateParsed.Day(),
		timeParsed.Hour(), timeParsed.Minute(), timeParsed.Second(), 0,
		dateParsed.Location(),
	)

	// 7. Obtener límite de cancelación configurado para el gimnasio
	limitMinutes, err := services.GetGymCancelTimeLimit(calendar.Class.GymId)
	if err != nil {
		// Si falla la consulta o no existe configuración, ponemos 60 minutos por defecto
		limitMinutes = 60
	}

	// 8. Calcular la diferencia entre hora de clase y hora actual
	now := time.Now()
	diff := classTime.Sub(now)

	// 9. Si estamos dentro del límite (menos minutos que el límite), marcar como ausente
	//    Si no, marcar como cancelado (sin penalización)
	if diff.Minutes() < float64(limitMinutes) {
		calendar.Status = "ausente"
	} else {
		calendar.Status = "cancelado"
	}

	// 10. Guardar el cambio de estado en la base de datos (actualizar el registro)
	if err := database.DB.Save(&calendar).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo cancelar la inscripción"})
		return
	}

	// 11. Responder con éxito y el estado final de la inscripción
	c.JSON(http.StatusOK, gin.H{
		"message": "Inscripción cancelada con éxito",
		"status":  calendar.Status,
	})
}




func GetUserUsedClasses(c *gin.Context) {
	userID := c.Param("id")

	var userPack models.UserPack
	err := database.DB.Preload("Pack").Preload("Disciplines").
			Where("user_id = ? AND status = 1",
					userID).
			First(&userPack).Error
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró un pack activo para el usuario"})
			return
	}

	// Extraer los IDs de disciplina del pack
	disciplineIDs := make([]uint, 0, len(userPack.Disciplines))
	for _, d := range userPack.Disciplines {
			disciplineIDs = append(disciplineIDs, d.DisciplineId)
	}

	// Contar clases usadas para todas las disciplinas del pack
	used, err := services.CountUsedClasses(userPack.UserId, userPack.GymId, disciplineIDs, userPack.StartDate, userPack.ExpirationDate)
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar clases usadas"})
			return
	}

	// Armar pack usage
	packUsage := &services.PackUsage{
		PackID:        userPack.PackId,
		ClassQuantity: userPack.Pack.ClassQuantity,
		Used:          used,
		Remaining:     userPack.Pack.ClassQuantity - used,
	}

	response := packUsage

	c.JSON(http.StatusOK, response)
}




func GetUserPacksUsage(c *gin.Context) {
	userID := c.Param("id")
	
	statusParam := c.DefaultQuery("status", "all") // all | 1 | 0

	   db := database.DB.
			   Preload("Pack").
			   Preload("Disciplines.Discipline").
			   Preload("Gym").
			   Where("user_id = ?", userID)

	switch statusParam {
	case "1", "active", "activo":
		db = db.Where("status = 1")
	case "0", "inactive", "inactivo", "vencido":
		db = db.Where("status = 0")
	// "all" → no agregamos filtro
	}

	var userPacks []models.UserPack
	if err := db.Order("start_date DESC").Find(&userPacks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener packs del usuario"})
		return
	}

	if len(userPacks) == 0 {
		// vacía pero 200, útil para frontend
		c.JSON(http.StatusOK, []dto.UserPackUsageItem{})
		return
	}

	// construir respuesta
	items := make([]dto.UserPackUsageItem, 0, len(userPacks))
	   for _, up := range userPacks {
			   // Extraer los IDs de disciplina del pack
			   disciplineIDs := make([]uint, 0, len(up.Disciplines))
			   disciplinesResp := make([]dto.DisciplineResponse, 0, len(up.Disciplines))
			   for _, d := range up.Disciplines {
					   disciplineIDs = append(disciplineIDs, d.DisciplineId)
					   disciplinesResp = append(disciplinesResp, dto.DisciplineResponse{
							   ID:   d.DisciplineId,
							   Name: d.Discipline.Name,
					   })
			   }

			   used, err := services.CountUsedClasses(
					   up.UserId,
					   up.GymId,
					   disciplineIDs,
					   up.StartDate,
					   up.ExpirationDate,
			   )
			   if err != nil {
					   // si falla el conteo, seguimos pero marcamos error local
					   used = 0
			   }
			   remaining := up.Pack.ClassQuantity - used
			   if remaining < 0 {
					   remaining = 0
			   }

			   items = append(items, dto.UserPackUsageItem{
					   UserPackID:     up.Id,
					   Status:         up.Status,
					   StartDate:      up.StartDate.Format("2006-01-02"),
					   ExpirationDate: up.ExpirationDate.Format("2006-01-02"),
					   Used:           used,
					   Remaining:      remaining,
					   ClassQuantity:  up.Pack.ClassQuantity,
					   Pack: dto.PackResponseMin{
							   ID:       up.PackId,
							   PackName: up.Pack.PackName,
							   Price:    up.Pack.Price, // ajustá tipo
					   },
					   Disciplines: disciplinesResp,
					   Gym: dto.GymResponseMin{
							   ID:   up.GymId,
							   Name: up.Gym.Name,
					   },
			   })
	   }

	c.JSON(http.StatusOK, items)
}
