package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

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
	//if the password and username matches from the database
	if username == "doctor" && password == "password" {
		sessions, _ := store.Get(request, "session")
		sessions.Values["username"] = username

		//save before the redirect
		sessions.Save(request, response)
		//if their username and password matches then redirect them to the dashboard(?) or whatever is the
		//main page of a succesful login
		http.Redirect(response, request, "/dashboard", http.StatusSeeOther)
	} else { //username and password doesn't match
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
