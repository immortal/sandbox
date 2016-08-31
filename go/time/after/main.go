package main

import (
	"fmt"
	"time"
)

func myFunc() error {
	for i := 1; i < 10; i++ {
		fmt.Printf("i = %+v\n", i)
		if i%3 == 0 {
			return fmt.Errorf("error")
		}
	}
	return nil
}

func main() {

	run := make(chan struct{}, 1)

	run <- struct{}{}
	for {
		select {
		case <-run:
			err := myFunc()
			if err != nil {
				time.AfterFunc(3*time.Second, func() {
					run <- struct{}{}
				})
			}
		default:
		}
	}
}
