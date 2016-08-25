package main

import "fmt"

type myCount struct {
	count int
	chch  chan interface{}
	quit  chan struct{}
}

type addOne struct {
	ch chan int
}

func (m *myCount) loop() {
	for {
		select {
		case ch := <-m.chch:
			switch c := ch.(type) {
			case addOne:
				m.count++
				c.ch <- m.count
			default:
				fmt.Println(c)
			}
		case <-m.quit:
			return
		}
	}
}

func (m *myCount) AddA() (i int) {
	ch := make(chan int)
	m.chch <- addOne{ch}
	return <-ch
}

func (m *myCount) AddB() (i int) {
	ch := make(chan int)
	m.chch <- addOne{ch}
	i = <-ch
	close(ch)
	return
}

func main() {
	m := &myCount{
		chch: make(chan interface{}),
		quit: make(chan struct{}),
	}
	go m.loop()
	m.chch <- "start.."
	for i := 0; i < 10; i++ {
		fmt.Printf("m.Add() = %+v\n", m.AddA())
	}
	close(m.quit)
}
