package routes

import (
	"net/http"
	"rvkc/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())
	r.Use(middleware.CORSMiddleware())

	r.Static("/static", "./static")
	r.LoadHTMLFiles("static/home.html", "static/admin.html")
	
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin.html", nil)
	})

	SetupAccountsRouter(r)
	SetupClubRouter(r)
	SetupConfirmationRouter(r)

	return r
}