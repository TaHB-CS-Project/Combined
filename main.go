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
	http.HandleFunc("/create_record", create_record) //gotta check this one
	http.HandleFunc("/user_dashboard.html", user_dashboard)
	http.HandleFunc("/user_diagnosis.html", user_diagnosis)
	http.HandleFunc("/user_procedure.html", user_procedure)
	http.HandleFunc("/user_record-draft.html", user_record_draft)
	http.HandleFunc("/user_record-list.html", user_record_list)

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
	http.HandleFunc("/create-account-second.html", Create_account_second)
	http.HandleFunc("/create_account_second", Hospitaladmin_signup)

	// http.HandleFunc("/dashboard.html", Dashboard)
	// http.HandleFunc("/procedure.html", Procedurelist)
	// http.HandleFunc("/diagnosis.html", Diagnosislist)
	// http.HandleFunc("/staff-list.html", Staff_list)
	// http.HandleFunc("/add-record.html", Add_record)
	// http.HandleFunc("/create_record", create_record)
	// http.HandleFunc("/record-draft.html", Record_draft)
	// http.HandleFunc("/record-list.html", Record_list)
	// http.HandleFunc("/create-account-second.html", Create_account_second)
	// http.HandleFunc("/create_account_second", Hospitaladmin_signup)

	http.ListenAndServe(":8090", nil)
}

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
