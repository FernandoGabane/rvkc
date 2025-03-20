package routes

import (
	"rvkc/controllers"
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite qualquer origem
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))


	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})


	repositoryPilot := repositories.NewGenericRepository[models.Pilot]()
	servicePilot := services.NewGenericService(repositoryPilot)
	pilotController := controllers.NewPilotController(*servicePilot)

	repositoryClub := repositories.NewGenericRepository[models.Club]()
	serviceClub := services.NewGenericService(repositoryClub)
	clubController := controllers.NewClubController(*serviceClub)


	pilotRoutes := r.Group("/pilots")
	{
		pilotRoutes.POST("/", pilotController.CreatePilot)
		pilotRoutes.GET("/", pilotController.GetPilots)
		pilotRoutes.GET("/:document", pilotController.GetPilotByDocument)
		pilotRoutes.PUT("/", pilotController.UpdatePilot)
	}


	clubRoutes := r.Group("/clubs")
	{
		clubRoutes.POST("/", clubController.CreateClub)
		clubRoutes.GET("/", clubController.GetClubs)
		clubRoutes.GET("/:id", clubController.GetClub)
		clubRoutes.PUT("/", clubController.UpdateClub)
	}

	return r
}
