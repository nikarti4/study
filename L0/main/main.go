package main

import (
	"L0/dbpart"

	//"net/http"

	//"example.com/model"

	//"github.com/gin-gonic/gin"
)



func main() {
	db := dbpart.ConnectDB()
	dbpart.CreateAllTables(db)


	// router := gin.Default();
	// router.GET("/orders/:id", GetOrdersByID)

	// router.Run("localhost:8080")

}
