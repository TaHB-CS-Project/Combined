package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type user_entity struct {
	// user_id               string
	// medicalemployee_id    int
	// email                 string
	// email_confirmed       bool
	// email_confirmed_token string
	Username      string `json: email`
	Password_hash string `json: password`
	// salt                  string
	// lockout               bool
	// reset_password_stamp  string
	// reset_password_date   string
}

type correct struct {
	Correctcredentials bool `json: "correctcredentials"`
}

type incorrect struct {
	Incorrectcredentials bool `json: "incorrectcredentials"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	//allow for CORS
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "POST":
		user := user_entity{}

		correctcred := correct{
			Correctcredentials: true,
		}

		incorrectcred := incorrect{
			Incorrectcredentials: false,
		}

		correctcredJson, err := json.Marshal(correctcred)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		incorrectcredJson, err := json.Marshal(incorrectcred)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}

		jsn, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Error reading the body", err)
		}

		//fmt.Printf("ioutil.ReadAll Body: ", string(jsn))

		err = json.Unmarshal(jsn, &user)
		if err != nil {
			log.Fatal("Decoding error: ", err)
		}

		log.Printf("Received: %v\n", user)

		// Get the existing entry present in the database for the given username
		result := db.QueryRow("SELECT password_hash FROM user_entity where username=$1", user.Username)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// We create another instance of `Credentials` to store the credentials we get from the database
		storedCreds := &user_entity{}
		// Store the obtained password in `storedCreds`
		err = result.Scan(&storedCreds.Password_hash)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// Compare the stored passwords
		check := storedCreds.Password_hash == user.Password_hash
		if !check {
			w.Header().Set("Content-Type", "application/json")
			w.Write(incorrectcredJson)
			return
		}

		// If we reach this point, that means the users password was correct
		w.Header().Set("Content-Type", "application/json")
		w.Write(correctcredJson)
	}
}

//used for testing
func server() {
	http.HandleFunc("/", signin)
	// http.ListenAndServe(":8088", nil)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
