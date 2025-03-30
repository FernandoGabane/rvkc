package routes

import (
	"rvkc/models"
	"rvkc/repositories"
	"rvkc/service"
	"github.com/gin-gonic/gin"
)


func SetupClubRouter(engine *gin.Engine) {
	repositoryClub 		:= repositories.NewGenericRepository[models.Club]()
	genericServiceClub  := service.NewGenericService(repositoryClub)
	serviceClub		    := service.NewClubClubService(*genericServiceClub)

	clubsRoutes    		:= engine.Group("/clubs")
	{
		clubsRoutes.POST("/", serviceClub.CreateClub)
		clubsRoutes.GET("/", serviceClub.GetClubs)
		clubsRoutes.PUT("/:id", serviceClub.UpdateClub)
		clubsRoutes.GET("/:id", serviceClub.GetClub)
	}
}
