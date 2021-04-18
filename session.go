package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type user_entity struct {
	// user_id               string
	// medicalemployee_id    int
	// email                 string
	// email_confirmed       bool
	// email_confirmed_token string
	Username      string `json:email`
	Password_hash string `json:password`
	// salt                  string
	// lockout               bool
	// reset_password_stamp  string
	// reset_password_date   string
}

var store = sessions.NewCookieStore([]byte("session"))

func Index(response http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("Template/index.html")
	tmp.Execute(response, nil)

}
func SignUp(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("email")
	password := request.Form.Get("password")
	repassword := request.Form.Get("psw-repeat")
	firstname := request.Form.Get("fname")
	lastname := request.Form.Get("lname")
	classification := request.Form.Get("classification")
	hospital := request.Form.Get("hospital")
	department := request.Form.Get("department")
	supervisor := request.Form.Get("supervisor")

	checkrepassword := password == repassword
	if !checkrepassword {
		fmt.Printf("Passwords do not match")
	}

	passwordhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	sqlStatementHospital := `
	SELECT hospital_id
	FROM hospital
	WHERE hospital_name = $1
	`
	var hospitalstruct Hospital
	rows := db.QueryRow(sqlStatementHospital, hospital)
	error := rows.Scan(&hospitalstruct.Hospital_id)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(hospitalstruct.Hospital_id)
	default:
		panic(error)
	}

	sqlStatementEmployee := `
	INSERT INTO medical_employee (hospital_id, medicalemployee_firstname, medicalemployee_lastname, medicalemployee_department, medicalemployee_classification, medicalemployee_supervisor)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING medicalemployee_id
	`
	var medicalemployee_id int64
	error = db.QueryRow(sqlStatementEmployee, hospitalstruct.Hospital_id, firstname, lastname, department, classification, supervisor).Scan(&medicalemployee_id)
	if error != nil {
		panic(error)
	}

	sqlStatementUser := `
	INSERT INTO user_entity (medicalemployee_id, username, password_hash, role)
	VALUES ($1, $2, $3, $4)
	`
	//	var user_id int64
	_, error = db.Exec(sqlStatementUser, medicalemployee_id, username, passwordhash, 2)
	//	error = db.QueryRow(sqlStatementUser,medicalemployee_id, username, passwordhash).Scan(&user_id)
	if error != nil {
		panic(error)
	}

	fmt.Println("Created Account")
	fmt.Println("Created email:", username)
	fmt.Println("Created password:", password)
	http.Redirect(response, request, "/index.html", http.StatusSeeOther)

}

func Login(response http.ResponseWriter, request *http.Request) {
	//read in the data from the login page bar
	request.ParseForm()
	username := request.Form.Get("email")
	password := request.Form.Get("password")
	fmt.Println("email:", username)
	fmt.Println("password:", password)

	user := user_entity{}
	//create a bcrypt hash to compare with the database stored hash
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	// TODO: Properly handle error
	// 	panic(err)
	// }

	sqlStatement := `
	SELECT password_hash 
	FROM user_entity 
	WHERE username=$1`
	result := db.QueryRow(sqlStatement, username)
	err := result.Scan(&user.Password_hash)
	if err != nil {
		fmt.Println("User/Pass was invalid")
	}
	//compare the two hashes
	if check := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password)); check != nil {
		fmt.Println("Invalid Username or Password.")
		// fmt.Println("Password Hash", []byte(user.Password_hash))
		// fmt.Println("Hash", []byte(user.Password_hash))

		tmp, _ := template.ParseFiles("Template/index.html")
		tmp.Execute(response, nil)
	} else {
		sessions, _ := store.Get(request, "session")
		sessions.Values["username"] = username

		//save before the redirect
		sessions.Save(request, response)
		//if their username and password matches then redirect them to the dashboard(?) or whatever is the
		//main page of a succesful login
		http.Redirect(response, request, "/dashboard.html", http.StatusSeeOther)
	}
}

func Logout(response http.ResponseWriter, request *http.Request) {
	//get the current session
	sessions, _ := store.Get(request, "session")
	//set the sessions time
	sessions.Options.MaxAge = -1
	//save the new session
	sessions.Save(request, response)
	//redirect the session
	http.Redirect(response, request, "/signin", http.StatusSeeOther)
}

func Dashboard(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/dashboard.html")
	tmp.Execute(response, nil)
}

func Staff_list(response http.ResponseWriter, request *http.Request) {
	getstaff_list(response, request)
	tmp, _ := template.ParseFiles("Template/staff-list.html")
	tmp.Execute(response, nil)
}

func Add_record(response http.ResponseWriter, request *http.Request) {
	gethospital_list(response, request)
	getdiagnosis(response, request)
	getprocedure(response, request)
	tmp, _ := template.ParseFiles("Template/add-record.html")
	tmp.Execute(response, nil)
}

func Create_account_registerd(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/create-account_registerd.html")
	tmp.Execute(response, nil)
}

func Create_account(response http.ResponseWriter, request *http.Request) {
	gethospital_list(response, request)
	tmp, _ := template.ParseFiles("Template/create-account.html")
	tmp.Execute(response, nil)
}

func Diagnosislist(response http.ResponseWriter, request *http.Request) {
	getdiagnosis(response, request)
	tmp, _ := template.ParseFiles("Template/diagnosis.html")
	tmp.Execute(response, nil)
}

func Forgot_password_submit(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/forgot-password-submit.html")
	tmp.Execute(response, nil)
}

func Forgot_password(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/forgot-password.html")
	tmp.Execute(response, nil)
}

func Procedurelist(response http.ResponseWriter, request *http.Request) {
	getprocedure(response, request)
	tmp, _ := template.ParseFiles("Template/procedure.html")
	tmp.Execute(response, nil)
}

func Record_draft(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/record-draft.html")
	tmp.Execute(response, nil)
}

func Record_list(response http.ResponseWriter, request *http.Request) {
	getrecord_list(response, request)
	tmp, _ := template.ParseFiles("Template/record-list.html")
	tmp.Execute(response, nil)
}

func Staff_test(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/staff-test.html")
	tmp.Execute(response, nil)
}
