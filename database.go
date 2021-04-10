package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

//Need to figure out birthday (datetime) CRUD
type Hospital struct {
	Hospital_id      int    `json:hospital_id`
	Hospital_city    string `json:hospital_city`
	Hospital_address string `json:hospital_address`
	Hospital_name    string `json:hospital_name`
}

type MedicalEmployee struct {
	Medicalemployee_id             int            `json:staff_id`
	Hospital_id                    int            `json:hospital_id`
	Medicalemployee_firstname      string         `json:staff_name`
	Medicalemployee_lastname       string         `json:staff_lastname`
	Medicalemployee_department     string         `json:staff_department`
	Medicalemployee_classification string         `json:staff_role`
	Medicalemployee_supervisor     sql.NullString `json:supervisor`
}

type Record struct {
	Record_id           int            `json:record_id`
	Medical_employee_id int            `json:medicalemployee_id`
	Start_datetime      time.Time      `json:starttime`
	End_datetime        time.Time      `json:endtime`
	Special_notes       sql.NullString `json:special_notes`
	Outcome             string         `json:result`

	Hospital_id      int    `json:hospital_id`
	Hospital_city    string `json:hospital_city`
	Hospital_address string `json:hospital_address`
	Hospital_name    string `json:hospital_name`

	Patient_id        int       `json:patient_id`
	Patient_age       int       `json:age`
	Patient_birthday  time.Time `json:dob`
	Patient_sex       string    `json:gender`
	Patient_weightlbs float32   `json:weight`

	Procedure_id   int    `json:procedure_id`
	Procedure_name string `json:procedure_name`

	Diagnosis_id   int    `json:diagnosis_id`
	Diagnosis_name string `json:diagnosis_name`
}

func makehospital(city, address, name string) {
	sqlStatement_create := `
	 INSERT INTO hospital3 ( hospital_city, hospital_address, hospital_name)
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
	error := row.Scan(&hospital.Hospital_city)
	switch error {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(hospital.Hospital_city)
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

func getstaff_list(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot to getstaff_list\n")
	w.Header().Set("Content-Type", "application/json")
	sqlStatement_get := `
		SELECT * FROM medical_employee 
		ORDER BY medicalemployee_id DESC LIMIT 10`
	medicalemployee := MedicalEmployee{}
	//medicalemployeearray := []MedicalEmployee{}
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {

		err := row.Scan(&medicalemployee.Medicalemployee_id, &medicalemployee.Hospital_id, &medicalemployee.Medicalemployee_firstname, &medicalemployee.Medicalemployee_lastname,
			&medicalemployee.Medicalemployee_department, &medicalemployee.Medicalemployee_classification, &medicalemployee.Medicalemployee_supervisor)
		if err != nil {
			panic(err)
		}
		var medicalemployeearray []MedicalEmployee
		medicalemployeearray = append(medicalemployeearray, MedicalEmployee{medicalemployee.Medicalemployee_id, medicalemployee.Hospital_id,
			medicalemployee.Medicalemployee_firstname, medicalemployee.Medicalemployee_lastname, medicalemployee.Medicalemployee_department,
			medicalemployee.Medicalemployee_classification, medicalemployee.Medicalemployee_supervisor})
		medicalemployeeJson, err := json.Marshal(medicalemployeearray)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
		}
		fmt.Println(medicalemployeearray)
		w.Write(medicalemployeeJson)
	}
}

func set_record(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot to set_record\n")
	w.Header().Set("Content-Type", "application/json")
	record := Record{}

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}

	//fmt.Printf("ioutil.ReadAll Body: ", string(jsn))

	err = json.Unmarshal(jsn, &record)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	//for testing
	log.Printf("Received: %v\n", user)

	sqlStatement_set := `
	UPDATE Record
	SET record_result = $2, Special_notes = $3, outcome = $4
	WHERE record_id = $1;`
	_, error := db.Exec(sqlStatement_set, record.Record_id, record.Special_notes, record.Outcome)
	if error != nil {
		panic(error)
	}
	fmt.Printf("\nSuccessfully updated record\n")
}

func create_record(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot to create_record\n")
	w.Header().Set("Content-Type", "application/json")
	record := Record{}

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}

	//fmt.Printf("ioutil.ReadAll Body: ", string(jsn))

	err = json.Unmarshal(jsn, &record)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	//for testing
	log.Printf("Received: %v\n", user)

	sqlStatement_create := `
	INSERT INTO patient (patient_age, birthday, sex, weight_lbs)
	VALUES ($1, $2, $3)
	RETURNING patient_id`
	var patientid int64
	error := db.QueryRow(sqlStatement_create, record.Patient_age, record.Patient_birthday, record.Patient_sex, record.Patient_weightlbs).Scan(&patientid)
	if error != nil {
		panic(error)
	}

	sqlStatement_create2 := `
	INSERT INTO record (start_datetime, end_datetime, special_notes, outcome)
	VALUES ($1, $2, $3, $4)
	RETURNING record_id`
	var recordid int64
	error1 := db.QueryRow(sqlStatement_create2, record.Start_datetime, record.End_datetime, record.Special_notes).Scan(&recordid)
	if error1 != nil {
		panic(error1)
	}
	fmt.Println("New record ID is: ", recordid)

	fmt.Printf("\nSuccessfully created record\n")
}

// func gethospital_address(id int) {
// 	sqlStatement_read := `
// 	SELECT hospital_address FROM Hospital
// 	WHERE hospital_id = $1;`
// 	var hospital Hospital
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&hospital.hospital_address)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(hospital.hospital_address)
// 	default:
// 		panic(error)
// 	}
// }

// func sethospital_address(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Hospital
// 	SET hospital_address = $2
// 	WHERE hospital_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func gethospital_name(id int) {
// 	sqlStatement_read := `
// 	SELECT hospital_name FROM Hospital
// 	WHERE hospital_id = $1;`
// 	var hospital Hospital
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&hospital.hospital_name)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(hospital.hospital_name)
// 	default:
// 		panic(error)
// 	}
// }

// func sethospital_name(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Hospital
// 	SET hospital_name = $2
// 	WHERE hospital_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func makemedicalemployee(firstname, lastname, department, classification, supervisor string) {
// 	sqlStatement_create := `
// 	 INSERT INTO Medical_Employee (medicalemployee_firstname, medicalemployee_lastname, medicalemployee_department, medicalemployee_classification, medicalemployee_supervisor)
// 	 VALUES ($1, $2, $3, $4, $5)
// 	 RETURNING medicalemployee_id`

// 	var id int64
// 	err := db.QueryRow(sqlStatement_create, firstname, lastname, department, classification, supervisor).Scan(&id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New record ID is: ", id)
// }

// func deletemedicalemployee(id int) {
// 	sqlStatement_delete := `
// 	DELETE FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
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

// func getmedicalemployee_firstname(id int) {
// 	sqlStatement_read := `
// 	SELECT medicalemployee_firstname FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
// 	var medicalemployee MedicalEmployee
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&medicalemployee.medicalemployee_firstname)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(medicalemployee.medicalemployee_firstname)
// 	default:
// 		panic(error)
// 	}
// }

// func setmmedicalemployee_firstname(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Medical_Employee
// 	SET medicalemployee_firstname = $2
// 	WHERE medicalemployee_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getmedicalemployee_lastname(id int) {
// 	sqlStatement_read := `
// 	SELECT medicalemployee_lastname FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
// 	var medicalemployee MedicalEmployee
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&medicalemployee.medicalemployee_lastname)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(medicalemployee.medicalemployee_lastname)
// 	default:
// 		panic(error)
// 	}
// }

// func setmedicalemployee_lastname(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Medical_Employee
// 	SET medicalemployee_lastname = $2
// 	WHERE medicalemployee_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getmedicalemployee_department(id int) {
// 	sqlStatement_read := `
// 	SELECT medicalemployee_department FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
// 	var medicalemployee MedicalEmployee
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&medicalemployee.medicalemployee_department)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(medicalemployee.medicalemployee_department)
// 	default:
// 		panic(error)
// 	}
// }

// func setmedicalemployee_department(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Medical_Employee
// 	SET medicalemployee_department = $2
// 	WHERE medicalemployee_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getmedicalemployee_classification(id int) {
// 	sqlStatement_read := `
// 	SELECT medicalemployee_classification FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
// 	var medicalemployee MedicalEmployee
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&medicalemployee.medicalemployee_classification)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(medicalemployee.medicalemployee_classification)
// 	default:
// 		panic(error)
// 	}
// }

// func setmedicalemployee_classification(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Medical_Employee
// 	SET medicalemployee_classification = $2
// 	WHERE medicalemployee_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func getmedicalemployee_supervisor(id int) {
// 	sqlStatement_read := `
// 	SELECT medicalemployee_supervisor FROM Medical_Employee
// 	WHERE medicalemployee_id = $1;`
// 	var medicalemployee MedicalEmployee
// 	row := db.QueryRow(sqlStatement_read, id)
// 	error := row.Scan(&medicalemployee.medicalemployee_supervisor)
// 	switch error {
// 	case sql.ErrNoRows:
// 		fmt.Println("No rows were returned!")
// 	case nil:
// 		fmt.Println(medicalemployee.medicalemployee_supervisor)
// 	default:
// 		panic(error)
// 	}
// }

// func setmedicalemployee_supervisor(id int, name string) {
// 	sqlStatement_update := `
// 	UPDATE Medical_Employee
// 	SET medicalemployee_supervisor = $2
// 	WHERE medicalemployee_id = $1;`
// 	_, err := db.Exec(sqlStatement_update, id, name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func makepatient_lbs(age, sex, weightlbs string) {
// 	sqlStatement_create := `
// 	 INSERT INTO patient (patient_age, patient_sex, patient_weightlbs)
// 	 VALUES ($1, $2, $3)
// 	 RETURNING patient_id`

// 	var id int64
// 	err := db.QueryRow(sqlStatement_create, age, sex, weightlbs).Scan(&id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New record ID is: ", id)
// }

// func makepatient_kilo(age, sex, weightkilos string) {
// 	sqlStatement_create := `
// 	 INSERT INTO patient (patient_age, patient_sex, patient_weightkilos)
// 	 VALUES ($1, $2, $3)
// 	 RETURNING patient_id`

// 	var id int64
// 	err := db.QueryRow(sqlStatement_create, age, sex, weightkilos).Scan(&id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("New record ID is: ", id)
// }

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
