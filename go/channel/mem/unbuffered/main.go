package main

import "time"

var c = make(chan int)
var a string

func f() {
	a = "hello, world"
	println("waiting for int....")
	<-c
}

func main() {
	go f()
	time.Sleep(10 * time.Second)
	c <- 0
	print(a)
}
