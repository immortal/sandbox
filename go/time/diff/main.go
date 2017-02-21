package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Unix(1411691219, 0)

	diff := time.Since(start)
	fmt.Printf("diff = %+v\n", diff)
	fmt.Printf("days = %d\n", diff/(24*time.Hour))
	hours := diff % (24 * time.Hour)
	fmt.Printf("hours = %d\n", hours/time.Hour)
	minutes := hours % time.Hour
	fmt.Printf("minutes = %d\n", minutes/time.Minute)
	seconds := int(minutes.Seconds()) % 60
	fmt.Printf("seconds = %d\n", seconds)

}
