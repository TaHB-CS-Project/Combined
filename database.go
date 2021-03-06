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
type Record_draft_id struct {
	Record_draft_id int `json:record_id`
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

	Patient_birthday  time.Time `json:dob`
	Patient_sex       string    `json:gender`
	Patient_weightlbs float32   `json:weight`
}

type Recordlist_old struct {
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
func getrecord_list(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "session")
	record := Recordlist{}
	var recordarray []Recordlist
	if sessions.Values["role"] == 0 {
		sqlStatement_get := `
		SELECT record.record_id, record.start_datetime, hospital.hospital_name, medical_employee.medicalemployee_firstname, medical_employee.medicalemployee_lastname, procedure.procedure_name, diagnosis.diagnosis_name , record.outcome, record.special_notes, patient.patient_birthday, patient.patient_sex, patient.patient_weightlbs
		FROM (((((record
		JOIN hospital ON record.hospital_id = hospital.hospital_id)
		JOIN medical_employee ON record.medicalemployee_id = medical_employee.medicalemployee_id)
		JOIN diagnosis ON record.diagnosis_id = diagnosis.diagnosis_id)
		JOIN procedure ON record.procedure_id = procedure.procedure_id)
		JOIN patient ON record.patient_id = patient.patient_id)
		`
		row, _ := db.Query(sqlStatement_get)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Start_datetime, &record.Hospital_name, &record.Medicalemployee_firstname,
				&record.Medicalemployee_lastname, &record.Procedure_name, &record.Diagnosis_name, &record.Outcome, &record.Special_notes,
				&record.Patient_birthday, &record.Patient_sex, &record.Patient_weightlbs)
			if err != nil {
				log.Fatal(err)
			}
			recordarray = append(recordarray, Recordlist{record.Record_id, record.Hospital_name, record.Start_datetime, record.Medicalemployee_firstname,
				record.Medicalemployee_lastname, record.Procedure_name, record.Diagnosis_name, record.Outcome, record.Special_notes, record.Patient_birthday,
				record.Patient_sex, record.Patient_weightlbs})
		}
	} else if sessions.Values["role"] == 1 {
		fmt.Println("Got to record list role 1")
		sqlStatement_create1 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
		var Medical_employee_id int64
		sessions, _ := store.Get(r, "session")

		error := db.QueryRow(sqlStatement_create1, sessions.Values["username"]).Scan(&Medical_employee_id)
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
		SELECT record.record_id, record.start_datetime, hospital.hospital_name, medical_employee.medicalemployee_firstname, medical_employee.medicalemployee_lastname, procedure.procedure_name, diagnosis.diagnosis_name, record.outcome, record.special_notes, patient.patient_birthday, patient.patient_sex, patient.patient_weightlbs
		FROM (((((record
		JOIN hospital ON record.hospital_id = hospital.hospital_id)
		JOIN medical_employee ON record.medicalemployee_id = medical_employee.medicalemployee_id)
		JOIN diagnosis ON record.diagnosis_id = diagnosis.diagnosis_id)
		JOIN procedure ON record.procedure_id = procedure.procedure_id)
		JOIN patient ON record.patient_id = patient.patient_id)
		WHERE hospital.hospital_id = $1
		`
		row, _ := db.Query(sqlStatement_get, hospital_id)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Start_datetime, &record.Hospital_name, &record.Medicalemployee_firstname,
				&record.Medicalemployee_lastname, &record.Procedure_name, &record.Diagnosis_name, &record.Outcome, &record.Special_notes,
				&record.Patient_birthday, &record.Patient_sex, &record.Patient_weightlbs)
			if err != nil {
				log.Fatal(err)
			}
			recordarray = append(recordarray, Recordlist{record.Record_id, record.Hospital_name, record.Start_datetime, record.Medicalemployee_firstname,
				record.Medicalemployee_lastname, record.Procedure_name, record.Diagnosis_name, record.Outcome, record.Special_notes, record.Patient_birthday,
				record.Patient_sex, record.Patient_weightlbs})
		}
	} else {
		fmt.Println("Got to record list role 2")
		sqlStatement_create1 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
		var Medical_employee_id int64
		sessions, _ := store.Get(r, "session")

		error := db.QueryRow(sqlStatement_create1, sessions.Values["username"]).Scan(&Medical_employee_id)
		if error != nil {
			panic(error)
		}

		sqlStatement_get := `
		SELECT record.record_id, record.start_datetime, hospital.hospital_name, medical_employee.medicalemployee_firstname, medical_employee.medicalemployee_lastname, procedure.procedure_name, diagnosis.diagnosis_name , record.outcome, record.special_notes, patient.patient_birthday, patient.patient_sex, patient.patient_weightlbs
		FROM (((((record
		JOIN hospital ON record.hospital_id = hospital.hospital_id)
		JOIN medical_employee ON record.medicalemployee_id = medical_employee.medicalemployee_id)
		JOIN diagnosis ON record.diagnosis_id = diagnosis.diagnosis_id)
		JOIN procedure ON record.procedure_id = procedure.procedure_id)
		JOIN patient ON record.patient_id = patient.patient_id)
		WHERE medical_employee.medicalemployee_id = $1
		`
		row, _ := db.Query(sqlStatement_get, Medical_employee_id)
		defer row.Close()
		for row.Next() {
			err := row.Scan(&record.Record_id, &record.Start_datetime, &record.Hospital_name, &record.Medicalemployee_firstname,
				&record.Medicalemployee_lastname, &record.Procedure_name, &record.Diagnosis_name, &record.Outcome, &record.Special_notes,
				&record.Patient_birthday, &record.Patient_sex, &record.Patient_weightlbs)
			if err != nil {
				log.Fatal(err)
			}
			recordarray = append(recordarray, Recordlist{record.Record_id, record.Hospital_name, record.Start_datetime, record.Medicalemployee_firstname,
				record.Medicalemployee_lastname, record.Procedure_name, record.Diagnosis_name, record.Outcome, record.Special_notes, record.Patient_birthday,
				record.Patient_sex, record.Patient_weightlbs})
		}
	}
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
	//log.Printf("Received: %v\n", user)

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
	//log.Printf("Received: %v\n", user)

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
	fmt.Println("Medical employee ID is: ", Medical_employee_id)

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
	INSERT INTO record (hospital_id, medicalemployee_id, patient_id, procedure_id, diagnosis_id, start_datetime, special_notes, outcome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING record_id`
	var record_id int64
	error1 := db.QueryRow(sqlStatement_create2, hospital_id, Medical_employee_id, patient_id, procedure_id, diagnosis_id, Start_datetime, Special_notes, Outcome).Scan(&record_id)
	if error1 != nil {
		panic(error1)
	}
	fmt.Println("New record ID is: ", record_id)

	fmt.Printf("\nSuccessfully created record\n")

	http.Redirect(w, r, "/user_dashboard.html", http.StatusSeeOther)
}

func create_record_draft(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\nGot to create_record_draft\n")
	w.Header().Set("Content-Type", "application/json")

	//hospital := Hospital{}
	//record := Record{}
	r.ParseForm()

	Hospital_name := r.Form.Get("hospital")
	Start_datetime := r.Form.Get("record_date")
	if Start_datetime == "" {
		Start_datetime = "2000-01-1T00:00:00Z"
	}
	Patient_sex := r.Form.Get("gender")
	if Patient_sex == "" {
		Patient_sex = "Empty"
	}
	Patient_weightlbs := r.Form.Get("weight")
	if Patient_weightlbs == "" {
		Patient_weightlbs = "0"
	}
	Patient_birthday := r.Form.Get("record_birthday")
	if Patient_birthday == "" {
		Patient_birthday = "2000-01-1T00:00:00Z"
	}
	Diagnosis_name := r.Form.Get("diagnosis")
	if Diagnosis_name == "" {
		Diagnosis_name = "Empty"
	}
	Procedure_name := r.Form.Get("procedure")
	if Procedure_name == "" {
		Procedure_name = "Empty"
	}
	Outcome := r.Form.Get("result")
	if Outcome == "" {
		Outcome = "Empty"
	}
	Special_notes := r.Form.Get("special_notes")
	if Special_notes == "" {
		Special_notes = "Empty"
	}
	//for testing
	//log.Printf("Received: %v\n", user)

	sqlStatement_create3 := `
	SELECT hospital_id
	FROM hospital 
	WHERE hospital_name = $1`
	var hospital_id int64
	error := db.QueryRow(sqlStatement_create3, Hospital_name).Scan(&hospital_id)
	if error != nil {
		hospital_id = 1
		fmt.Println("Do I get here hospital id = 1")
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
		Medical_employee_id = 1
	}
	fmt.Println("New medical employee ID is: ", Medical_employee_id)

	//the data names is the DATABASES name
	sqlStatement_create := `
	INSERT INTO patient_draft (hospital_id, medicalemployee_id, patient_birthday, patient_sex, patient_weightlbs)
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
		diagnosis_id = 1
	}
	fmt.Println("New diagnosis ID is: ", diagnosis_id)

	sqlStatement_create5 := `
	SELECT procedure_id
	FROM procedure 
	WHERE procedure_name = $1`
	var procedure_id int64
	error = db.QueryRow(sqlStatement_create5, Procedure_name).Scan(&procedure_id)
	if error != nil {
		procedure_id = 1
	}
	fmt.Println("New procedure ID is: ", procedure_id)

	//final statement to make record with all the foreign keys available
	sqlStatement_create2 := `
	INSERT INTO record_draft (medicalemployee_id, procedure_id, hospital_id, diagnosis_id, patient_id, start_datetime, special_notes, outcome)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING record_id`
	var record_id int64
	error1 := db.QueryRow(sqlStatement_create2, Medical_employee_id, procedure_id, hospital_id, diagnosis_id, patient_id, Start_datetime, Special_notes, Outcome).Scan(&record_id)
	if error1 != nil {
		log.Fatal(error1)
	}
	fmt.Println("New record draft ID is: ", record_id)

	fmt.Printf("\nSuccessfully created record draft\n")
	http.Redirect(w, r, "/user_dashboard.html", http.StatusSeeOther)
}

func getrecord_draft_list(w http.ResponseWriter, r *http.Request) {
	record_draft := Recordlist{}
	var recordarray []Recordlist

	fmt.Println("Got to record draft list role 2")

	sqlStatement_create1 := `
		SELECT medicalemployee_id
		FROM user_entity
		WHERE username = $1`
	var Medical_employee_id int64
	sessions, _ := store.Get(r, "session")

	error := db.QueryRow(sqlStatement_create1, sessions.Values["username"]).Scan(&Medical_employee_id)
	if error != nil {
		panic(error)
	}

	sqlStatement_get := `
	SELECT record_draft.record_id, record_draft.start_datetime, hospital.hospital_name, medical_employee.medicalemployee_firstname, medical_employee.medicalemployee_lastname, procedure.procedure_name, diagnosis.diagnosis_name , record_draft.outcome, record_draft.special_notes, patient.patient_birthday, patient.patient_sex, patient.patient_weightlbs
	FROM (((((record_draft
	JOIN hospital ON record_draft.hospital_id = hospital.hospital_id)
	JOIN medical_employee ON record_draft.medicalemployee_id = medical_employee.medicalemployee_id)
	JOIN diagnosis ON record_draft.diagnosis_id = diagnosis.diagnosis_id)
	JOIN procedure ON record_draft.procedure_id = procedure.procedure_id)
	JOIN patient ON record_draft.patient_id = patient.patient_id)
	WHERE medical_employee.medicalemployee_id = $1
		`
	row, _ := db.Query(sqlStatement_get, Medical_employee_id)
	defer row.Close()
	for row.Next() {
		err := row.Scan(&record_draft.Record_id, &record_draft.Start_datetime, &record_draft.Hospital_name, &record_draft.Medicalemployee_firstname,
			&record_draft.Medicalemployee_lastname, &record_draft.Procedure_name, &record_draft.Diagnosis_name, &record_draft.Outcome, &record_draft.Special_notes,
			&record_draft.Patient_birthday, &record_draft.Patient_sex, &record_draft.Patient_weightlbs)
		if err != nil {
			log.Fatal(err)
		}
		recordarray = append(recordarray, Recordlist{record_draft.Record_id, record_draft.Hospital_name, record_draft.Start_datetime, record_draft.Medicalemployee_firstname,
			record_draft.Medicalemployee_lastname, record_draft.Procedure_name, record_draft.Diagnosis_name, record_draft.Outcome, record_draft.Special_notes, record_draft.Patient_birthday,
			record_draft.Patient_sex, record_draft.Patient_weightlbs})
	}
	file, _ := json.MarshalIndent(recordarray, "", " ")
	_ = ioutil.WriteFile("js/record-draft-list.json", file, 0644)
}

func submit_record_draft(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// r_id := &Record_draft_id{}
	// error := json.NewDecoder(r.Body).Decode(r_id)
	// if error != nil {
	// 	log.Fatal(error)
	// 	return
	// }
	// jsn, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatal("Error reading the body", err)
	// }

	// fmt.Printf("ioutil.ReadAll Body: ", string(jsn))

	// err = json.Unmarshal(jsn, &r_id)
	// if err != nil {
	// 	log.Fatal("Decoding error: ", err)
	// }
	//hospital := Hospital{}
	//record := Record{}
	r.ParseForm()

	Record_Draft_id := r.Form.Get("record_draft_id")
	Hospital_name := r.Form.Get("hospital")
	Start_datetime := r.Form.Get("record_date")
	Patient_sex := r.Form.Get("gender")
	Patient_weightlbs := r.Form.Get("weight")
	Patient_birthday := r.Form.Get("record_birthday")
	Diagnosis_name := r.Form.Get("diagnosis")
	Procedure_name := r.Form.Get("procedure")
	Outcome := r.Form.Get("result")
	Special_notes := r.Form.Get("special_notes")

	sqlStatement_create_hospital := `
	SELECT hospital_id
	FROM hospital 
	WHERE hospital_name = $1`
	var hospital_id int64
	error := db.QueryRow(sqlStatement_create_hospital, Hospital_name).Scan(&hospital_id)
	if error != nil {
		panic(error)
	}
	//fmt.Println("New hospital ID is: ", hospital_id)

	sqlStatement_create_employee := `
	SELECT medicalemployee_id
	FROM user_entity
	WHERE username = $1`
	var Medical_employee_id int64
	sessions, _ := store.Get(r, "session")

	error = db.QueryRow(sqlStatement_create_employee, sessions.Values["username"]).Scan(&Medical_employee_id)
	if error != nil {
		panic(error)
	}
	fmt.Println("Medical employee ID is: ", Medical_employee_id)

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

	sqlStatement_delete := `
		DELETE FROM record_draft
		WHERE record_id = $1
		`
	delete_error := db.QueryRow(sqlStatement_delete, Record_Draft_id)
	if delete_error != nil {
		//log.Fatal(delete_error)
	}
	fmt.Println("Succesfully deleted draft.")

	http.Redirect(w, r, "/user_dashboard.html", http.StatusSeeOther)
}

func delete_record_draft(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// r.ParseForm()
	// Record_Draft_id := r.Form.Get("record_draft_id")

	r_id := &Record_draft_id{}
	// error := json.NewDecoder(r.Body).Decode(r_id)
	// if error != nil {
	// 	log.Fatal(error)
	// 	return
	// }
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//log.Fatal("Error reading the body", err)
	}

	fmt.Printf("Record id for deleted draft is: %v", r_id.Record_draft_id)
	fmt.Printf("\nioutil.ReadAll Body: ", string(jsn))

	err = json.Unmarshal(jsn, &r_id)
	if err != nil {
		//log.Fatal("Decoding error: ", err)
	}

	sqlStatement_delete := `
		DELETE FROM record_draft
		WHERE record_id = $1`
	error := db.QueryRow(sqlStatement_delete, r_id.Record_draft_id)
	if error != nil {
		//log.Fatal(delete_error)
	}

	http.Redirect(w, r, "/user_dashboard.html", http.StatusSeeOther)
}
