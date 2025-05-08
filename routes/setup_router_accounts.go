package routes

import (
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/service"
	"github.com/gin-gonic/gin"
)


func SetupAccountsRouter(engine *gin.Engine) {
	roleRepository        := repositories.NewGenericRepository[models.Role]()
	roleServiceGeneric 	  := service.NewGenericService(roleRepository)
	roleService  	      := service.NewRoleService(*roleServiceGeneric)

	accountRepository     := repositories.NewGenericRepository[models.Account]()
	accountServiceGeneric := service.NewGenericService(accountRepository)
	accountService 	      := service.NewAccountService(*accountServiceGeneric, *roleService)

	accountsRoutes := engine.Group("/accounts")
	{
		accountsRoutes.POST("", accountService.CreateAccount)
		accountsRoutes.GET("", accountService.GetAccounts)
		// accountsRoutes.GET("/:document", accountService.GetAccountByDocument) disable, enable it later if needed
		accountsRoutes.GET("/:document/simple", accountService.GetAccountSimpleByDocument)
		accountsRoutes.GET("/simple", accountService.GetAccountsSimple)
		accountsRoutes.PUT("/", accountService.UpdateAccount)
	}
}