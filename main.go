package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

type Joke struct {
	ID     int     `json:"id" binding:"required"`
	Likes  int     `json:"likes"`
	Joke   string  `json:"joke" binding:"required"`
}

var jokes = []Joke{
	Joke{1, 0, "Hello world 1."},
	Joke{2, 0, "Hello world 2."},
	Joke{3, 0, "Hello world 3."},
}

func main() {
  router := gin.Default()
  router.Use(static.Serve("/", static.LocalFile("./views", true)))
  api := router.Group("/api")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })
  }
  api.GET("/jokes", JokeHandler)
  api.POST("/jokes/like/:jokeID", LikeJoke)
  api.PATCH("/jokes/update/:jokeID/:newJoke", UpdateJoke)
  api.DELETE("/jokes/delete/:jokeID", DeleteJoke)
  router.Run(":3000")
}

func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

func LikeJoke(c *gin.Context) {
	if jokeId, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeId {
				jokes[i].Likes += 1
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateJoke(c *gin.Context) {
	if jokeId, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		newJoke := c.Param("newJoke")
		jokes[jokeId].Joke = newJoke
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}	
	
}

func DeleteJoke(c *gin.Context) {
	if jokeId, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		jokes = append(jokes[:jokeId], jokes[jokeId + 1])
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}