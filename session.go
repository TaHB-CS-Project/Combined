package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
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

func Login(response http.ResponseWriter, request *http.Request) {
	//read in the data from the login page bar
	request.ParseForm()
	username := request.Form.Get("email")
	password := request.Form.Get("password")
	fmt.Println("email:", username)
	fmt.Println("password:", password)

	user := user_entity{}

	sqlStatement := `
	SELECT password_hash 
	FROM user_entity 
	WHERE username=$1`
	result := db.QueryRow(sqlStatement, username)
	err := result.Scan(&user.Password_hash)
	if err != nil {
		fmt.Println("User/Pass was invalid")
	}

	check := password == user.Password_hash
	if check {
		sessions, _ := store.Get(request, "session")
		sessions.Values["username"] = username

		//save before the redirect
		sessions.Save(request, response)
		//if their username and password matches then redirect them to the dashboard(?) or whatever is the
		//main page of a succesful login
		http.Redirect(response, request, "/dashboard", http.StatusSeeOther)
	}
	if !check { //username and password doesn't match
		data := map[string]interface{}{
			"err": "Invalid Username or Password.",
		}
		tmp, _ := template.ParseFiles("Template/index.html")
		tmp.Execute(response, data)
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
	//session, _ := store.Get(request, "session")
	tmp, _ := template.ParseFiles("Template/dashboard.html")
	tmp.Execute(response, nil)
}
