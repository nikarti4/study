package main

import (
	"example.com/dbpart"
)

func main() {
	db := dbpart.ConnectDB()
	dbpart.CreateAllTables(db)
}
