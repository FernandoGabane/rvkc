package routes

import (
	"rvkc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())
	r.Use(middleware.CORSMiddleware())

	SetupAccountsRouter(r)
	SetupClubRouter(r)
	SetupConfirmationRouter(r)

	return r
}
