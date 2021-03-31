package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "ec2-3-12-163-23.us-east-2.compute.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func dbconnect() {
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

	//fmt.Println("Successfully connected!")
}

type Hospital struct {
	hospital_id      int
	hospital_city    string
	hospital_address string
	hospital_name    string
}

func makehospital(city, address, name string) {
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

	sqlStatement_create := `
	INSERT INTO Hospital (hospital_id, hospital_city, hospital_address, hospital_name)
	VALUES (default, $1, $2, $3)
	RETURNING hospital_id`
	hospital_id := 0
	err = db.QueryRow(sqlStatement_create, city, address, name).Scan(&hospital_id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", hospital_id)
}

func gethospital_city(id int) {
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

	sqlStatement_read := `
	SELECT hospital_city FROM Hospital
	WHERE hospital_id = $1;`
	var hospital Hospital
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&hospital.hospital_city)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(hospital.hospital_city)
	default:
		panic(error)
	}
}

func sethospital_city(id int, name string) {
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

	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_city = $2
	WHERE hospital_id = $1;`
	_, err = db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}
