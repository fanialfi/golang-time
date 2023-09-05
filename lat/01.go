package lat

import (
	"fmt"
	"time"
)

func LatSatu() {
	// beberapa method milik time.Time
	now := time.Now()
	fmt.Printf("tahun %v, hari ke %v sejak tanggal 1 januari\n", now.Year(), now.YearDay())
	fmt.Printf("hari ini hari %v\n", now.Weekday().String())
	fmt.Printf("hari ini bulan %s\n", now.Month())
	year, week := now.ISOWeek()
	fmt.Printf("hari ini tahun %d dan minggu ke %d sejak tanggal 1 januari\n", year, week)

	fmt.Printf("Hari ini tanggal %d\n", now.Day())
	fmt.Printf("hari ini jam %d:%d:%d:%d\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000)
	fmt.Printf("lokasi saya %s\n", now.Location().String())

	timeZone, offsite := now.Zone()
	fmt.Printf("lokasi saya di %s dengan offsite dengan UTC : %d\n", timeZone, offsite/3600)

	date := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(date.IsZero()) // mengembalikan true jika now sama dengan zero time sejak tanggal 1 january 1 0:0:0 UTC
	fmt.Println(now.UTC())     // mengembalikan date time sekarang dalam timezone utc
	fmt.Println(now.Local())   // mengembalikan date time sekarang dalam timezone local

	fmt.Printf("%T\t%v\n%T\t%v\n", now, now, now.String(), now.String())
}
