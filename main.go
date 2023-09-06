package main

import (
	"fmt"
	"golang-time/lat"
	"time"
)

func main() {
	now := time.Now()

	// time.Duration
	lat.LatEmpat()
	fmt.Println()

	// Time, Ticker, Scheduler
	lat.LatTiga()

	// parsing dari string ke time.Time dan sebaliknya
	fmt.Println()
	lat.LatDua()

	// beberapa method milik time.Time
	fmt.Println()
	lat.LatSatu()

	fmt.Println()
	time1 := time.Now().UTC()
	fmt.Printf("time1 %v\n", time1)

	time2 := time.Date(2023, 8, 26, 18, 59, 00, 00, time.Local)
	fmt.Printf("time2 %v\n", time2)

	fmt.Printf("program berjalan selama %v\n", time.Since(now).Round(time.Second))
}
