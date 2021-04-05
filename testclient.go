package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//used for testing
func client() {
	locJson, err := json.Marshal(user_entity{Username: "doctor1", Password_hash: "password"})
	req, err := http.NewRequest("POST", "http://localhost:8090", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("Input: ", string(locJson))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}
