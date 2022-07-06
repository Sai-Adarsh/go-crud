package likejoke

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
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
