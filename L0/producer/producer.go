package producer

import (
	"strconv"
	"time"
	"io/ioutil"
	"encoding/json"

	"L0/model"

	stan "github.com/nats-io/stan.go"
)

func ProduceOrder() {

	order, err := ReadOrderFromFile("../input_data/model.json")

	// соединение, указываем ID кластера и ID клиента
	p, _ := stan.Connect("prod", "example")
	defer p.Close()

	// цикл посылки сообщений
	for i := 1; ; i++ {
		// имя канала и сообщения
		p.Publish("msg", []byte("Msg "+strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}

}

func ReadOrderFromFile(path string) (out []model.Order_t, err error) {
	fromFile := ioutil.ReadFile(path)

	err = json.Unmarshall(fromFile, &out)

	return out, err
}



