package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Unix(1411691219, 0)

	diff := time.Since(start)
	days := diff / (24 * time.Hour)
	fmt.Printf("diff = %s\n", diff)
	fmt.Printf("days = %d\n", days)
	hours := diff % (24 * time.Hour)
	fmt.Printf("hours = %d\n", hours/time.Hour)
	minutes := hours % time.Hour
	fmt.Printf("minutes = %d\n", minutes/time.Minute)
	seconds := int(minutes.Seconds()) % 60
	fmt.Printf("seconds = %d\n", seconds)
}
