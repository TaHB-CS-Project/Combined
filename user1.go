package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type user_entity1 struct {
	user_id               string
	medicalemployee_id    int
	email                 string
	email_confirmed       bool
	email_confirmed_token string
	username              string
	password              string
	salt                  string
	lockout               bool
	reset_password_stamp  string
	reset_password_date   string
}

type creds1 struct {
	username string `json: "email"`
	password string `json: "password"`
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	jsonFeed, err := ioutil.ReadAll(r.Body)
	fmt.Println("jsonFeed", jsonFeed)

	user := user_entity1{}
	json.Unmarshal([]byte(jsonFeed), &user)
	fmt.Println("Email", user.username, "Password", user.password)

	status := Status{RespStatus: "OK"}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
