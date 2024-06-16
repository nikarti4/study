package main

import (
	"fmt"
	"sync"
)

func putInCh(val int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- val
	fmt.Println("Put", val, "in ch", ch)
}

func main() {

	//размер массива
	n := 10

	// массив
	a := make([]int, n)
	// заполняем массив
	for i := 0; i < n; i++ {
		a[i] = i
	}

	// создаем два канала
	chX := make(chan int, n)
	chX2 := make(chan int, n)

	// для отработки горутин
	var wgX, wgX2 sync.WaitGroup

	// запись в канал из массива
	go func() {
		for _, elem := range a {
			wgX.Add(1)
			putInCh(elem, chX, &wgX)
		}

		close(chX)
	}()

	// запись из канала в канал квадратов
	go func() {
		for out := range chX {
			wgX2.Add(1)
			putInCh(out*out, chX2, &wgX2)
		}

		close(chX2)
	}()

	// вывод значений канала квадратов
	for out := range chX2 {
		fmt.Println("chX2 val = ", out)
	}

	wgX.Wait()
	wgX2.Wait()
}
