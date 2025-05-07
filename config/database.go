package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"rvkc/middleware"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connectionUrl := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + dbname + " port=" + dbPort + " sslmode=disable"

	database, err := gorm.Open(postgres.Open(connectionUrl), &gorm.Config{
		Logger: &middleware.GormLogger{},
	})
	
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados:", err)
	}

	DB = database
}
