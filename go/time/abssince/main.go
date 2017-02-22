package main

import (
	"fmt"
	"time"
)

func AbsSince(t time.Time) string {
	const (
		Decisecond = 100 * time.Millisecond
		Day        = 24 * time.Hour
	)
	ts := time.Since(t) + Decisecond/2
	d := ts / Day
	ts = ts % Day
	h := ts / time.Hour
	ts = ts % time.Hour
	m := ts / time.Minute
	ts = ts % time.Minute
	s := ts / time.Second
	ts = ts % time.Second
	f := ts / Decisecond
	return fmt.Sprintf("%dd%dh%dm%d.%ds", d, h, m, s, f)
}

func main() {
	start := time.Unix(1487782985, 0)
	diff := time.Since(start)
	fmt.Printf("diff = %s\n", diff)
	fmt.Printf("diff = %s\n", AbsSince(start))
}
