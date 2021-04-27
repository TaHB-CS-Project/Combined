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

type Hospitallist struct {
	Hospital_name string `json:hospital_name`
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
	Record_id           int       `json:record_id`
	Medical_employee_id int       `json:medicalemployee_id`
	Start_datetime      time.Time `json:starttime`
	Special_notes       string    `json:special_notes`
	Outcome             string    `json:result`

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
	Special_notes             string    `json:special_notes`
}

type Diagnosis struct {
	Diagnosis_id   int    `json:diagnosis_id`
	Diagnosis_name string `json:diagnosis_name`
}

type Procedure struct {
	Procedure_id   int    `json:procedure_id`
	Procedure_name string `json:procedure_name`
}

func getdiagnosis(w http.ResponseWriter, r *http.Request) {
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
	//fmt.Printf("%v", diagnosisarray)
	file, _ := json.MarshalIndent(diagnosisarray, "", " ")
	_ = ioutil.WriteFile("js/diagnosis.json", file, 0644)
}

func getprocedure(w http.ResponseWriter, r *http.Request) {
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
	//fmt.Printf("%v", procedurearray)
	file, _ := json.MarshalIndent(procedurearray, "", " ")
	_ = ioutil.WriteFile("js/procedure.json", file, 0644)
}
func getrecord_list2() {
	//sessions, _ := store.Get(r, "session")
	record := Recordlist{}
	// hospital := Hospital{}
	// medicalemployee := MedicalEmployee{}
	// procedure := Procedure{}
	// diagnosis := Diagnosis{}
	var recordarray []Recordlist
	// if sessions.Values["role"] == 0 {
	sqlStatement_get := `
	SELECT record.record_id, record.start_datetime, hospital.hospital_name, medical_employee.medicalemployee_firstname, medical_employee.medicalemployee_lastname, diagnosis.diagnosis_name, procedure.procedure_name, record.outcome, record.special_notes
	FROM ((((record
	JOIN hospital ON record.hospital_id = hospital.hospital_id)
	JOIN medical_employee ON record.medicalemployee_id = medical_employee.medicalemployee_id)
	JOIN diagnosis ON record.diagnosis_id = diagnosis.diagnosis_id)
	JOIN procedure ON record.procedure_id = procedure.procedure_id)
	`
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&record.Record_id, &record.Start_datetime, &record.Hospital_name, &record.Medicalemployee_firstname,
			&record.Medicalemployee_lastname, &record.Diagnosis_name, &record.Procedure_name, &record.Outcome, &record.Special_notes)
		if err != nil {
			log.Fatal(err)
		}
		recordarray = append(recordarray, Recordlist{record.Record_id, record.Hospital_name, record.Start_datetime, record.Medicalemployee_firstname,
			record.Medicalemployee_lastname, record.Diagnosis_name, record.Procedure_name, record.Outcome, record.Special_notes})
	}
	file, _ := json.MarshalIndent(recordarray, "", " ")
	_ = ioutil.WriteFile("js/record-test.json", file, 0644)

}

func getrecord_list(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("\nGot to getrecord_list\n")
	sessions, _ := store.Get(r, "session")
	record := Record{}
	var recordarray []Recordlist

	if sessions.Values["role"] == 0 {
		fmt.Println("Got to record list role 0")
		sqlStatement_get := `
		SELECT * FROM record 
		ORDER BY record_id`
		row, _ := db.Query(sqlStatement_get)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Hospital_id, &record.Medical_employee_id, &record.Patient_id, &record.Procedure_id, &record.Diagnosis_id,
				&record.Start_datetime, &record.Special_notes, &record.Outcome)
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
			error1 := db.QueryRow(sqlStatement_get1, record.Medical_employee_id).Scan(&record.Medicalemployee_firstname)
			if error1 != nil {
				panic(error1)
			}

			sqlStatement_get2 := `
		SELECT medicalemployee_lastname
		FROM medical_employee 
		WHERE medicalemployee_id = $1`
			error2 := db.QueryRow(sqlStatement_get2, record.Medical_employee_id).Scan(&record.Medicalemployee_lastname)
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
				record.Medicalemployee_lastname, record.Diagnosis_name, record.Procedure_name, record.Outcome, record.Special_notes})
		}
	} else if sessions.Values["role"] == 1 {
		fmt.Println("Got to record list role 1")
		sqlStatement_create6 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
		var Medical_employee_id int64
		sessions, _ := store.Get(r, "session")

		error := db.QueryRow(sqlStatement_create6, sessions.Values["username"]).Scan(&Medical_employee_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_gethospital := `
		SELECT hospital_id
		FROM medical_employee
		WHERE medicalemployee_id = $1`
		var hospital_id int64
		error = db.QueryRow(sqlStatement_gethospital, Medical_employee_id).Scan(&hospital_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_get := `
		SELECT * FROM record 
		WHERE hospital_id = $1`
		row, _ := db.Query(sqlStatement_get, hospital_id)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Hospital_id, &record.Medical_employee_id, &record.Patient_id, &record.Procedure_id, &record.Diagnosis_id,
				&record.Start_datetime, &record.Special_notes, &record.Outcome)
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
			error1 := db.QueryRow(sqlStatement_get1, record.Medical_employee_id).Scan(&record.Medicalemployee_firstname)
			if error1 != nil {
				panic(error1)
			}

			sqlStatement_get2 := `
		SELECT medicalemployee_lastname
		FROM medical_employee 
		WHERE medicalemployee_id = $1`
			error2 := db.QueryRow(sqlStatement_get2, record.Medical_employee_id).Scan(&record.Medicalemployee_lastname)
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
				record.Medicalemployee_lastname, record.Diagnosis_name, record.Procedure_name, record.Outcome, record.Special_notes})
		}
	} else {
		fmt.Println("Got to record list role 2")
		sqlStatement_create6 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
		var Medical_employee_id int64
		sessions, _ := store.Get(r, "session")

		error := db.QueryRow(sqlStatement_create6, sessions.Values["username"]).Scan(&Medical_employee_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_get := `
		SELECT * FROM record 
		WHERE medicalemployee_id = $1`
		row, _ := db.Query(sqlStatement_get, Medical_employee_id)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Hospital_id, &record.Medical_employee_id, &record.Patient_id, &record.Procedure_id, &record.Diagnosis_id,
				&record.Start_datetime, &record.Special_notes, &record.Outcome)
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
			error1 := db.QueryRow(sqlStatement_get1, record.Medical_employee_id).Scan(&record.Medicalemployee_firstname)
			if error1 != nil {
				panic(error1)
			}

			sqlStatement_get2 := `
		SELECT medicalemployee_lastname
		FROM medical_employee 
		WHERE medicalemployee_id = $1`
			error2 := db.QueryRow(sqlStatement_get2, record.Medical_employee_id).Scan(&record.Medicalemployee_lastname)
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
				record.Medicalemployee_lastname, record.Diagnosis_name, record.Procedure_name, record.Outcome, record.Special_notes})
		}
	}

	//fmt.Printf("%v", recordarray)
	file, _ := json.MarshalIndent(recordarray, "", " ")
	_ = ioutil.WriteFile("js/record-list.json", file, 0644)
}

func gethospital_list(w http.ResponseWriter, r *http.Request) {
	sqlStatement_get := `
	SELECT hospital_name 
	FROM hospital`
	hospital := Hospitallist{}
	var hospitalarray []Hospitallist
	row, _ := db.Query(sqlStatement_get)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&hospital.Hospital_name)
		if err != nil {
			panic(err)
		}
		hospitalarray = append(hospitalarray, Hospitallist{hospital.Hospital_name})
	}
	//fmt.Printf("%v", hospitalarray)
	file, _ := json.MarshalIndent(hospitalarray, "", " ")
	_ = ioutil.WriteFile("js/hospitalnamelist.json", file, 0644)
}

func getstaff_list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	sessions, _ := store.Get(r, "session")
	medicalemployee := MedicalEmployee{}
	var medicalemployeearray []MedicalEmployee

	if sessions.Values["role"] == 0 {
		fmt.Println("Got to staff list role 0")
		sqlStatement_get := `
		SELECT * FROM medical_employee 
		ORDER BY medicalemployee_id`
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
			//fmt.Printf("%v", medicalemployeearray)
		}
	} else {
		fmt.Println("Got to record list role 1")
		sqlStatement_create6 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
		var Medical_employee_id int64
		sessions, _ := store.Get(r, "session")

		error := db.QueryRow(sqlStatement_create6, sessions.Values["username"]).Scan(&Medical_employee_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_gethospital := `
		SELECT hospital_id
		FROM medical_employee
		WHERE medicalemployee_id = $1`
		var hospital_id int64
		error = db.QueryRow(sqlStatement_gethospital, Medical_employee_id).Scan(&hospital_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_get := `
		SELECT * FROM medical_employee
		WHERE hospital_id = $1`
		row, _ := db.Query(sqlStatement_get, hospital_id)
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
			//fmt.Printf("%v", medicalemployeearray)
		}
	}
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
	Start_datetime := r.Form.Get("record_date")
	Patient_sex := r.Form.Get("gender")
	Patient_weightlbs := r.Form.Get("weight")
	Patient_birthday := r.Form.Get("record_birthday")
	Diagnosis_name := r.Form.Get("diagnosis")
	Procedure_name := r.Form.Get("procedure")
	Outcome := r.Form.Get("result")
	Special_notes := r.Form.Get("special_notes")
	//for testing
	log.Printf("Received: %v\n", user)

	fmt.Printf("\nStart Date: %v\n", Start_datetime)
	fmt.Printf("\nBirthday Date: %v\n", Patient_birthday)

	sqlStatement_create3 := `
	SELECT hospital_id
	FROM hospital 
	WHERE hospital_name = $1`
	var hospital_id int64
	error := db.QueryRow(sqlStatement_create3, Hospital_name).Scan(&hospital_id)
	if error != nil {
		panic(error)
	}
	//fmt.Println("New hospital ID is: ", hospital_id)

	sqlStatement_create6 := `
	SELECT medicalemployee_id
	FROM user_entity
	WHERE username = $1`
	var Medical_employee_id int64
	sessions, _ := store.Get(r, "session")

	error = db.QueryRow(sqlStatement_create6, sessions.Values["username"]).Scan(&Medical_employee_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New medical employee ID is: ", Medical_employee_id)

	//the data names is the DATABASES name
	sqlStatement_create := `
	INSERT INTO patient (hospital_id, medicalemployee_id, patient_birthday, patient_sex, patient_weightlbs)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING patient_id`
	var patient_id int64
	//the names for the record.query is the STRUCT names
	error = db.QueryRow(sqlStatement_create, hospital_id, Medical_employee_id, Patient_birthday, Patient_sex, Patient_weightlbs).Scan(&patient_id)
	if error != nil {
		panic(error)
	}

	sqlStatement_create4 := `
	SELECT diagnosis_id
	FROM diagnosis 
	WHERE diagnosis_name = $1`
	var diagnosis_id int64
	error = db.QueryRow(sqlStatement_create4, Diagnosis_name).Scan(&diagnosis_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New diagnosis ID is: ", diagnosis_id)

	sqlStatement_create5 := `
	SELECT procedure_id
	FROM procedure 
	WHERE procedure_name = $1`
	var procedure_id int64
	error = db.QueryRow(sqlStatement_create5, Procedure_name).Scan(&procedure_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("New procedure ID is: ", procedure_id)

	//final statement to make record with all the foreign keys available
	sqlStatement_create2 := `
	INSERT INTO record (medicalemployee_id, procedure_id, hospital_id, diagnosis_id, patient_id, start_datetime, special_notes, outcome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING record_id`
	var record_id int64
	error1 := db.QueryRow(sqlStatement_create2, Medical_employee_id, procedure_id, hospital_id, diagnosis_id, patient_id, Start_datetime, Special_notes, Outcome).Scan(&record_id)
	if error1 != nil {
		panic(error1)
	}
	fmt.Println("New record ID is: ", record_id)

	fmt.Printf("\nSuccessfully created record\n")

	http.Redirect(w, r, "/user_dashboard.html", http.StatusSeeOther)
}
