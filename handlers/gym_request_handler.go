package handlers

import (
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/transformers"
	"net/http"
	"github.com/gin-gonic/gin"
    "errors"
    "gorm.io/gorm"
)

func RequestGymMembership(c *gin.Context) {
    var req struct {
        UserId uint `json:"user_id" binding:"required"`
        GymId  uint `json:"gym_id" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validar que el usuario y el gimnasio existan
    var user models.User
    if err := database.DB.First(&user, req.UserId).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado"})
        return
    }
    var gym models.Gym
    if err := database.DB.First(&gym, req.GymId).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no encontrado"})
        return
    }

    // Evitar duplicados: ya existe una solicitud pendiente
    var existing models.GymRequest
    if err := database.DB.Where("user_id = ? AND gym_id = ? AND status = ?", req.UserId, req.GymId, "pendiente").First(&existing).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Ya existe una solicitud pendiente para este usuario y gimnasio"})
        return
    }

    // Evitar si usuario ya pertenece al gym
    if user.GymId != nil && *user.GymId == req.GymId {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario ya pertenece a este gimnasio"})
        return
    }

    gymRequest := models.GymRequest{
        UserId: req.UserId,
        GymId:  req.GymId,
        Status: "pendiente",
    }
    if err := database.DB.Create(&gymRequest).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar la solicitud"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Solicitud enviada"})
}


func GetPendingGymRequests(c *gin.Context) {
    gymID := c.Param("id")
    var requests []models.GymRequest
    database.DB.Preload("User").Preload("Gym").Where("gym_id = ? AND status = ?", gymID, "pendiente").Find(&requests)

    response := make([]dto.GymRequestResponse, len(requests))
    for i, req := range requests {
        response[i] = transformers.TransformGymRequest(req)
    }

    c.JSON(http.StatusOK, response)
}


func AcceptGymRequest(c *gin.Context) {
    id := c.Param("id")

    err := database.DB.Transaction(func(tx *gorm.DB) error {
        var req models.GymRequest
        if err := tx.First(&req, id).Error; err != nil {
            return err
        }

        if req.Status != "pendiente" {
            return errors.New("la solicitud no está en estado pendiente")
        }

        // Marcar aceptada
        req.Status = "aceptado"
        if err := tx.Save(&req).Error; err != nil {
            return err
        }

        // Asociar usuario al gimnasio
        if err := tx.Model(&models.User{}).Where("id = ?", req.UserId).Update("gym_id", req.GymId).Error; err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Solicitud no encontrada"})
            return
        }
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuario asociado al gimnasio"})
}



func RejectGymRequest(c *gin.Context) {
    id := c.Param("id")
    var req models.GymRequest
    if err := database.DB.First(&req, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Solicitud no encontrada"})
        return
    }
    if req.Status != "pendiente"{
        c.JSON(http.StatusBadRequest, gin.H{"error": "Solo se pueden rechazar solicitudes en estado pendiente"})
        return
    }
    req.Status = "rechazado"
    database.DB.Save(&req)
    c.JSON(http.StatusOK, gin.H{"message": "Solicitud rechazada"})
}


