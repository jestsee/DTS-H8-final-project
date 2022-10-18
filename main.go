package main

import (
	"log"
	"myGram/config"
	"myGram/controllers"
	"myGram/database"
	"myGram/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	db := database.ConnectDB(&conf)
	inDB := &controllers.InDB{DB: db, Conf: &conf}
	router := gin.Default()

	router.POST("api/v1/register", inDB.Register)
	router.POST("api/v1/login", inDB.Login)

	r := router.Group("api/v1")
	{
		r.Use(middleware.Authentication(*inDB.Conf))
		r.PUT("/users", inDB.UpdateUser)
		r.DELETE("/users", inDB.DeleteUser)
		
		r.GET("/photos", inDB.GetPhotos)
		r.POST("/photos", inDB.AddPhoto)
		r.PUT("/photos/:photoId", inDB.UpdatePhoto)
		r.DELETE("/photos/:photoId", inDB.DeletePhoto)

		r.GET("/comments", inDB.GetComments)
		r.POST("/comments", inDB.AddComment)
		r.PUT("/comments/:commentId", inDB.UpdateComment)
		r.DELETE("/comments/:commentId", inDB.DeleteComment)
		
		r.GET("/socialmedias", inDB.GetSocialMedias)
		r.POST("/socialmedias", inDB.AddSocialMedia)
		r.PUT("/socialmedias/:socialMediaId", inDB.UpdateSocialMedia)
		r.DELETE("/socialmedias/:socialMediaId", inDB.DeleteSocialMedia)
	}
	
	router.Run(":3000")
}
