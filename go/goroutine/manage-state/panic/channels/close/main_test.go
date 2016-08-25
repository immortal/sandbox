package main

import (
	"fmt"
	"testing"
)

func BenchmarkA(b *testing.B) {
	m := &myCount{
		chch: make(chan interface{}),
		quit: make(chan struct{}),
	}
	go m.loop()
	m.chch <- "start.."
	for i := 0; i < 10000000; i++ {
		m.AddA()
	}
	fmt.Printf("m.count = %+v\n", m.count)
	close(m.quit)
}

func BenchmarkB(b *testing.B) {
	m := &myCount{
		chch: make(chan interface{}),
		quit: make(chan struct{}),
	}
	go m.loop()
	m.chch <- "start.."
	for i := 0; i < 10000000; i++ {
		m.AddB()
	}
	fmt.Printf("m.count = %+v\n", m.count)
	close(m.quit)
}
