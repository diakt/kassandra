package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Env struct {
	DB *sql.DB
}


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "my_password"
	dbname = "temp"
)

func ConnectDb() (*sql.DB, error){
	conn_db := fmt.Sprintf("host=%s port=%d user=%s dbname=%s", host, port, user, dbname)
	db, err := sql.Open("postgres", conn_db)
	if err!=nil{
		fmt.Printf("conndb fails yet again")
		return &sql.DB{}, err
	}
	return db, nil




}