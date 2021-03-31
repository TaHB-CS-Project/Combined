package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-3-12-163-23.us-east-2.compute.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
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

func Signin(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Parse and decode the request body into a new `Credentials` instance
	creds := &user_entity{}
	error := json.NewDecoder(r.Body).Decode(creds)
	if error != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	//hashedPassword, error := bcrypt.GenerateFromPassword([]byte(creds.password_hash), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, error = db.Query("insert into users values ($1, $2)", creds.username, creds.password_hash); error != nil {
		// If there is any issue with inserting into the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
}
