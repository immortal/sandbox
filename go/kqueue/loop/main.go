package main

import (
	"fmt"
)

func main() {

	kevents := []string{"a", "b", "c"}
	for len(kevents) > 0 {

		fmt.Println(kevents[0])

		// Move to next event
		kevents = kevents[1:]
	}
}
