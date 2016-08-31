package main

import (
	"fmt"
	"time"
)

var wait time.Duration

func main() {
	time.Sleep(wait)
	fmt.Printf("wait = %+v\n", wait)
}
