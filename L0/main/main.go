package main

import (
	"L0/common"
	"L0/dbpart"
	"L0/producer"
	"L0/subscriber"

	//"fmt"

	//"net/http"

	//"example.com/model"

	//"github.com/gin-gonic/gin"
	stan "github.com/nats-io/stan.go"
)

func main() {
	db := dbpart.ConnectDB()
	dbpart.CreateAllTables(db)

	ns, err := stan.Connect("prod", "u1")

	//fmt.Printf("%T", ns)

	common.CheckFatal(err)

	chanName := "orders"

	subscriber.ConsumeOrder(ns, chanName)

	producer.ProduceOrder(ns, chanName)

	// router := gin.Default();
	// router.GET("/orders/:id", GetOrdersByID)

	// router.Run("localhost:8080")

}
