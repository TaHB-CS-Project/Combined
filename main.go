package main

import (
	_ "github.com/lib/pq"
)

func main() {
	dbconnect()
	makehospital("Dallas", "Westheimer Rd", "Freedom Hospital")
	//sethospital_city(1, "Test City for Testing")
	//gethospital_city(1)
}
