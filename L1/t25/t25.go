package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	select {
	// посылает значение в канал после t
	case <-time.After(t):
		fmt.Println("Time out")
	}
}

func main() {
	fmt.Println("Waiting...")
	sleep(5 * time.Second)
}
