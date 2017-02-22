package main

import (
	"bytes"
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Unix(1411691219, 0)
	diff := time.Since(start)
	fmt.Printf("diff = %s\n", diff)
	fmt.Printf("diff = %s\n", TimeDiff(start))
}

func TimeDiff(t time.Time) string {
	diff := time.Since(t)
	days := diff / (24 * time.Hour)
	hours := diff % (24 * time.Hour)
	minutes := hours % time.Hour
	seconds := math.Mod(minutes.Seconds(), 60)
	var buffer bytes.Buffer
	if days > 0 {
		buffer.WriteString(fmt.Sprintf("%dd", days))
	}
	if hours/time.Hour > 0 {
		buffer.WriteString(fmt.Sprintf("%dh", hours/time.Hour))
	}
	if minutes/time.Minute > 0 {
		buffer.WriteString(fmt.Sprintf("%dm", minutes/time.Minute))
	}
	if seconds > 0 {
		buffer.WriteString(fmt.Sprintf("%.1fs", seconds))
	}
	return buffer.String()
}
