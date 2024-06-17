package main

import (
	"fmt"
)

// объект DVI порта
type DVI struct{}

// метод DVI
func (d *DVI) ConnectDVI() {
	fmt.Println("Connected to DVI !")
}

// объект HDMI порта
type HDMI struct{}

// метод HDMI, пусть еще будет указан разъем
func (h *HDMI) ConnectHDMI(port int) {
	fmt.Printf("Connected to HDMI in port %v !\n", port)
}

// целевой интерфейс
type ConnectAdapter interface {
	Connection()
}

// адаптер для DVI
type DVIadapter struct {
	*DVI
}

// подключение к DVI
func (adapter *DVIadapter) Connection() {
	adapter.ConnectDVI()
}

// конструктор адаптера для DVI
func NewDVIAdapter(d *DVI) ConnectAdapter {
	return &DVIadapter{d}
}

// адаптер для HDMI
type HDMIadapter struct {
	*HDMI
}

// подключение к HDMI
func (adapter *HDMIadapter) Connection() {
	adapter.ConnectHDMI(1)
}

// конструктор адаптера для HDMI
func NewHDMIAdapter(h *HDMI) ConnectAdapter {
	return &HDMIadapter{h}
}

// объект USB порта
type USB struct{}

// подключение к USB
func (u *USB) Connection() {
	fmt.Println("Connected to USB !")
}

func main() {
	// создаем два адаптера
	a1 := NewDVIAdapter(&DVI{})
	a2 := NewHDMIAdapter(&HDMI{})

	a3 := &USB{}

	// вызываем метод интерфейса
	a1.Connection()
	a2.Connection()
	a3.Connection()
}
