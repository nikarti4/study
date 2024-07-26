package main

import (
	"fmt"
)

// объект, который мы хотим построить
type Phone struct {
	OS 				string // ОС
	year			uint16  // год выпуска
	isTouchScreen	bool   // сенсорный?
}

// интерфейс строителя
type PhoneBuilder interface {
	SetOS(OS string) PhoneBuilder
	SetYear(year uint16) PhoneBuilder
	SetTouchScreen(isTouchScreen bool) PhoneBuilder
	Build() *Phone
}


// описывает интерфейс PhoneBuilder
type phoneBuilder struct {
	phone *Phone
}

// создает нового строителя
func NewPhoneBuilder() PhoneBuilder {
	return &phoneBuilder{
		phone: &Phone{},
	}
}

func (pb *phoneBuilder) SetOS(OS string) PhoneBuilder {
	pb.phone.OS = OS
	return pb
}

func (pb *phoneBuilder) SetYear(year uint16) PhoneBuilder {
	pb.phone.year = year
	return pb
}

func (pb *phoneBuilder) SetTouchScreen(isTouchScreen bool) PhoneBuilder {
	pb.phone.isTouchScreen = isTouchScreen
	return pb 
}

func (pb *phoneBuilder) Build() *Phone {
	return pb.phone
}

// обеспечивает интерфейс строителя
type Director struct {
	builder PhoneBuilder
}


func (d *Director) MakePhone(OS string, year uint16, isTouchScreen bool) *Phone {

	d.builder.SetOS(OS).SetYear(year).SetTouchScreen(isTouchScreen)

	return d.builder.Build()
}

func main() {

	// новый строитель
	b := NewPhoneBuilder()

	// новый директор
	d := &Director{builder: b}

	// создать новый телефон
	p := d.MakePhone("Android", 2020, true)

	fmt.Println(p)
}