package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	for i := 1; i < 10; i++ {
		if i%3 == 0 {
			fmt.Fprintf(os.Stderr, "STDERR i: %d\n", i)
		} else {
			fmt.Printf("STDOUT i: %d\n", i)
		}
		time.Sleep(time.Second)
	}

}
