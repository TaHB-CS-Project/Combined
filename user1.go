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
	user_id               string
	medicalemployee_id    int
	email                 string
	email_confirmed       bool
	email_confirmed_token string
	username              string `json: email`
	password_hash         string `json: password`
	salt                  string
	lockout               bool
	reset_password_stamp  string
	reset_password_date   string
}

type correct struct {
	correctcredentials bool `json: "correctcredentials"`
}

type incorrect struct {
	incorrectcredentials bool `json: "incorrectcredentials"`
}

func signin(w http.ResponseWriter, r *http.Request) {
	user := user_entity{}

	correctcred := correct{
		correctcredentials: true,
	}

	incorrectcred := incorrect{
		incorrectcredentials: false,
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

	err = json.Unmarshal(jsn, &user)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	log.Printf("Received: %v\n", user)

	result := db.QueryRow("SELECT password_hash FROM user_entity where username=$1", user.username)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(incorrectcredJson)
	}

	storedCreds := &user_entity{}

	err = result.Scan(&storedCreds.password_hash)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(incorrectcredJson)
	}

	check := storedCreds.username == user.username
	if check != true {
		w.Header().Set("Content-Type", "application/json")
		w.Write(incorrectcredJson)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(correctcredJson)
}
