package main

import (
	"fmt"
	"sync"
	"time"
)

// структура для конкурентой записи в файл,
// мьютекс и данные (мапа)
type MyMap struct {
	mu   sync.Mutex
	data map[string]int
}

// увеличиваем значение ключа на 1
func (m *MyMap) Inc(key string) {
	// "лочим", чтобы только 1 горутина имела доступ в тек.мом.вр.
	m.mu.Lock()
	m.data[key]++
	m.mu.Unlock()
}

// получаем значение ключа
func (m *MyMap) Value(key string) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[key]
}

func main() {
	// создаем объект MyMap
	a := MyMap{data: make(map[string]int)}
	// запускаем горутины
	for i := 0; i < 100; i++ {
		go a.Inc("a")
	}

	time.Sleep(time.Second)
	// выводим значение
	fmt.Println(a.Value("a"))
}
