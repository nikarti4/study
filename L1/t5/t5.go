package main

import (
	"fmt"
	"time"
)


func main() {

	// переменная для ввода количества секунд
	var t_str string

	// создаем канал с размером буффера в 1
	ch := make(chan time.Duration, 1)
	
	// ввод секунд
	fmt.Print("Select number of seconds: ")
	fmt.Scanln(&t_str)
	// добавляем секунды в конце строки
	t_str += "s"

	// задаем заданное количество секунд типом time.Duration
	t_k, _ := time.ParseDuration(t_str)

	// задаем начальную точку отсчета времени
	start := time.Now()
	elapsed := time.Since(start)

	// цикл до заданного времени
	for t_k > elapsed {
		// получаем время с момента отсчета
		elapsed = time.Since(start)
		// посылка в канал
		ch <- elapsed
		// чтение из канала
		temp_ch := <- ch

		fmt.Println("temp_ch data:", temp_ch)
	}

	// закрываем канал
	close(ch)
}