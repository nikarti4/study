package main

import (
	"fmt"
	"time"
)

// способы остановки 
// - использовать каналы
// - использовать контекст
// https://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine

func worker(status <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Active")
			time.Sleep(500 * time.Millisecond)
		case <-status:
			fmt.Println("Stop!")
			return
		}
	}
}

func main() {

	status := make(chan bool)

	go worker(status)

	time.Sleep(3 * time.Second)
	status <- true

	fmt.Println("End of sub-program")
}