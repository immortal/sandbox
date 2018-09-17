package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	lock := sync.Mutex{}
	lock.Lock()

	cond := sync.NewCond(&lock)

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		fmt.Println("First go routine has started and waits for 1 second before broadcasting condition")
		time.Sleep(1 * time.Second)
		fmt.Println("First go routine broadcasts condition")
		cond.Broadcast()
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("Second go routine has started and is waiting on condition")
		cond.Wait()
		fmt.Println("Second go routine unlocked by condition broadcast")
	}()

	fmt.Println("Main go routine starts waiting")
	waitGroup.Wait()
	fmt.Println("Main go routine ends")
}
