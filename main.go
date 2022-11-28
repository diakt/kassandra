package main

import (
	"fmt"
	// "io/ioutil"
	_ "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	routes "./routing"
	seed "./seed"
)





func main(){
	// env := New(Env)
	// var env error
	// env.DB, err = ConnectDb()



	fmt.Printf("In main\n")
	seed.Seed()
	routes.Routing()
	







	fmt.Printf("Exiting main\n")

	




}














