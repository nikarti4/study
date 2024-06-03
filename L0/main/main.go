package main

import (
	"example.com/dbpart"
)

func main() {
	db := dbpart.ConnectDB()
	//defer db.Close()
	//dbpart.CreateTestsTabel(db)
	dbpart.CreateDeliveryTable(db)
	dbpart.CreatePaymentTable(db)
}
