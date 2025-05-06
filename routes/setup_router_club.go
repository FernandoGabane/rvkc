package routes

import (
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/service"
	"github.com/gin-gonic/gin"
)


func SetupClubRouter(engine *gin.Engine) {
	roleRepository        := repositories.NewGenericRepository[models.Role]()
	roleServiceGeneric 	  := service.NewGenericService(roleRepository)
	roleService  	      := service.NewRoleService(*roleServiceGeneric)

	accountRepository     := repositories.NewGenericRepository[models.Account]()
	accountServiceGeneric := service.NewGenericService(accountRepository)
	accountService 	      := service.NewAccountService(*accountServiceGeneric, *roleService)

	clubRepository 		  := repositories.NewGenericRepository[models.Club]()
	clubServiceGeneric    := service.NewGenericService(clubRepository)
	clubService		      := service.NewClubService(*clubServiceGeneric, *accountService)

	clubsRoutes    		  := engine.Group("/clubs")
	{
		clubsRoutes.POST("", 	clubService.CreateClub)
		clubsRoutes.GET("", 	clubService.GetClubs)
		clubsRoutes.PUT("/:id", clubService.UpdateClub)
		clubsRoutes.GET("/:id", clubService.GetClub)
	}
}