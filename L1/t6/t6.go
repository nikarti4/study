package main

import (
	"context"
	"fmt"
	"time"
)

// способы остановки
// - использовать каналы
// - использовать контекст

// воркер, "слушает" канал на остановку
func worker_channel(status <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Active")
			time.Sleep(500 * time.Millisecond)
		case <-status:
			return
		}
	}
}

func main() {
	// создаем канал и запускаем горутину
	status := make(chan bool)
	go worker_channel(status)

	// ждем 3 секунды и посылаем в канал сигнал
	time.Sleep(3 * time.Second)
	status <- true
	// закрываем канал
	close(status)

	fmt.Println("End of sub-program 1")

	// создаем новый контекст из переданного родительского
	// (передаем пустой), возвращается контекст и функция отмены
	ctx2, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			// если был вызван cancel()
			case <-ctx.Done():
				return
			default:
				fmt.Println("Active")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx2)

	time.Sleep(3 * time.Second)
	cancel()

	fmt.Println("End of sub-program 2")

}
