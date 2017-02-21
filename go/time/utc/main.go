package main

import (
	"fmt"
	"time"
)

func main() {

	i := time.Now().UTC().Unix()
	y := time.Now().Unix()
	fmt.Println(i)
	fmt.Println(y)
}
