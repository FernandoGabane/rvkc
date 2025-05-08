package routes

import (
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/service"

	"github.com/gin-gonic/gin"
)

func SetupConfirmationRouter(engine *gin.Engine) {
	roleRepository        := repositories.NewGenericRepository[models.Role]()
	roleServiceGeneric 	  := service.NewGenericService(roleRepository)
	roleService  	      := service.NewRoleService(*roleServiceGeneric)

	accountRepository := repositories.NewGenericRepository[models.Account]()
	accountServiceGeneric := service.NewGenericService(accountRepository)
	accountService := service.NewAccountService(*accountServiceGeneric, *roleService)

	clubRepository := repositories.NewGenericRepository[models.Club]()
	clubServiceGeneric := service.NewGenericService(clubRepository)
	clubService := service.NewClubService(*clubServiceGeneric, *accountService)

	confirmationRepository := repositories.NewGenericRepository[models.Confirmation]()
	confirmationServiceGeneric := service.NewGenericService(confirmationRepository)
	confirmationService := service.NewConfirmationService(*confirmationServiceGeneric, *clubService, *accountService)
	confirmationRoutes := engine.Group("/confirmations")
	{
		confirmationRoutes.POST("", confirmationService.CreateConfirmation)
		confirmationRoutes.GET("", confirmationService.GetConfirmationByQueryParam)
		confirmationRoutes.GET("/:id", confirmationService.GetConfirmation)
	}
}
