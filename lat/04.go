package lat

import (
	"fmt"
	"time"
)

func LatEmpat() {
	// penggunaan time.Duration yang sederhana
	now := time.Now()

	duration, _ := time.ParseDuration("1h1s")

	for i := duration; i > time.Second; i = i - time.Minute {
		now := time.Now()
		fmt.Println(now.Format(time.DateTime))
		time.Sleep(time.Second)
	}

	duration = time.Since(now)
	fmt.Println(duration)
	fmt.Println("time elapsed in seconds", duration.Seconds())
	fmt.Println("time elapsed in minutes", duration.Minutes())
	fmt.Println("time elapsed in hours", duration.Hours())
}
