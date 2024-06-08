package main

import (
	"fmt"
)

// структура - "базовая"
type Human struct {
	Name string
	Age  uint8
}

// метод Human - вывод информации о человеке
func (h Human) printInfo() {
	fmt.Printf("This human is %v, %v y.o.\n", h.Name, h.Age)
}

// метод Human - инкрементировать возраст и поздравить с днём рождения
func (h *Human) bDay() {
	h.Age++
	fmt.Printf("Happy Birthday, %v! You are %v y.o. now!\n", h.Name, h.Age)
}

// структура - "наследник"
type Action struct {
	Human
}

func main() {

	// объявляем h типа Human и вызываем все ее методы
	h := Human{"Test", 1}
	h.printInfo()
	h.bDay()

	// объявляем a типа Action и вызываем все методы "базовой" структуры Human
	a := Action{Human: Human{"Test-Action", 3}}
	a.printInfo()
	a.bDay()

}
