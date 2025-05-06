package routes

import (
	"net/http"
	"rvkc/middleware"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.Static("/static", "./static")
	r.LoadHTMLFiles("static/home.html")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", nil)
	})

	SetupAccountsRouter(r)
	SetupClubRouter(r)
	SetupConfirmationRouter(r)

	return r
}