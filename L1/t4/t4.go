package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	var workersCount int
	var wg sync.WaitGroup

	fmt.Print("Select number of workers: ")
	fmt.Scanln(&workersCount)

	// канал с данными
	taskChan := make(chan int)
	// канал для завершения горутины
	signalChan := make(chan os.Signal, 1)

	// запускаем воркеры
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go Worker(i, taskChan, &wg)
	}

	// читаем данные из канала, пока не придет сигнал на завершение
	go func() {
		for {
			random := rand.Int()

			select {

			case <-signalChan:
				fmt.Println("\nEND OF PROGRAMM")
				close(taskChan)
				fmt.Println("\nCLOSE WORKERS")
				close(signalChan)
				return

			default:
				taskChan <- random
				time.Sleep(time.Millisecond * 100)

			}
		}
	}()

	signal.Notify(signalChan, os.Interrupt)
	wg.Wait()
}

func Worker(id int, dataChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// пока в канал поступают данные
	for data := range dataChan {
		fmt.Printf("Worker %d recieve data: %v\n", id, data)
		time.Sleep(time.Millisecond * 100)
	}
}
