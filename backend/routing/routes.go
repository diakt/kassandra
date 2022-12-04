package routes

import (
	"net/http"

	controller "../controller"
	"github.com/gin-gonic/gin"
)

func lost(c *gin.Context){
	c.JSON(http.StatusNotFound, gin.H{
		"status": "404",
		"message": "I may have messed up here. My bad.",
	})
}

func Routing(){
	router := gin.Default()
	
	//homepage
	router.GET("/", controller.GetHomepage)

	//practices
	router.GET("/practices", controller.GetPractices)
	router.GET("/practices/:practice", controller.GetPractice)

	//workouts
	router.GET("/workouts", controller.GetWorkouts)
	router.GET("/workouts/:workout", controller.GetWorkout)

	//boats
	router.GET("/boats", controller.GetBoats)
	router.GET("/boats/:boat", controller.GetBoat)

	//athletes
	router.GET("/athletes", controller.GetAthletes)
	router.GET("/athletes/:athlete", controller.GetAthlete)
	router.POST("/athletes", controller.PostAthlete)




	//coxswains
	router.GET("/coxswains", controller.GetCoxswains)
	router.GET("/coxswains/:coxswain", controller.GetCoxswain)

	//athletes
	router.GET("/shells", controller.GetShells)
	router.GET("/shells/:shell", controller.GetShell)

	//otherwise
	router.NoRoute(lost)


	router.Run()

}