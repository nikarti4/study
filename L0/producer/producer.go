package producer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"L0/common"
	"L0/model"

	stan "github.com/nats-io/stan.go"
)

func ProduceOrder(cn stan.Conn, chanName string) {

	order, err := ReadOrderFromFile("../input_data/model.json")
	common.CheckFatal(err)

	strOrder, err := json.Marshal(order)
	common.CheckFatal(err)
	fmt.Println("Send some message via NATS!")

	err = cn.Publish(chanName, []byte(strOrder))
	common.CheckFatal(err)
	//fmt.Println(string(strOrder))
	time.Sleep(2 * time.Second)
}

func ReadOrderFromFile(path string) (out model.Order_t, err error) {
	fromFile, err := ioutil.ReadFile(path)
	common.CheckFatal(err)
	err = json.Unmarshal(fromFile, &out)

	return out, err
}
