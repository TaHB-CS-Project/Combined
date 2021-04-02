package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-3-21-100-78.us-east-2.compute.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

// declare global db to use across other files
var db *sql.DB

func main() {

	//start database instance for use
	initDB()
	//http.HandleFunc("/signin", Signin)
	//http.HandleFunc("/signup", Signup)
	makehospital("Dallas", "Westheimer Rd", "Freedom Hospital")
	//sethospital_city(1, "Test City for Testing")
	//gethospital_city(1)
	//deletehospital(150)
}
func initDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	/* 	defer db.Close()

	   	err = db.Ping()
	   	if err != nil {
	   		panic(err)
	   	} */

	//fmt.Println("Successfully connected!")
}
