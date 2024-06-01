package main

import (
	"fmt"
	"sync"

	stan "github.com/nats-io/stan.go"
)

func main() {

	s, _ := stan.Connect("prod", "s1")
	defer s.Close()

	s.Subscribe("msg", func(m *stan.Msg) {
		fmt.Printf("Receive: %s\n", string(m.Data))
	})

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}