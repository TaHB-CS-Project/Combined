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
	Username      string `json:email`
	Password_hash string `json:password`
	Role          int
}

var store = sessions.NewCookieStore([]byte("session"))

func SignUp(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.Form.Get("email")
	password := request.Form.Get("password")
	repassword := request.Form.Get("psw-repeat")
	firstname := request.Form.Get("fname")
	lastname := request.Form.Get("lname")
	//classification := request.Form.Get("classification")
	hospital := request.Form.Get("hospital")
	department := request.Form.Get("department")
	//supervisor := request.Form.Get("supervisor")

	checkrepassword := password == repassword
	if !checkrepassword {
		fmt.Printf("Passwords do not match")
	}
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")

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
	INSERT INTO medical_employee (hospital_id, medicalemployee_firstname, medicalemployee_lastname, medicalemployee_department, medicalemployee_classification)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING medicalemployee_id
	`
	var medicalemployee_id int64
	error = db.QueryRow(sqlStatementEmployee, hospitalstruct.Hospital_id, firstname, lastname, department, "Admin").Scan(&medicalemployee_id)
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

func Hospitaladmin_signup(response http.ResponseWriter, request *http.Request) {
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
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")

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
	_, error = db.Exec(sqlStatementUser, medicalemployee_id, username, passwordhash, 1)
	//	error = db.QueryRow(sqlStatementUser,medicalemployee_id, username, passwordhash).Scan(&user_id)
	if error != nil {
		panic(error)
	}

	fmt.Println("Created Hospital Admin Account")
	fmt.Println("Created email:", username)
	fmt.Println("Created password:", password)
	http.Redirect(response, request, "/admin_dashboard.html", http.StatusSeeOther)

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
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
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
		//main page of a successful login
		sqlStatement := `
		SELECT role
		FROM user_entity
		WHERE username = $1`
		error := db.QueryRow(sqlStatement, username).Scan(&user.Role)
		if error != nil {
			fmt.Println("Role not found")
		}

		sessions.Values["role"] = user.Role
		sessions.Save(request, response)

		if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
			fmt.Println("Got to dashboard with role 0")
			http.Redirect(response, request, "/admin_dashboard.html", http.StatusSeeOther)
		} else if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
			fmt.Println("Got to dashboard with role 1")
			http.Redirect(response, request, "/hospital_admin_dashboard.html", http.StatusSeeOther)
		} else if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
			fmt.Println("Got to dashboard with role 2")
			http.Redirect(response, request, "/user_dashboard.html", http.StatusSeeOther)
		}
	}
}

func Logout(response http.ResponseWriter, request *http.Request) {
	//get the current session
	sessions, _ := store.Get(request, "session")

	// //set the sessions time
	sessions.Options.MaxAge = -1
	sessions.Values["username"] = ""
	// //save the new session
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	sessions.Save(request, response)
	//redirect the session
	/*
		response.Header().Set("Cache-Control", "no-cache, private, max-age=0")
		response.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		response.Header().Set("Pragma", "no-cache")
		response.Header().Set("X-Accel-Expires", "0") */

	// Proxies.
	http.Redirect(response, request, "/signin", http.StatusSeeOther)
}

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/index.html")
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

func Forgot_password_submit(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/forgot-password-submit.html")
	tmp.Execute(response, nil)
}

func Forgot_password(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/forgot-password.html")
	tmp.Execute(response, nil)
}

//User Pages
///////////////////////////////////////////////////////////////////////////////////////

func user_add_record(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		tmp, _ := template.ParseFiles("Template/user_add-record.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func user_dashboard(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		tmp, _ := template.ParseFiles("Template/user_dashboard.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func user_diagnosis(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		getdiagnosis(response, request)
		tmp, _ := template.ParseFiles("Template/user_diagnosis.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func user_procedure(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		getprocedure(response, request)
		tmp, _ := template.ParseFiles("Template/user_procedure.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func user_record_draft(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		tmp, _ := template.ParseFiles("Template/user_record-draft.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func user_record_list(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 2 && sessions.Values["username"] != "" {
		getrecord_list(response, request)
		tmp, _ := template.ParseFiles("Template/user_record-list.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

//Hospital Admin Pages
///////////////////////////////////////////////////////////////////////////////////////

func hospital_admin_dashboard(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
		tmp, _ := template.ParseFiles("Template/hospital_admin_dashboard.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func hospital_admin_diagnosis(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
		getdiagnosis(response, request)
		tmp, _ := template.ParseFiles("Template/hospital_admin_diagnosis.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func hospital_admin_procedure(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
		getprocedure(response, request)
		tmp, _ := template.ParseFiles("Template/hospital_admin_procedure.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func hospital_admin_record_list(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
		getrecord_list(response, request)
		tmp, _ := template.ParseFiles("Template/hospital_admin_record-list.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func hospital_admin_staff_list(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 1 && sessions.Values["username"] != "" {
		getstaff_list(response, request)
		tmp, _ := template.ParseFiles("Template/hospital_admin_staff-list.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

//Admin Pages
///////////////////////////////////////////////////////////////////////////////////////

func admin_create_account_second(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles("Template/admin_create-account-second.html")
	tmp.Execute(response, nil)
}

func admin_dashboard(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
		tmp, _ := template.ParseFiles("Template/admin_dashboard.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func admin_diagnosis(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
		getdiagnosis(response, request)
		tmp, _ := template.ParseFiles("Template/admin_diagnosis.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func admin_procedure(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
		getprocedure(response, request)
		tmp, _ := template.ParseFiles("Template/admin_procedure.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func admin_record_list(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
		getrecord_list(response, request)
		tmp, _ := template.ParseFiles("Template/admin_record-list.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}

func admin_staff_list(response http.ResponseWriter, request *http.Request) {
	sessions, _ := store.Get(request, "session")
	response.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	response.Header().Set("Expires", "0")
	if sessions.Values["role"] == 0 && sessions.Values["username"] != "" {
		getstaff_list(response, request)
		tmp, _ := template.ParseFiles("Template/admin_staff-list.html")
		tmp.Execute(response, nil)
	} else {
		http.Redirect(response, request, "/signin", http.StatusSeeOther)
	}
}
