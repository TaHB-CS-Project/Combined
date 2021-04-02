package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Hospital struct {
	hospital_id      int
	hospital_city    string
	hospital_address string
	hospital_name    string
}

func makehospital(city, address, name string) {
	/* 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
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
	   	} */

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

func gethospital_city(id int) {
	/* 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
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
	   	} */

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
	/* 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
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
	   	} */

	sqlStatement_update := `
	UPDATE Hospital
	SET hospital_city = $2
	WHERE hospital_id = $1;`
	_, err := db.Exec(sqlStatement_update, id, name)
	if err != nil {
		panic(err)
	}
}

func deletehospital(id int) {
	/* 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
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
	   	} */

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
