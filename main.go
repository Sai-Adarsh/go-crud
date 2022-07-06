package main

import (
	"net/http"

	"context"
	"fmt"
	"go-crud/controllers/createjoke"
	"go-crud/controllers/deletejoke"
	"go-crud/controllers/jokehandler"
	"go-crud/controllers/likejoke"
	"go-crud/controllers/updatejoke"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
)

var ctx = context.Background()

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func main() {
	initRedis()
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
	api.GET("/jokes/", jokehandler.JokeHandler)
	api.POST("/jokes/:newJoke", createjoke.CreateJoke)
	api.POST("/jokes/like/:jokeID", likejoke.LikeJoke)
	api.PATCH("/jokes/update/:jokeID/:newJoke", updatejoke.UpdateJoke)
	api.DELETE("/jokes/delete/:jokeID", deletejoke.DeleteJoke)
	router.Run(":3000")
}
