package main

import (
	"fmt"
)

type MainInterface interface {
	SubInterfaceA
	SubInterfaceB
}

type SubInterfaceA interface {
	MethodA()
	GetterA(s implementMain)
}

type SubInterfaceB interface {
	MethodB()
	GetterB(s implementMain)
}

type implementA struct{}

func (ia *implementA) MethodA() { fmt.Println("I am method A") }
func (ia *implementA) GetterA(s implementMain) {
	fmt.Println(s.Data)
}

type implementB struct{}

func (ib *implementB) MethodB() { fmt.Println("I am method B") }
func (ib *implementB) GetterB(s implementMain) {
	fmt.Println(s.Data)
}

type implementMain struct {
	Data string
	SubInterfaceA
	SubInterfaceB
}

func New(d string) implementMain {
	return implementMain{
		Data:          d,
		SubInterfaceA: &implementA{},
		SubInterfaceB: &implementB{},
	}
}

func main() {
	var m MainInterface

	m = New("something")

	fmt.Println(m.(implementMain).Data)

	m.MethodA()                  // prints I am method A
	m.MethodB()                  // prints I am method B
	m.GetterA(m.(implementMain)) // prints "something"
	m.GetterB(m.(implementMain)) // prints "something"
}
