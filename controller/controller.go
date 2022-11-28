package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)



func GetHomepage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Homepage":"/",
	})

}


func GetPractices(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Practices":"/practices",
	})

}

func GetPractice(c *gin.Context){
	specPractice := c.Param("practice")
	c.JSON(http.StatusOK, gin.H{
		"Specpractice":specPractice,
	})

}

func GetWorkouts(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Workouts":"/workouts",
	})

}

func GetWorkout(c *gin.Context){
	specWorkout := c.Param("workout")
	c.JSON(http.StatusOK, gin.H{
		"Specworkout":specWorkout,
	})

}

func GetBoats(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Boats":"/boats",
	})

}

func GetBoat(c *gin.Context){
	specBoat := c.Param("boat")
	c.JSON(http.StatusOK, gin.H{
		"Specboat":specBoat,
	})

}

func GetAthletes(c *gin.Context){
	connreq := "host=localhost port=5432 user=postgres dbname=temp sslmode=disable"
	db, err:= sql.Open("postgres",connreq)
	if err!=nil{
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM athletes;")
    if err != nil {
        panic(err)
    }
	fmt.Println(rows)
    var result []interface{}

    cols, _ := rows.Columns()
    pointers := make([]interface{}, len(cols))
    container := make([]json.RawMessage, len(cols))
    for i, _ := range pointers {
        pointers[i] = &container[i]
    }
    for rows.Next() {
        rows.Scan(pointers...)
        result = append(result, container)
    }

    fmt.Println(container)

    c.JSON(200, container)





	// c.JSON(http.StatusOK, gin.H{
	// 	"Athletes":"/athletes",
	// })

}

func GetAthlete(c *gin.Context){
	specAthlete := c.Param("athlete")
	c.JSON(http.StatusOK, gin.H{
		"Specboat":specAthlete,
	})

}

func Dingo(c *gin.Context){

}












