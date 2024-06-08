package main

import (
	"fmt"
	"sync"
	"time"
)

// рассчет и вывод квадрата числа
func pow2(a int) {
	fmt.Printf("%v^2 = %v\n", a, a*a)
}

// рассчет и вывод квадрата числа с использованием sync
func pow2sync(a int, s *sync.WaitGroup) {
	defer s.Done() // вычитает 1 из счетчика
	fmt.Printf("%v^2 = %v\n", a, a*a)
}

func main() {
	// по условию задан массив
	a := [5]int{2, 4, 6, 8, 10}

	fmt.Println("Вариант 1 - time.Sleep")

	// вызываем pow2 для каждого элемента массива
	for i := 0; i < 5; i++ {
		go pow2(a[i])
	}

	// задержка 1 мс, т.к. основной поток может завершиться раньше,
	// чем успеют выполниться горутины
	time.Sleep(time.Millisecond)

	// (*) ИЛИ можно через sync
	fmt.Println("Вариант 2 - sync.WaitGroup")
	var s sync.WaitGroup

	// вызываем pow2 для каждого элемента массива
	for i := 0; i < 5; i++ {
		s.Add(1) // добавить 1 к счетчику
		go pow2sync(a[i], &s)
	}

	s.Wait() // ждем завершения всех горутин
}
