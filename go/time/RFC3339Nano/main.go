package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Printf("t = %+v\n", t)
	t2 := time.Now().Format(time.RFC3339Nano)
	fmt.Printf("t2 = %+v\n", t2)
}
