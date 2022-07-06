package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
