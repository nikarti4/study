package subscriber

import (
	"encoding/json"
	"fmt"

	"L0/common"
	"L0/model"

	stan "github.com/nats-io/stan.go"
)

func ConsumeOrder(cn stan.Conn, chanName string) {

	var order model.Order_t
	//var s stan.Subscription
	//s, err := stan.Connect("prod", "s1")
	//common.CheckFatal(err)
	//defer s.Close()

	_, err := cn.Subscribe(chanName, func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &order)
		common.CheckFatal(err)
		fmt.Println("Get something from NATS!")
		fmt.Println(string(m.Data))
	})

	common.CheckFatal(err)

}
