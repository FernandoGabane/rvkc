package main


import (
    "rvkc/config"
    "rvkc/routes"
)


func main() {
	config.ConnectDatabase()

    r := routes.SetupRouter()

	r.Run(":8080")
}
