package main

import (
	"svi-backend/config/database"
	"svi-backend/handler"
	"svi-backend/repository"
	"svi-backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(cors.Default())

	//Database Connection
	conn := database.DatabaseConn()

	//Migration
	database.DatabaseMigration(conn)

	rGlobal := repository.NewRepository(conn)

	// Post instance
	postService := service.PostNewService(rGlobal)
	postHandler := handler.NewPostHandler(postService)

	v1 := router.Group("api/v1")
	{
		grPost := v1.Group("/article")
		{
			grPost.POST("", postHandler.PostArticleHandler)
			grPost.GET("/:id", postHandler.FindByIDBookHandler)
			grPost.POST("/:id", postHandler.UpdateArticleHandler)
			grPost.DELETE("/:id", postHandler.DeleteArticleHandler)
			grPost.GET("/:id/:page", postHandler.FindAllArticleHandler)
		}
	}

	router.Run(":8081")

}
