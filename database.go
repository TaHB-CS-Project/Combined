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

	Hospital_id   int    `json:hospital_id`
	Hospital_name string `json:hospital_name`

	Medicalemployee_firstname string `json:staff_name`
	Medicalemployee_lastname  string `json:staff_lastname`

	Patient_id        int       `json:patient_id`
	Patient_birthday  time.Time `json:dob`
	Patient_sex       string    `json:gender`
	Patient_weightlbs float32   `json:weight`

	Procedure_id   int    `json:procedure_id`
	Procedure_name string `json:procedure_name`

	Diagnosis_id   int    `json:diagnosis_id`
	Diagnosis_name string `json:diagnosis_name`
}

type Recordlist struct {
	Record_id                 int       `json:record_id`
	Hospital_name             string    `json:hospital_name`
	Start_datetime            time.Time `json:starttime`
	Medicalemployee_firstname string    `json:staff_name`
	Medicalemployee_lastname  string    `json:staff_lastname`
	Procedure_name            string    `json:procedure_name`
	Diagnosis_name            string    `json:diagnosis_name`
	Outcome                   string    `json:result`
}

type Diagnosis struct {
	Diagnosis_id   int    `json:diagnosis_id`
	Diagnosis_name string `json:diagnosis_name`
}

type Procedure struct {
	Procedure_id   int    `json:procedure_id`
	Procedure_name string `json:procedure_name`
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
	SELECT hospital_city 
	FROM Hospital
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

func getdiagnosis(w http.Response, r *http.Request) {
	sqlStatement_get := `
	SELECT * FROM diagnosis`
	diagnosis := Diagnosis{}
	var diagnosisarray []Diagnosis
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&diagnosis.Diagnosis_id, &diagnosis.Diagnosis_name)
		if err != nil {
			panic(err)
		}
		diagnosisarray = append(diagnosisarray, Diagnosis{diagnosis.Diagnosis_id, diagnosis.Diagnosis_name})
	}
	fmt.Printf("%v", diagnosisarray)
	file, _ := json.MarshalIndent(diagnosisarray, "", " ")
	_ = ioutil.WriteFile("js/diagnosis.json", file, 0644)
}

func getprocedure(w http.Response, r *http.Request) {
	sqlStatement_get := `
	SELECT * FROM procedure`
	procedure := Procedure{}
	var procedurearray []Procedure
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&procedure.Procedure_id, &procedure.Procedure_name)
		if err != nil {
			panic(err)
		}
		procedurearray = append(procedurearray, Procedure{procedure.Procedure_id, procedure.Procedure_name})
	}
	fmt.Printf("%v", procedurearray)
	file, _ := json.MarshalIndent(procedurearray, "", " ")
	_ = ioutil.WriteFile("js/procedure.json", file, 0644)
}

func getrecord_list(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot to getrecord_list\n")
	sqlStatement_get := `
		SELECT * FROM record 
		ORDER BY record_id DESC LIMIT 10`
	record := Record{}
	var recordarray []Recordlist
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&record.Record_id, &record.Hospital_id, &record.Medical_employee_id, &record.Patient_id, &record.Procedure_id, &record.Diagnosis_id,
			&record.Start_datetime, &record.End_datetime, &record.Special_notes, &record.Outcome)
		if err != nil {
			panic(err)
		}

		sqlStatement_get0 := `
		SELECT hospital_name
		FROM hospital 
		WHERE hospital_id = $1`
		//var hospital_name string
		error := db.QueryRow(sqlStatement_get0, record.Hospital_id).Scan(&record.Hospital_name)
		if error != nil {
			panic(error)
		}

		sqlStatement_get1 := `
		SELECT medicalemployee_firstname
		FROM medical_employee 
		WHERE medicalemployee_id = $1`
		error1 := db.QueryRow(sqlStatement_get1, record.Hospital_id).Scan(&record.Medicalemployee_firstname)
		if error1 != nil {
			panic(error1)
		}

		sqlStatement_get2 := `
		SELECT medicalemployee_lastname
		FROM medical_employee 
		WHERE medicalemployee_id = $1`
		error2 := db.QueryRow(sqlStatement_get2, record.Hospital_id).Scan(&record.Medicalemployee_lastname)
		if error2 != nil {
			panic(error2)
		}

		sqlStatement_get3 := `
		SELECT diagnosis_name
		FROM diagnosis
		WHERE diagnosis_id = $1`
		error3 := db.QueryRow(sqlStatement_get3, record.Diagnosis_id).Scan(&record.Diagnosis_name)
		if error3 != nil {
			panic(error3)
		}

		sqlStatement_get4 := `
		SELECT procedure_name
		FROM procedure
		WHERE procedure_id = $1`
		error4 := db.QueryRow(sqlStatement_get4, record.Procedure_id).Scan(&record.Procedure_name)
		if error4 != nil {
			panic(error4)
		}

		recordarray = append(recordarray, Recordlist{record.Record_id, record.Hospital_name, record.Start_datetime, record.Medicalemployee_firstname,
			record.Medicalemployee_lastname, record.Diagnosis_name, record.Procedure_name, record.Outcome})
	}
	fmt.Printf("%v", recordarray)
	file, _ := json.MarshalIndent(recordarray, "", " ")
	_ = ioutil.WriteFile("js/record-list.json", file, 0644)
}

func getstaff_list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	sqlStatement_get := `
		SELECT * FROM medical_employee 
		ORDER BY medicalemployee_id DESC LIMIT 10`
	medicalemployee := MedicalEmployee{}
	var medicalemployeearray []MedicalEmployee
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {

		err := row.Scan(&medicalemployee.Medicalemployee_id, &medicalemployee.Hospital_id, &medicalemployee.Medicalemployee_firstname, &medicalemployee.Medicalemployee_lastname,
			&medicalemployee.Medicalemployee_department, &medicalemployee.Medicalemployee_classification, &medicalemployee.Medicalemployee_supervisor)
		if err != nil {
			panic(err)
		}
		medicalemployeearray = append(medicalemployeearray, MedicalEmployee{medicalemployee.Medicalemployee_id, medicalemployee.Hospital_id,
			medicalemployee.Medicalemployee_firstname, medicalemployee.Medicalemployee_lastname, medicalemployee.Medicalemployee_department,
			medicalemployee.Medicalemployee_classification, medicalemployee.Medicalemployee_supervisor})
		//fmt.Println(medicalemployeearray)
	}
	//fmt.Printf("%v", medicalemployeearray)
	file, _ := json.MarshalIndent(medicalemployeearray, "", " ")
	_ = ioutil.WriteFile("js/staff-list.json", file, 0644)
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

	//hospital := Hospital{}
	//record := Record{}
	r.ParseForm()

	Hospital_name := r.Form.Get("hospital")
	Start_datetime := r.Form.Get("date")
	Patient_sex := r.Form.Get("gender")
	Patient_weightlbs := r.Form.Get("weight")
	Patient_birthday := r.Form.Get("dob_date")
	Diagnosis_name := r.Form.Get("diagnosis")
	Procedure_name := r.Form.Get("procedure")
	Outcome := r.Form.Get("result")
	Special_notes := r.Form.Get("special_notes")
	//for testing
	log.Printf("Received: %v\n", user)

	//the data names is the DATABASES name
	sqlStatement_create := `
	INSERT INTO patient (patient_birthday, patient_sex, patient_weightlbs)
	VALUES ($1, $2, $3)
	RETURNING patient_id`
	var patient_id int64
	//the names for the record.query is the STRUCT names
	error := db.QueryRow(sqlStatement_create, Patient_birthday, Patient_sex, Patient_weightlbs).Scan(&patient_id)
	if error != nil {
		panic(error)
	}

	sqlStatement_create3 := `
	SELECT FROM hospital 
	WHERE hospital_name = $1
	RETURNING hospital_id`
	var hospital_id int64
	error = db.QueryRow(sqlStatement_create3, Hospital_name).Scan(&hospital_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New hospital ID is: ", hospital_id)

	sqlStatement_create4 := `
	SELECT FROM diagnosis 
	WHERE diagonsis_name = $1
	RETURNING diagnosis_id`
	var diagnosis_id int64
	error = db.QueryRow(sqlStatement_create4, Diagnosis_name).Scan(&diagnosis_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New diagnosis ID is: ", diagnosis_id)

	sqlStatement_create5 := `
	SELECT FROM procedure 
	WHERE procedure_name = $1
	RETURNING procedure_id`
	var procedure_id int64
	error = db.QueryRow(sqlStatement_create5, Procedure_name).Scan(&procedure_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New procedure ID is: ", procedure_id)

	sqlStatement_create6 := `
	SELECT FROM user_entity
	WHERE username = $1
	RETURNING medicalemployee_id`

	var Medical_employee_id int64
	sessions, _ := store.Get(r, "session")

	error = db.QueryRow(sqlStatement_create6, sessions.Values["username"]).Scan(&Medical_employee_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New medical employee ID is: ", Medical_employee_id)

	//final statement to make record with all the foreign keys available
	sqlStatement_create2 := `
	INSERT INTO record (medicalemployee_id, procedure_id, hospital_id, diagnosis_id, patient_id, start_datetime, special_notes, outcome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING record_id`
	var record_id int64
	error1 := db.QueryRow(sqlStatement_create2, Medical_employee_id, procedure_id, hospital_id, diagnosis_id, patient_id, Start_datetime, Special_notes, Outcome).Scan(&record_id)
	if error1 != nil {
		panic(error1)
	}
	fmt.Println("New record ID is: ", record_id)

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
