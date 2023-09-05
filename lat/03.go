package lat

import (
	"fmt"
	"os"
	"time"
)

func timeAfter() {
	now := time.Now()
	h, m, s := now.Clock()
	fmt.Printf("%d:%d:%d:%d\n", h, m, s, now.Nanosecond()/1000000)

	<-time.After(time.Second * 3)

	now = time.Now()
	h, m, s = now.Clock()
	fmt.Printf("%d:%d:%d:%d\n", h, m, s, now.Nanosecond()/1000000)
}

func timeAfterFunc() {
	ch := make(chan bool)

	// parameter kedua dari function AfterFunc berjalan secara asynchronous
	time.AfterFunc(time.Second*4, func() {
		ch <- true
		fmt.Println("expired DxD", time.Now())
	})

	fmt.Println("start time.AfterFunc", time.Now().Format(time.TimeOnly))
	<-ch
	fmt.Println("finish time.AfterFunc", time.Now().Format(time.TimeOnly))
}

func timeNewTimer() {
	timer := time.NewTimer(time.Second * 4)
	fmt.Println("start time.NewTimer", time.Now().Format(time.TimeOnly))
	c := <-timer.C
	fmt.Println("isi data channel C :", c.Format(time.TimeOnly))
	fmt.Println("finish time.NewTimer", time.Now().Format(time.TimeOnly))
	timer.Reset(time.Second * 2)
	c = <-timer.C
	fmt.Println("NewTimer", c.Format(time.TimeOnly))
	fmt.Println("finish time.NewTimer", time.Now().Format(time.TimeOnly))
}

func timeSleep() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 second")

	// scheduler sederhana menggunakan time.Sleep()
	for i := 0; i < 5; i++ {
		fmt.Println(time.Now().UnixMilli())
		time.Sleep(time.Second)
	}
}

func timeTicker() {
	fmt.Println("ini ticker", time.Now())
	ticker := time.NewTicker(time.Second.Abs() * 2)
	done := make(chan bool)

	go func() {
		time.Sleep(10 * time.Second.Abs())
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			// return digunakan untuk keluar dari statement for - select
			// agar tidak terjadi deadlock
			return
		case t := <-ticker.C:
			fmt.Println("ini ticker", t)
		}
	}
}

func timeTicker2() {
	ticker := time.NewTicker(time.Second)
	duration, _ := time.ParseDuration("10s")

	// defer function for handling ticker.Reset()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic error terjadi dengan message", r)
		}
	}()

	for i := time.Second; i < duration; i = i + time.Second {
		c := <-ticker.C
		fmt.Println(c.Format(time.DateTime))

		if i == duration-time.Second {
			ticker.Stop()
			fmt.Println("ticker is stop")
		}
	}

	ticker.Reset(time.Second - time.Second)
	fmt.Println("apa bisa ?", <-ticker.C)
}

// kombinasi timer dan goruntine
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out! no answer more than", timeout, "seconds")
	os.Exit(1)
}

func LatTiga() {
	// kombinasi timer dan goruntine
	timeout := 5
	ch := make(chan bool)
	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725 / 25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}

	// scheduler menggunakan ticker
	timeTicker2()
	timeTicker()

	// time.Sleep(d time.Duration)
	timeSleep()

	// penggunaan function time.NewTimer()
	timeNewTimer()

	// scheduler dengan menggunakan time.AfterFunc()
	timeAfterFunc()

	// penggunaan function time.After(d time.Duration)
	timeAfter()
}
