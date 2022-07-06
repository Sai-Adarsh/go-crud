package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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
