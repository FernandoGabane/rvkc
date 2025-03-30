package routes

import (
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/service"
	"github.com/gin-gonic/gin"
)


func SetupAccountsRouter(engine *gin.Engine) {
	repositoryPilot   	:= repositories.NewGenericRepository[models.Account]()
	genericServicePilot := service.NewGenericService(repositoryPilot)
	serviceAccount 	    := service.NewAccountService(*genericServicePilot)

	accountsRoutes := engine.Group("/accounts")
	{
		accountsRoutes.POST("/", serviceAccount.CreatePilot)
		accountsRoutes.GET("/", serviceAccount.GetPilots)
		accountsRoutes.PUT("/", serviceAccount.UpdatePilot)
		accountsRoutes.GET("/:document", serviceAccount.GetPilotByDocument)
	}
}
