package main

import (
	"wod-go/models"
	"wod-go/handlers"
	"wod-go/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// AutoMigrar la tabla si no existe
	database.DB.AutoMigrate(&models.Country{}, &models.Gym{})

	r := gin.Default()

	//countries
	r.GET("/countries", handlers.GetCountries)
	r.GET("/countries/:id", handlers.GetCountryId)
	r.POST("/countries", handlers.CreateCountry)
	r.PUT("/countries/:id", handlers.UpdatedCountry)
	r.DELETE("/countries/:id", handlers.DeleteCountry)

	//gyms
	r.GET("/gyms", handlers.GetGyms)
	r.GET("/gyms/:id", handlers.GetGymId)
	r.POST("/gyms", handlers.CreateGym)
	r.PUT("/gyms/:id", handlers.UpdatedGym)
	r.DELETE("/gyms/:id", handlers.DeleteGym)


	//users
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserId)
	r.POST("/users", handlers.CreateUser)
	/*r.PUT("/users/:id", handlers.UpdatedGym)
	r.DELETE("/users/:id", handlers.DeleteGym)*/

	r.Run(":8080")
}
