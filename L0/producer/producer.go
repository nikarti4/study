package main

import (
	"strconv"
	"time"

	stan "github.com/nats-io/stan.go"
)

func main() {

	// соединение, указываем ID кластера и ID клиента
	p, _ := stan.Connect("prod", "example")
	defer p.Close()

	// цикл посылки сообщений
	for i := 1; ; i++ {
		// имя канала и сообщения
		p.Publish("msg", []byte("Msg " + strconv.Itoa(i)))
		time.Sleep(2 * time.Second)
	}

}