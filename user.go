package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type user_entity struct {
	user_id               string
	medicalemployee_id    int
	email                 string
	email_confirmed       bool
	email_confirmed_token string
	username              string
	password_hash         string
	salt                  string
	lockout               bool
	reset_password_stamp  string
	reset_password_date   string
}

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

type login struct {
	login *bool `json: "login, omitempty"`
}

func Signup(w http.ResponseWriter, r *http.Request) {

	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Signin(w http.ResponseWriter, r *http.Request) {
	t := new(bool)
	f := new(bool)

	*t = true
	*f = false

	// Parse and decode the request body into a new `Credentials` instance
	creds := &user_entity{}
	error := json.NewDecoder(r.Body).Decode(creds)
	if error != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the existing entry present in the database for the given username
	result := db.QueryRow("SELECT password_hash FROM user_entity where username=$1", creds.username)
	if error != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		json.Marshal(login{f})
		return
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &user_entity{}
	// Store the obtained password in `storedCreds`
	error = result.Scan(&storedCreds.password_hash)
	if error != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if error == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			json.Marshal(login{f})
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		json.Marshal(login{f})
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if error = bcrypt.CompareHashAndPassword([]byte(storedCreds.password_hash), []byte(creds.password_hash)); error != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		json.Marshal(login{f})
	}
	json.Marshal(login{t})
	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}
