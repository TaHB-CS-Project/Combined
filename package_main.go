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

//there is a datetime package that can be imported
//https://pkg.go.dev/google.golang.org/genproto/googleapis/type/datetime

type Hospital struct {
	hospital_id      int
	hospital_city    int
	hospital_address string
	hospital_name    string
}

type Medical_Employee struct {
	medicalemployee_id             int
	hospital_id                    int
	medicalemployee_firstname      string
	medicalemployee_lastname       string
	medicalemployee_department     string
	medicalemployee_classification string
	medicalemployee_supervisor     string
}

type Patient struct {
	patient_id                int
	hospital_id               int
	medicalemployee_id        int
	patient_age               int
	patient_ageclassification string
	patient_birthday          string
	patient_sex               string
	patient_weightlbs         float32
	patient_weightkilo        float32
}

type Record struct {
	record_id          int
	hospital_id        int
	medicalemployee_id int
	patient_id         int
	procedure_id       int
	symptom_id         int
	start_datetime     string
	end_datetime       string
	special_notes      string
	outcome            string
}

type Procedure struct {
	procedure_id   int
	procedure_name string
}

type Symptom struct {
	symptom_id   int
	symptom_name string
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
