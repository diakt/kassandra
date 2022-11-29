package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	models "../models"
	"github.com/gin-gonic/gin"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "my_password"
	dbname = "temp"

)


func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}



func GetHomepage(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Homepage":"/",
	})

}


func GetPractices(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM practices")
	if err!=nil{
		log.Fatal(err)
	}
	
	var practices []models.Practice

	for rows.Next(){
		var practice models.Practice
		rows.Scan(&practice.Practice_ID, &practice.Practice_date, &practice.Workout_ID, &practice.Boat1, &practice.Boat2)
		practices = append(practices, practice)
	}
	rows.Close()
	db.Close()

	c.IndentedJSON(200, practices)

}

func GetPractice(c *gin.Context){
	specPractice := c.Param("practice")
	query := fmt.Sprintf("SELECT * FROM practices WHERE practice_id=%s", specPractice)
	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var practices []models.Practice

	for rows.Next(){
		var practice models.Practice
		rows.Scan(&practice.Practice_ID, &practice.Practice_date, &practice.Workout_ID, &practice.Boat1, &practice.Boat2)
		practices = append(practices, practice)
	}
	rows.Close()
	db.Close()

	c.IndentedJSON(200, practices)

}

func GetWorkouts(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM workouts")
	if err!=nil{
		log.Fatal(err)
	}
	
	var workouts []models.Workout

	for rows.Next(){
		var workout models.Workout
		rows.Scan(&workout.Workout_ID, &workout.Intensity,&workout.Volume,&workout.Structure)
		workouts = append(workouts, workout)
	}
	rows.Close()
	db.Close()

	c.IndentedJSON(200, workouts)

}

func GetWorkout(c *gin.Context){
	specWorkout := c.Param("workout")
	query := fmt.Sprintf("SELECT * FROM workouts WHERE workout_id=%s", specWorkout)
	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var workouts []models.Workout

	for rows.Next(){
		var workout models.Workout
		rows.Scan(&workout.Workout_ID, &workout.Intensity,&workout.Volume,&workout.Structure)
		workouts = append(workouts, workout)
	}
	rows.Close()
	db.Close()

	c.IndentedJSON(200, workouts)

}

func GetBoats(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM boats")
	if err!=nil{
		log.Fatal(err)
	}
	
	var boats []models.Boat

	for rows.Next(){
		var boat models.Boat
		rows.Scan(&boat.Boat_ID, &boat.Bow,&boat.Two,&boat.Three,&boat.Four,&boat.Five,&boat.Six,&boat.Seven,&boat.Eight, &boat.Cox, &boat.Shell)
		boats = append(boats, boat)
	}


	rows.Close()
	db.Close()

	c.IndentedJSON(200, boats)
	

}

func GetBoat(c *gin.Context){
	specBoat := c.Param("boat")
	query := fmt.Sprintf("SELECT * FROM boats WHERE boat_id=%s", specBoat)
	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var boats []models.Boat

	for rows.Next(){
		var boat models.Boat
		rows.Scan(&boat.Boat_ID, &boat.Bow,&boat.Two,&boat.Three,&boat.Four,&boat.Five,&boat.Six,&boat.Seven,&boat.Eight, &boat.Cox, &boat.Shell)
		boats = append(boats, boat)
	}


	rows.Close()
	db.Close()

	c.IndentedJSON(200, boats)

}


func GetAthletes(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM athletes")
	if err!=nil{
		log.Fatal(err)
	}
	
	var athletes []models.Athlete

	for rows.Next(){
		var athlete models.Athlete
		rows.Scan(&athlete.Athlete_ID, &athlete.Ath_name, &athlete.Ath_2k_watt, &athlete.Ath_6k_watt)
		athletes = append(athletes, athlete)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, athletes)

}

func GetAthlete(c *gin.Context){
	specAthlete := c.Param("athlete")
	query := fmt.Sprintf("SELECT * FROM athletes WHERE ath_id=%s", specAthlete)

	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var athletes []models.Athlete

	for rows.Next(){
		var athlete models.Athlete
		rows.Scan(&athlete.Athlete_ID, &athlete.Ath_name, &athlete.Ath_2k_watt, &athlete.Ath_6k_watt)
		athletes = append(athletes, athlete)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, athletes)

}


func GetCoxswains(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM coxswains")
	if err!=nil{
		log.Fatal(err)
	}
	
	var coxswains []models.Coxswain

	for rows.Next(){
		var coxswain models.Coxswain
		rows.Scan(&coxswain.Cox_ID, &coxswain.Cox_name, &coxswain.Cox_weight)
		coxswains = append(coxswains, coxswain)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, coxswains)

}

func GetCoxswain(c *gin.Context){
	specAthlete := c.Param("coxswain")
	query := fmt.Sprintf("SELECT * FROM coxswains WHERE cox_id=%s", specAthlete)

	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var coxswains []models.Coxswain

	for rows.Next(){
		var coxswain models.Coxswain
		rows.Scan(&coxswain.Cox_ID, &coxswain.Cox_name, &coxswain.Cox_weight)
		coxswains = append(coxswains, coxswain)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, coxswains)

}


func GetShells(c *gin.Context){
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM shells")
	if err!=nil{
		log.Fatal(err)
	}
	
	var shells []models.Shell

	for rows.Next(){
		var shell models.Shell
		rows.Scan(&shell.Shell_ID, &shell.Shell_name) 
		shells = append(shells, shell)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, shells)

}

func GetShell(c *gin.Context){
	specShell := c.Param("shell")
	query := fmt.Sprintf("SELECT * FROM shells WHERE shell_id=%s", specShell)

	db := OpenConnection()
	rows, err := db.Query(query)
	if err!=nil{
		log.Fatal(err)
	}
	
	var shells []models.Shell

	for rows.Next(){
		var shell models.Shell
		rows.Scan(&shell.Shell_ID, &shell.Shell_name) 
		shells = append(shells, shell)
	}
	
	rows.Close()
	db.Close()

    c.IndentedJSON(200, shells)

}