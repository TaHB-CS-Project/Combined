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

type Hospital struct {
	hospital_id      int
	hospital_city    int
	hospital_address string
	hospital_name    string
}

func main() {
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

	fmt.Println("Successfully connected!")

	//START OF CRUD (Create, Read, Update, Delete)
	//CREATE
	sqlStatement_create := `
	INSERT INTO Hospital (hospital_city, hospital_address, hospital_name)
	VALUES ($2, $3, $4)
	RETURNING id`
	id := 0
	_, err = db.Exec(sqlStatement_create, "Dallas", "Westheimer Rd", "Freedom Hospital")
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)

	//READ
	sqlStatement_read := `
	SELECT * FROM Hospital
	WHERE hospital_id = $1;`
	var hospital Hospital
	row := db.QueryRow(sqlStatement_read, 1)
	err = row.Scan(&hospital.hospital_id, &hospital.hospital_address, &hospital.hospital_name)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned.")
		return
	case nil:
		fmt.Println(hospital)
	default:
		panic(err)
	}
	_, err = db.Exec(sqlStatement_read, 1, "NewFirstCityChangeTest")
	if err != nil {
		panic(err)
	}

	//UPDATE
	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_city = $2
	WHERE hospital_id = $1;`
	_, err = db.Exec(sqlStatement_update, 1, "NewFirstCityChangeTest")
	if err != nil {
		panic(err)
	}

	//DELETE
	sqlStatement_delete := `
	DELETE FROM Hospital
	WHERE hospital_id = $1;`
	//delete the row of the id number
	res, err := db.Exec(sqlStatement_delete, 1)
	if err != nil {
		panic(err)
	}
	//print out the number of rows affected by the delete command to confirm
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
