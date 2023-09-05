package lat

import (
	"fmt"
	"golang-time/lib"
	"time"
)

func LatDua() {
	// parsing dari string ke time.Time
	var layoutFormat, value string

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	lib.ParseTime(layoutFormat, value)

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	lib.ParseTime(layoutFormat, value)

	// error
	value = "2015-09-02 08:04:00"
	lib.ParseTime(layoutFormat, value)

	// menggunakan predefined layout umum
	data, err := time.Parse(time.RFC3339, "2023-08-30T14:20:00+07:00")
	if err == nil {
		fmt.Println(data.String())
	} else {
		fmt.Println(err)
	}

	// parsing dari time.Time ke string
	var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Printf("date\t->\t%v\ndateS1\t->\t%v\n", date, dateS1)
}
