package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	for i := 1; i < 1000; i++ {
		fmt.Println("---")
		for _, v := range os.Environ() {
			fmt.Fprintf(os.Stderr, "%s\n", v)
		}
		time.Sleep(time.Second)
	}

}
