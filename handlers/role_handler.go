package handlers

import (
	"wod-go/database"
	"wod-go/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role
	database.DB.Find(&roles)

	c.JSON(http.StatusOK, roles)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	database.DB.Create(&role)
	database.DB.First(&role, role.Id)
	c.JSON(http.StatusOK, role)
}