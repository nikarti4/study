package main

import (
	"fmt"
	"sync"
	"time"
)

// структура-счетчик
// мьютекс и значение счетчика
type Counter struct {
	mu sync.Mutex
	v  int
}

// увеличить счетчик на 1
func (c *Counter) Inc() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func main() {
	// объявляем счетчик
	var ctr Counter

	// количество проходов
	n := 5

	// цикл с горутинами
	for i := 0; i < n; i++ {
		go ctr.Inc()
	}

	//задержка, чтобы главная горутина не завершилась раньше
	time.Sleep(time.Millisecond)

	fmt.Println(ctr.v)
}
