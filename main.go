package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//for deployment
// var (
// 	host     = os.Getenv("PostgresDBHost")
// 	port     = os.Getenv("PostgresDBPort")
// 	user     = os.Getenv("PostgresDBUser")
// 	password = os.Getenv("PostgresDBPassword")
// 	dbname   = os.Getenv("PostgresDBUser")
// )

// testing locally on port 8090
var (
	host     = goDotEnvVariable("PostgresDBHost")
	port     = goDotEnvVariable("PostgresDBPort")
	user     = goDotEnvVariable("PostgresDBUser")
	password = goDotEnvVariable("PostgresDBPassword")
	dbname   = goDotEnvVariable("PostgresDBUser")
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load("db.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// declare global db to use across other files
var db *sql.DB

func main() {

	//start database instance for use
	initDB()
	initstyle()
	http.HandleFunc("/", Index)
	http.HandleFunc("/signin", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/create-account_registerd.html", Create_account_registerd)
	http.HandleFunc("/create-account.html", Create_account)
	http.HandleFunc("/create_account", SignUp)
	http.HandleFunc("/forgot-password-submit.html", Forgot_password_submit)
	http.HandleFunc("/forgot-password.html", Forgot_password)

	//User
	http.HandleFunc("/user_add-record.html", user_add_record)
	http.HandleFunc("/create_record", create_record)
	http.HandleFunc("/save_record", create_record_draft)
	http.HandleFunc("/user_dashboard.html", user_dashboard)
	http.HandleFunc("/user_diagnosis.html", user_diagnosis)
	http.HandleFunc("/user_procedure.html", user_procedure)
	http.HandleFunc("/user_record-draft.html", user_record_draft)
	http.HandleFunc("/user_record-list.html", user_record_list)
	http.HandleFunc("/submit_record_draft", submit_record_draft)
	http.HandleFunc("/delete_record_draft", delete_record_draft)

	//Hospital Admin
	http.HandleFunc("/hospital_admin_dashboard.html", hospital_admin_dashboard)
	http.HandleFunc("/hospital_admin_diagnosis.html", hospital_admin_diagnosis)
	http.HandleFunc("/hospital_admin_procedure.html", hospital_admin_procedure)
	http.HandleFunc("/hospital_admin_record-list.html", hospital_admin_record_list)
	http.HandleFunc("/hospital_admin_staff-list.html", hospital_admin_staff_list)

	//Admin
	http.HandleFunc("/admin_dashboard.html", admin_dashboard)
	http.HandleFunc("/admin_diagnosis.html", admin_diagnosis)
	http.HandleFunc("/admin_procedure.html", admin_procedure)
	http.HandleFunc("/admin_record-list.html", admin_record_list)
	http.HandleFunc("/admin_staff-list.html", admin_staff_list)
	http.HandleFunc("/admin_create-account-second.html", admin_create_account_second)
	http.HandleFunc("/create_account_second", Hospitaladmin_signup)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

//initalize connection to the DB
func initDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

}

//initialize frontend javascript, style, img, fonts
func initstyle() {
	http.Handle("/css/", //final url can be anything
		http.StripPrefix("/css/",
			http.FileServer(http.Dir("css"))))

	http.Handle("/img/", //final url can be anything
		http.StripPrefix("/img/",
			http.FileServer(http.Dir("img"))))

	http.Handle("/fonts/", //final url can be anything
		http.StripPrefix("/fonts/",
			http.FileServer(http.Dir("fonts"))))

	http.Handle("/js/", //final url can be anything
		http.StripPrefix("/js/",
			http.FileServer(http.Dir("js"))))

}
