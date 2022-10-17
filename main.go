package main

import (
	"log"
	"myGram/config"
	"myGram/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	db := config.ConnectDB(&conf)
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()

	router.POST("/register", inDB.Register)
	router.POST("/login", inDB.Login)
	router.Run(":3000")
}
