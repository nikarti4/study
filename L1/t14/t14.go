package main

import (
	"fmt"
	"reflect"
)

// определяем тип интерфейса в рантайме
func whatType(i interface{}) {
	fmt.Println("The type of ", i, " is ", reflect.TypeOf(i))
}

func main() {

	// объявляем интерфейсы
	var i1, i2, i3, i4 interface{}

	// присваиваем значения
	i1 = 123
	i2 = "abc"
	i3 = false
	i4 = make(chan int)

	// узнаем тип
	whatType(i1)
	whatType(i2)
	whatType(i3)
	whatType(i4)
}
