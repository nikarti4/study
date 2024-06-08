package main

import (
	"fmt"
	"sync"
	"time"
)

// структура с мьютексом и суммой квадратов
type psData struct {
	mu   sync.Mutex
	data int
}

// метод структуры, получаем на вход число, возводим в квадрат и добавляем
// к данным структуры
func (d *psData) Inc(a int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.data += (a * a)
}

// получаем число и записываем его квадрат в канал
func pow2chan(a int, c chan int) {
	var out int

	out = a * a

	c <- out
}

func main() {

	// создаем слайс
	a := make([]int, 0)
	// создаем канал
	c := make(chan int)
	// добавляем 5 элементов в слайс
	a = append(a, 2, 4, 6, 8, 10)
	// переменная - сумма квадратов
	var aSum int

	aSum = 0

	// кладем квадрат каждого элемента в канал, затем забираем это значение из
	// канала и считаем сумму
	for _, elem := range a {
		go pow2chan(elem, c)
		aSum += <-c
	}

	fmt.Println(aSum)

	aSum = 0

	// вариант с мьютексами
	var s psData
	// добавляем к данным квадрат елемента
	// и делаем задержку
	for _, elem := range a {
		go s.Inc(elem)
		time.Sleep(time.Millisecond)
	}

	fmt.Println(s.data)
}