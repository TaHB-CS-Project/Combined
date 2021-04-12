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

func server() {
	http.HandleFunc("/", getstaff_list)
	// http.ListenAndServe(":8088", nil)
	//http.HandleFunc("/agg", newsAggHandler)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

//used for testing
func client() {
	usernameinput := "doctor"
	passwordinput := "password1"
	locJson, err := json.Marshal(user_entity{Username: usernameinput, Password_hash: passwordinput})
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

func getstafflisttest() {
	fmt.Printf("\nDo we get to getstafflisttest\n")
	resp, err := http.Get("http://localhost:8090/staff_list")
	resp.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func dbsetrecord() {
	fmt.Printf("\nDo we get to dbsetrecord\n")
}
