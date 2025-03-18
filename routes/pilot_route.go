package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"rvkc/repositories"
	"rvkc/services"
	"rvkc/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite qualquer origem (pode restringir para domínios específicos)
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	repo := repositories.NewPilotRepository()
	service := services.NewPilotService(repo)
	pilotController := controllers.NewPilotController(service)

	pilotRoutes := r.Group("/pilots")
	{
		pilotRoutes.POST("/", pilotController.CreatePilot)   // Criar um piloto
		pilotRoutes.GET("/", pilotController.GetPilots)      // Listar todos os pilotos
		pilotRoutes.GET("/:document", pilotController.GetPilotByDocument) // Buscar piloto por ID
		pilotRoutes.PUT("/", pilotController.UpdatePilot)  // Atualizar piloto
		// pilotRoutes.DELETE("/:document", pilotController.DeletePilot) // Deletar piloto
	}

	return r
}
