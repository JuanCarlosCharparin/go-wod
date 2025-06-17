package handlers

import (
	"errors"
	"net/http"
	"os"
	"time"
	"wod-go/database"
	"wod-go/dto"
	"wod-go/models"
	"wod-go/transformers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar que el gimnasio exista
	var gym models.Gym
	if err := database.DB.Where("id = ? AND deleted_at IS NULL", input.GymId).First(&gym).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gimnasio no válido o eliminado"})
		return
	}

	// Verificar que el email no esté registrado
	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El email ya está registrado"})
		return
	}

	// Hashear contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
		return
	}

	// Crear usuario
	user := models.User{
		Name:      input.Name,
		Lastname:  input.Lastname,
		Gender:    input.Gender,
		Phone:     input.Phone,
		Email:     input.Email,
		DNI:       input.DNI,
		BirthDate: input.BirthDate,
		Password:  string(hashedPassword),
		GymId:     input.GymId,
	}

	database.DB.Create(&user)
	database.DB.Preload("Gym").First(&user, user.Id)

	response := transformers.TransformUser(user)
	c.JSON(http.StatusOK, response)
}



func LoginUser(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Verificar contraseña
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
		return
	}

	// Generar JWT
	tokenString, err := GenerateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":    user.Id,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}


func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Leer la clave secreta del .env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("clave JWT no definida")
	}

	return token.SignedString([]byte(secret))
}
