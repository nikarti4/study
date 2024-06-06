package main

import (
	"L0/cache"
	"L0/common"
	"L0/dbpart"
	"L0/producer"
	"L0/subscriber"
	"L0/http_handlers"

	//"fmt"

	//"net/http"

	"github.com/gin-gonic/gin"
	stan "github.com/nats-io/stan.go"
)

func main() {
	db := dbpart.ConnectDB()
	dbpart.CreateAllTables(db)

	c := cache.Init()
	c.RestoreCache(db)

	ns, err := stan.Connect("prod", "u1")

	common.CheckFatal(err)

	chanName := "orders"

	subscriber.ConsumeAndSaveOrder(ns, chanName, db, c)

	producer.ProduceOrder(ns, chanName)

	router := gin.Default();
	router.GET("/orders/:id", http_handlers.GetServerOrderByID(c))

	router.Run("localhost:8080")

}
