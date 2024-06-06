package subscriber

import (
	"encoding/json"
	"fmt"
	"database/sql"

	"L0/common"
	"L0/model"
	"L0/dbpart"
	"L0/cache"

	stan "github.com/nats-io/stan.go"
)

func ConsumeAndSaveOrder(cn stan.Conn, chanName string, db *sql.DB, c *cache.My_cache) {
	
	var order model.Order_t
	
	_, err := cn.Subscribe(chanName, func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &order)
		common.CheckFatal(err)
		fmt.Println("Get something from NATS!")
		//fmt.Println(string(m.Data))
		dbpart.InsertOrder(db, order)
		c.Write(order)
	})

	common.CheckFatal(err)
}
