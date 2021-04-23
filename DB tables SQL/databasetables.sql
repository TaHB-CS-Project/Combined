CREATE TABLE Hospital (
	hospital_id SERIAL PRIMARY KEY,
	hospital_city VARCHAR(50),
	hospital_address VARCHAR(100),
	hospital_name VARCHAR(100) 
);

CREATE TABLE Medical_Employee (
	medicalemployee_id SERIAL PRIMARY KEY,
	hospital_id SERIAL ,
	medicalemployee_firstname VARCHAR(50),
	medicalemployee_lastname VARCHAR(50),
	medicalemployee_department VARCHAR(50),
	medicalemployee_classification VARCHAR(50),
	medicalemployee_supervisor VARCHAR(50),
	FOREIGN KEY(hospital_id) REFERENCES hospital(hospital_id)
);

CREATE TABLE Patient (
	patient_id SERIAL PRIMARY KEY,
	hospital_id SERIAL,
	medicalemployee_id SERIAL,
	patient_age INT,
	patient_ageclassification VARCHAR(10),
	patient_birthday TIMESTAMP,
	patient_sex VARCHAR(10),
	patient_weightlbs DECIMAL,
	FOREIGN KEY(hospital_id) REFERENCES hospital(hospital_id),
	FOREIGN KEY(medicalemployee_id) REFERENCES medical_employee(medicalemployee_id)
);

CREATE TABLE Procedure (
	procedure_id SERIAL PRIMARY KEY,
	procedure_name VARCHAR(100) 
);

CREATE TABLE Diagnosis (
	diagnosis_id SERIAL PRIMARY KEY,
	diagnosis_name VARCHAR(100) 
);

CREATE TABLE Record (
	record_id SERIAL PRIMARY KEY,
	hospital_id SERIAL,
	medicalemployee_id SERIAL,
	patient_id SERIAL,
	procedure_id SERIAL,
	diagnosis_id SERIAL,
	start_datetime TIMESTAMP,
	special_notes TEXT,
	outcome VARCHAR(20),
	FOREIGN KEY(hospital_id) REFERENCES hospital(hospital_id),
	FOREIGN KEY(medicalemployee_id) REFERENCES medical_employee(medicalemployee_id),
	FOREIGN KEY(procedure_id) REFERENCES procedure(procedure_id),
	FOREIGN KEY(diagnosis_id) REFERENCES diagnosis(diagnosis_id),
	FOREIGN KEY(patient_id) REFERENCES patient(patient_id)
);

CREATE TABLE User_Entity (
	user_id SERIAL PRIMARY KEY,
	medicalemployee_id SERIAL,
	username VARCHAR,
	password_hash VARCHAR,
	role INT,
	FOREIGN KEY(medicalemployee_id) REFERENCES medical_employee(medicalemployee_id)
);

	