package main

import (
	"fmt"
	"time"
)

// константы
const (
	workMS  = 250 // время "обработки данных" для воркера
	whileMS = 200 // время одного шага посылки данных в канал
)

// воркер, принимает номер воркера и данные из канала
func worker(id int, jobs <-chan int) {
	for j := range jobs {
		fmt.Println("worker ", id, " start data ", j)
		time.Sleep(workMS * time.Millisecond)
		fmt.Println("worker ", id, " finish data ", j)
	}
}

func main() {

	var N, data int

	// ввод количества воркеров
	fmt.Print("Select number of workers: ")
	fmt.Scanln(&N)

	// канал с данными
	ch := make(chan int)

	// бесконечный цикл
	for {
		// проходимся по всем воркерам и даем им работы
		for id := 0; id < N; id++ {
			go worker(id, ch)

		}
		// посылаем новые данные в канал
		data++
		ch <- data
		// вывод текущего шага
		fmt.Println(data)
		// задержка для наглядности
		time.Sleep(whileMS * time.Millisecond)
	}
}
