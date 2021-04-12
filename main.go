package main

import (
	"database/sql"
	"fmt"
	"net/http"

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
	initstyle()
	http.HandleFunc("/", Index)
	// http.HandleFunc("/signin", Index)
	http.HandleFunc("/signin", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/dashboard", Dashboard)
	http.HandleFunc("/staff-list.html", Staff_list)
	http.HandleFunc("/add-record.html", Add_record)
	http.HandleFunc("/record-draft.html", Record_draft)
	http.HandleFunc("/record-list.html", Record_list)
	http.HandleFunc("/create-account_registerd.html", Create_account_registered)
	http.HandleFunc("/create-account.html", Create_account)
	http.HandleFunc("/diagnosis.html", Diagnosis)
	http.HandleFunc("/forgot-password-submit.html", Forgot_password_submit)
	http.HandleFunc("/forgot-password.html", Forgot_password)
	http.ListenAndServe(":8090", nil)
	// http.HandleFunc("/login", SessionLogin)
	// http.HandleFunc("/logout", SessionLogout)
	// http.HandleFunc("/dbgettest", getstaff_list)
	// fmt.Printf("Starting server for testing HTTP POST...\n")
	//http.ListenAndServe(":8090", context.ClearHandler(http.DefaultServeMux))

	//server()
	// client()
	// getstafflisttest()

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

func initstyle() {
	http.Handle("/css/", //final url can be anything
		http.StripPrefix("/css/",
			http.FileServer(http.Dir("css"))))

	http.Handle("/img/", //final url can be anything
		http.StripPrefix("/img/",
			http.FileServer(http.Dir("img"))))
}
