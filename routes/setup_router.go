package routes

import (
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

	SetupAccountsRouter(r)
	SetupClubRouter(r)

	return r
}