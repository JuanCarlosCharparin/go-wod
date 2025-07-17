package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wod-go/database"
	"wod-go/handlers"
	"wod-go/jobs"
	"wod-go/middleware"
	"wod-go/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
}

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
		&models.ScheduleTemplate{}, 
		&models.ScheduleBlock{},
		&models.Role{},
		&models.GymSetting{},
		&models.Waitlist{},
	)

	//r := gin.Default()

	//cors
	r := gin.New() // ← NO usar gin.Default() si aplicás middlewares manuales
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Configuración explícita de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rutas protegidas
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		/*protected.GET("/user-packs", handlers.GetUserPacks)
		protected.POST("/classes", handlers.CreateClass)*/
		// etc...
	}

	// Iniciar el cron de busqueda de packs vencidos
	c := cron.New()
	c.AddFunc("@daily", jobs.CheckExpiredUserPacks) 
	c.Start()


	//
	args := os.Args

	if len(args) > 1 && args[1] == "check-expired-packs" {
		fmt.Println("→ Ejecutando job de verificación de packs expirados")
		jobs.CheckExpiredUserPacks()
		return
	}



	//rutas protegidas login
	r.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)
		c.JSON(http.StatusOK, gin.H{
			"message": "Acceso autorizado",
			"user_id": userID,
		})
	})

	//countries
	r.GET("/countries", handlers.GetCountries)
	r.GET("/countries/:id", handlers.GetCountryId)
	r.POST("/countries", handlers.CreateCountry)
	r.PUT("/countries/:id", handlers.UpdatedCountry)
	r.DELETE("/countries/:id", handlers.DeleteCountry)

	//gyms
	r.GET("/gyms", handlers.GetGyms)
	r.GET("/gyms/:id", handlers.GetGymId)
	r.GET("/gyms/country/:id", handlers.GetGymsByCountryId)
	r.POST("/gyms", handlers.CreateGym)
	r.PUT("/gyms/:id", handlers.UpdatedGym)
	r.DELETE("/gyms/:id", handlers.DeleteGym)


	//users
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserId)
	r.GET("/users/gym/:id/:rol", handlers.GetUsersByGymId)
	r.POST("/users", handlers.CreateUser)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.LoginUser)
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
	r.GET("/wods/gym/:id", handlers.GetWodByGymId)
	r.POST("/wods", handlers.CreateWod)
	r.PUT("/wods/:id", handlers.UpdatedWod)
	r.DELETE("/wods/:id", handlers.DeleteWod)


	//classes
	r.GET("/classes", handlers.GetClasses)
	r.GET("/classes/:id", handlers.GetClassId)
	r.GET("/classes/gym/:id", handlers.GetClassesByGymId)
	r.GET("/calendar/gym/:id/upcoming", handlers.GetUpcomingClassesByGymId) //clases futuras
	r.POST("/classes", handlers.CreateClass)
	r.PUT("/classes/:id", handlers.UpdatedClass)
	r.DELETE("/classes/:id", handlers.DeleteClass)
	/* generar clases*/
	r.POST("/generate-classes", handlers.GenerateClassesHandler)


	//calendars
	r.GET("/calendar", handlers.GetCalendars)
	r.GET("/calendar/:id", handlers.GetCalendarId)
	r.GET("/calendar/class/:id", handlers.GetUsersByClassId)
	r.GET("/calendar/users/:id", handlers.GetClassesByUserId)
	r.POST("/calendar", handlers.CreateCalendar)
	r.PUT("/calendars/:id", handlers.UpdatedCalendar)
	r.DELETE("/calendar", handlers.CancelClassEnrollment)
	//clases usadas y restantes por usuario
	r.GET("/calendar/users/used-classes/:id", handlers.GetUserUsedClasses)


	//packs
	r.GET("/packs", handlers.GetPacks)
	r.GET("/packs/:id", handlers.GetPackId)
	r.GET("/packs/gym/:id", handlers.GetPackByGymId)
	r.POST("/packs", handlers.CreatePack)
	r.PUT("/packs/:id", handlers.UpdatedPack)
	r.DELETE("/packs/:id", handlers.DeletePack)


	//user_packs
	r.GET("/user_packs", handlers.GetUserPacks)
	r.GET("/user_packs/:id", handlers.GetUserPackId)
	r.GET("/user_packs/user/:id", handlers.GetUserPackByUserId)
	r.POST("/user_packs", handlers.CreateUserPack)
	r.PUT("/user_packs/:id", handlers.UpdatedUserPack)
	r.DELETE("/user_packs/:id", handlers.DeleteUserPack)


	//schedules
	r.GET("/templates/gym/:id", handlers.GetScheduleTemplatesByGymID)
	r.POST("/schedule-templates", handlers.CreateScheduleTemplate)
	r.POST("/schedule-blocks", handlers.CreateScheduleBlock)

	//roles
	r.GET("/roles", handlers.GetRoles)
	r.POST("/roles", handlers.CreateRole)

	//gym_settings
	r.GET("/settings", handlers.GetSettings)


	// Levantar la API normalmente
	r.Run(":8080")
}
