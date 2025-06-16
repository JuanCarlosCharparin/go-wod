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
	database.DB.AutoMigrate(
		&models.Country{}, 
		&models.Gym{}, 
		&models.User{}, 
		&models.Discipline{}, 
		&models.Wod{}, 
		&models.Calendar{},
		&models.Class{},
		&models.Pack{},
		&models.UserPack{},
	)

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
	r.PUT("/users/:id", handlers.UpdatedUser)
	r.DELETE("/users/:id", handlers.DeleteUser)


	//disciplines
	r.GET("/disciplines", handlers.GetDisciplines)
	r.GET("/disciplines/:id", handlers.GetDisciplineId)
	r.POST("/disciplines", handlers.CreateDiscipline)
	r.PUT("/disciplines/:id", handlers.UpdatedDiscipline)
	r.DELETE("/disciplines/:id", handlers.DeleteDiscipline)


	//wods
	r.GET("/wods", handlers.GetWods)
	r.GET("/wods/:id", handlers.GetWodId)
	r.POST("/wods", handlers.CreateWod)
	r.PUT("/wods/:id", handlers.UpdatedWod)
	r.DELETE("/wods/:id", handlers.DeleteWod)


	//classes
	r.GET("/classes", handlers.GetClasses)
	r.GET("/classes/:id", handlers.GetClassId)
	r.POST("/classes", handlers.CreateClass)
	r.PUT("/classes/:id", handlers.UpdatedClass)
	r.DELETE("/classes/:id", handlers.DeleteClass)


	//calendars
	r.GET("/calendar", handlers.GetCalendars)
	r.GET("/calendar/:id", handlers.GetCalendarId)
	r.POST("/calendar", handlers.CreateCalendar)
	r.PUT("/calendars/:id", handlers.UpdatedCalendar)
	r.DELETE("/calendars/:id", handlers.DeleteCalendar)


	//packs
	r.GET("/packs", handlers.GetPacks)
	r.GET("/packs/:id", handlers.GetPackId)
	r.POST("/packs", handlers.CreatePack)
	r.PUT("/packs/:id", handlers.UpdatedPack)
	r.DELETE("/packs/:id", handlers.DeletePack)


	//user_packs
	r.GET("/user_packs", handlers.GetUserPacks)
	r.GET("/user_packs/:id", handlers.GetUserPackId)
	r.POST("/user_packs", handlers.CreateUserPack)
	r.PUT("/user_packs/:id", handlers.UpdatedUserPack)
	r.DELETE("/user_packs/:id", handlers.DeleteUserPack)

	r.Run(":8080")
}
