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

/* var globalSessions *session.SessionManager */

func main() {

	//start database instance for use
	initDB()
	// http.HandleFunc("/", Home)
	// http.HandleFunc("/login", SessionLogin)
	// http.HandleFunc("/logout", SessionLogout)
	// http.HandleFunc("/dbgettest", getstaff_list)
	// fmt.Printf("Starting server for testing HTTP POST...\n")
	// http.ListenAndServe(":8090", context.ClearHandler(http.DefaultServeMux))

	go server()
	client()
	getstafflisttest()

	// http.HandleFunc("/signin", signin)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	//http.HandleFunc("/signup", Signup)
	//makehospital("Dallas", "Westheimer Rd", "Freedom Hospital")
	//sethospital_city(1, "Test City for Testing")
	//gethospital_city(1)
	//deletehospital(150)
}

/* func initSession() {
	globalSessions = NewManager("memory", "gosessionid", 3000)
} */

//initalize connection to the DB
func initDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

}
