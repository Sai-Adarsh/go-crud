package main

import (
	"database/sql"
	"net/http"
	"strconv"

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

func CreateJoke(c *gin.Context) {
	// initialize the database
	c.Header("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// fetch the existing data
	results, err := db.Query("SELECT * FROM `jokes` ORDER BY `ID` ASC")
	if err != nil {
		panic(err.Error())
	}

	// local variable to store the existing data
	var jokesTemp = []Joke{}

	// put the existing data into the local variable
	for results.Next() {
		var eachJoke Joke
		err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
		if err != nil {
			panic(err.Error())
		}
		jokesTemp = append(jokesTemp, eachJoke)
	}

	if newJoke := c.Param("newJoke"); newJoke != "" {
		// create a new joke
		query := "INSERT INTO `jokes` (`ID`, `Likes`, `Joke`) VALUES ('" + strconv.Itoa(len(jokesTemp)+1) + "', '0', '" + newJoke + "');"
		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		results, err := db.Query("SELECT * FROM `jokes` ORDER BY `ID` ASC")
		if err != nil {
			panic(err.Error())
		}

		// local variable to store the existing data
		var jokesTemp = []Joke{}

		// put the existing data into the local variable
		for results.Next() {
			var eachJoke Joke
			err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
			if err != nil {
				panic(err.Error())
			}
			jokesTemp = append(jokesTemp, eachJoke)
		}
		c.JSON(http.StatusOK, jokesTemp)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func JokeHandler(c *gin.Context) {
	// initialize the database
	c.Header("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// fetch the existing data
	results, err := db.Query("SELECT * FROM `jokes`")
	if err != nil {
		panic(err.Error())
	}

	// local variable to store the existing data
	var jokesTemp = []Joke{}

	// put the existing data into the local variable
	for results.Next() {
		var eachJoke Joke
		err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
		if err != nil {
			panic(err.Error())
		}
		jokesTemp = append(jokesTemp, eachJoke)
	}
	c.JSON(http.StatusOK, jokesTemp)
}

func LikeJoke(c *gin.Context) {
	// initialize the database
	c.Header("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// fetch the existing data
	results, err := db.Query("SELECT * FROM `jokes`")
	if err != nil {
		panic(err.Error())
	}

	// local variable to store the existing data
	var jokesTemp = []Joke{}

	// put the existing data into the local variable
	for results.Next() {
		var eachJoke Joke
		err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
		if err != nil {
			panic(err.Error())
		}
		jokesTemp = append(jokesTemp, eachJoke)
	}

	// fetch the jokeID from the URL
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokesTemp); i++ {
			// if the jokeID matches the jokeID in the database
			if jokesTemp[i].ID == jokeID {
				// increment the likes by 1
				jokesTemp[i].Likes += 1
				query := "UPDATE `jokes` SET `Likes`=" + strconv.Itoa(jokesTemp[i].Likes) + " WHERE `ID`=" + strconv.Itoa(jokeID)
				insert, err := db.Query(query)
				if err != nil {
					panic(err.Error())
				}
				defer insert.Close()
			}
		}
		// return the updated data
		c.JSON(http.StatusOK, &jokesTemp)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateJoke(c *gin.Context) {
	// initialize the database
	c.Header("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// fetch the existing data
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		// fetch the new joke from the URL
		if newJoke := c.Param("newJoke"); newJoke != "" {
			// update the joke in the database
			query := "UPDATE `jokes` SET `Joke`='" + newJoke + "' WHERE `ID`=" + strconv.Itoa(jokeID)
			insert, err := db.Query(query)
			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()

			// return the updated data
			results, err := db.Query("SELECT * FROM `jokes`")
			if err != nil {
				panic(err.Error())
			}

			// local variable to store the existing data
			var jokesTemp = []Joke{}

			// put the existing data into the local variable
			for results.Next() {
				var eachJoke Joke
				err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
				if err != nil {
					panic(err.Error())
				}
				jokesTemp = append(jokesTemp, eachJoke)
			}
			c.JSON(http.StatusOK, &jokesTemp)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func DeleteJoke(c *gin.Context) {
	// initialize the database
	c.Header("Content-Type", "application/json")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// fetch the existing data
	if jokeID, err := strconv.Atoi(c.Param("jokeID")); err == nil {

		// delete the joke from the database
		query := "DELETE FROM `jokes` WHERE `ID`=" + strconv.Itoa(jokeID)
		insert, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		// return the updated data
		results, err := db.Query("SELECT * FROM `jokes`")
		if err != nil {
			panic(err.Error())
		}

		// local variable to store the existing data
		var jokesTemp = []Joke{}

		// put the existing data into the local variable
		for results.Next() {
			var eachJoke Joke
			err = results.Scan(&eachJoke.ID, &eachJoke.Likes, &eachJoke.Joke)
			if err != nil {
				panic(err.Error())
			}
			jokesTemp = append(jokesTemp, eachJoke)

		}
		c.JSON(http.StatusOK, &jokesTemp)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
