package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

func main() {
	// initialize the database
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// add the routes
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	api.GET("/jokes/", JokeHandler)
	api.POST("/jokes/:newJoke", CreateJoke)
	api.POST("/jokes/like/:jokeID", LikeJoke)
	api.PATCH("/jokes/update/:jokeID/:newJoke", UpdateJoke)
	api.DELETE("/jokes/delete/:jokeID", DeleteJoke)
	router.Run(":3000")
}
