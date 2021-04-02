package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Need to figure out birthday (datetime) CRUD and if patient CRUD is even needed...
type Hospital struct {
	hospital_id      int
	hospital_city    string
	hospital_address string
	hospital_name    string
}

type MedicalEmployee struct {
	medicalemployee_firstname      string
	medicalemployee_lastname       string
	medicalemployee_department     string
	medicalemployee_classification string
	medicalemployee_supervisor     string
}

type Patient struct {
	patient_age               int
	patient_ageclassification string
	//patient_birthday string
	patient_sex        string
	patient_weightlbs  float32
	patient_weightkilo float32
}

type Record struct {
	//start_datetime string
	//end_datetime string
	//special_notes string
	outcome string
}

func makehospital(city, address, name string) {
	sqlStatement_create := `
	 INSERT INTO hospital ( hospital_city, hospital_address, hospital_name)
	 VALUES ($1, $2, $3)
	 RETURNING hospital_id`

	var id int64
	err := db.QueryRow(sqlStatement_create, city, address, name).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)
}

func deletehospital(id int) {
	sqlStatement_delete := `
	DELETE FROM Hospital
	WHERE hospital_id = $1;`
	res, err := db.Exec(sqlStatement_delete, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func gethospital_city(id int) {
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
	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_city = $2
	WHERE hospital_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func gethospital_address(id int) {
	sqlStatement_read := `
	SELECT hospital_address FROM Hospital
	WHERE hospital_id = $1;`
	var hospital Hospital
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&hospital.hospital_address)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(hospital.hospital_address)
	default:
		panic(error)
	}
}

func sethospital_address(id int, name string) {
	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_address = $2
	WHERE hospital_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func gethospital_name(id int) {
	sqlStatement_read := `
	SELECT hospital_name FROM Hospital
	WHERE hospital_id = $1;`
	var hospital Hospital
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&hospital.hospital_name)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(hospital.hospital_name)
	default:
		panic(error)
	}
}

func sethospital_name(id int, name string) {
	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_name = $2
	WHERE hospital_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func makemedicalemployee(firstname, lastname, department, classification, supervisor string) {
	sqlStatement_create := `
	 INSERT INTO Medical_Employee (medicalemployee_firstname, medicalemployee_lastname, medicalemployee_department, medicalemployee_classification, medicalemployee_supervisor)
	 VALUES ($1, $2, $3, $4, $5)
	 RETURNING medicalemployee_id`

	var id int64
	err := db.QueryRow(sqlStatement_create, firstname, lastname, department, classification, supervisor).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)
}

func deletemedicalemployee(id int) {
	sqlStatement_delete := `
	DELETE FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	res, err := db.Exec(sqlStatement_delete, id)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func getmedicalemployee_firstname(id int) {
	sqlStatement_read := `
	SELECT medicalemployee_firstname FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	var medicalemployee MedicalEmployee
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&medicalemployee.medicalemployee_firstname)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(medicalemployee.medicalemployee_firstname)
	default:
		panic(error)
	}
}

func setmmedicalemployee_firstname(id int, name string) {
	sqlStatement_update := `
	UPDATE Medical_Employee
	SET medicalemployee_firstname = $2
	WHERE medicalemployee_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func getmedicalemployee_lastname(id int) {
	sqlStatement_read := `
	SELECT medicalemployee_lastname FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	var medicalemployee MedicalEmployee
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&medicalemployee.medicalemployee_lastname)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(medicalemployee.medicalemployee_lastname)
	default:
		panic(error)
	}
}

func setmedicalemployee_lastname(id int, name string) {
	sqlStatement_update := `
	UPDATE Medical_Employee
	SET medicalemployee_lastname = $2
	WHERE medicalemployee_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func getmedicalemployee_department(id int) {
	sqlStatement_read := `
	SELECT medicalemployee_department FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	var medicalemployee MedicalEmployee
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&medicalemployee.medicalemployee_department)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(medicalemployee.medicalemployee_department)
	default:
		panic(error)
	}
}

func setmedicalemployee_department(id int, name string) {
	sqlStatement_update := `
	UPDATE Medical_Employee
	SET medicalemployee_department = $2
	WHERE medicalemployee_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func getmedicalemployee_classification(id int) {
	sqlStatement_read := `
	SELECT medicalemployee_classification FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	var medicalemployee MedicalEmployee
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&medicalemployee.medicalemployee_classification)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(medicalemployee.medicalemployee_classification)
	default:
		panic(error)
	}
}

func setmedicalemployee_classification(id int, name string) {
	sqlStatement_update := `
	UPDATE Medical_Employee
	SET medicalemployee_classification = $2
	WHERE medicalemployee_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func getmedicalemployee_supervisor(id int) {
	sqlStatement_read := `
	SELECT medicalemployee_supervisor FROM Medical_Employee
	WHERE medicalemployee_id = $1;`
	var medicalemployee MedicalEmployee
	row := db.QueryRow(sqlStatement_read, id)
	error := row.Scan(&medicalemployee.medicalemployee_supervisor)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(medicalemployee.medicalemployee_supervisor)
	default:
		panic(error)
	}
}

func setmedicalemployee_supervisor(id int, name string) {
	sqlStatement_update := `
	UPDATE Medical_Employee
	SET medicalemployee_supervisor = $2
	WHERE medicalemployee_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func makepatient_lbs(age, sex, weightlbs string) {
	sqlStatement_create := `
	 INSERT INTO patient (patient_age, patient_sex, patient_weightlbs)
	 VALUES ($1, $2, $3)
	 RETURNING patient_id`

	var id int64
	err := db.QueryRow(sqlStatement_create, age, sex, weightlbs).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)
}

func makepatient_kilo(age, sex, weightkilos string) {
	sqlStatement_create := `
	 INSERT INTO patient (patient_age, patient_sex, patient_weightkilos)
	 VALUES ($1, $2, $3)
	 RETURNING patient_id`

	var id int64
	err := db.QueryRow(sqlStatement_create, age, sex, weightkilos).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is: ", id)
}

//Not sure if delete patient needed
// func deletepatient(id int) {
// 	sqlStatement_delete := `
// 	DELETE FROM Patient
// 	WHERE Patient_id = $1;`
// 	res, err := db.Exec(sqlStatement_delete, id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(count)
// }
