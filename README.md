# date time

go menyediakan package untuk keperluan pemanfaatan _date-time_ salahsatu nya adalah `time.Time` yang berupakan tipe untuk tanggal dan waktu di go.

> time disini maksudnya adalah gabungan antara `date` dan `time`, bukan hanya waktu saja.

## time.Time

tipe `time.Time` merupakan representasi untuk object date-time, ada 2 cara untuk membuat data bertipe ini :
1. menjadikan informasi waktu sekarang sebagai object `time.Time` menggunakan function `time.Now()`
2. membuat object baru bertipe `time.Time` dengan informasi ditentukan sendiri menggunakan `time.Date()`

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  time1 := time.Now().UTC()
	fmt.Printf("time1 %v\n", time1)

	time2 := time.Date(2023, 8, 26, 18, 59, 00, 00, time.Local)
	fmt.Printf("time2 %v\n", time2)
}
```

function `time.Now()` mengembalikan object `time.Time` dengan informasi date-time sesuai dengan waktu saat statement tersebut dijalankan. 
Sedangkan function `time.Date()` digunakan untuk membuat object `time.Time` baru dimana informasi date-time nya ditentukan sendiri. 
Function ini memiliki 8 buah parameter yang _mandatory_ (wajib) dengan skema sebagai berikut

```go
time.Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location)
```

hasil dari function `time.Now()`, informasi timezone nya adalah relatif dari lokasi kita, untuk penentuan timezone seperti pada variabel `time` selain menggunakan function `UTC()` juga bisa dengan menggunakan function `Local()` yang nilainya relatif dengan date-time local kita.

## method milik struct `time.Time`

selain method `time.Now()` dan `time.Date()` ada beberapa lagi method pada struct `time.Time` ini diantaranya :

|method|return type|penjelasan|
|------|-----------|----------|
|`now.Year()`|_int_|mengembalikan tahun sekarang|
|`now.YearDay()`|_int_|mengembalikan hari ke-? dimulai dari awal tahun|
|`now.Month()`|_int_|mengembalikan bulan dimana dimulai dari 1|
|`now.Weekday()`|_int_|mengembalikan nama hari dalam bahasa inggris, untuk mengambil bentuk string-nya bisa dengan menambahkan method `.String()`|
|`now.ISOWeek()`|(_int_, _int_)|mengembalikan tahun dan minggu ke-? dimulai dari awal tahun|
|`now.Day()`|_int_|mengembalikan tanggal|
|`now.Hour()`|_int_|mengembalikan jam|
|`now.Minute()`|_int_|mengembalikan menit|
|`now.Second()`|_int_|mengembalikan detik|
|`now.Nanosecond()`|_int_|mengembalikan nano detik|
|`now.Local()`|_time.Time_|mengembalikan date time sekarang dalam timezone local|
|`now.Location()`|_*time.Location_|mengambil informasi lokasi, apakah _utc_ atau _local_, untuk mengambil bentuk string-nya bisa dengan menambahkan method `.String()`|
|`now.Zone()`|(_string_, _int_)|mengembalikan informasi timezone dalam bentuk string, dan offset dalam bentuk int second, contoh `WIB`, `25200`|
|`now.IsZero()`|_bool_|cek apakah variabel _now_ adalah 1 januari tahun 1, 00:00:00 UTC, jika iya maka return _true_|
|`now.UTC()`|_time.Time_|mengembalikan date time dalam timezone UTC|
|`now.Unix()`|_int64_|mengembalikan date time dalam format unix time|
|`now.UnixNano()`|_int64_|mengambil date time dalam format unix time, informasi nanosecond juga dimasukkan|
|`now.String()`|_string_|mengembalikan date time dalam string|

## parsing dari _string_ ke _time.Time_

untuk konversi dari _string_ ke _time.Time_ bisa memanfaatkan function `time.Parse()`, function ini membutuhkan 2 parameter :

- parameter ke 1 adalah layout format data waktu yang akan di parsing
- parameter ke 2 adalah data string yang akan di parsing

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  var layoutFormat, value string
  var data time.Time

  layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
  data, _ = time.Parse(layoutFormat, value)
  fmt.Printf("%s\t->\t%v\n", value, data)

  layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	data, _ = time.Parse(layoutFormat, value)
  fmt.Printf("%s\t->\t%v\n", value, data)
}
```

Go memiliki standar layout format yang cukup unik, contohnya seperti pada kode di atas _"2006-01-02 15:04:05"_. 
Go menggunakan _2006_ untuk parsing tahun, bukan _YYYY_; _01_ untuk parsing bulan; _02_ untuk parsing tanggal, dst. 
Detail-nya bisa dilihat pada table berikut :

|layout format|penjelasan|contoh data|
|-------------|----------|-----------|
|2006|tahun 4 digit|2006|
|006|tahun 3 digit|006|
|06|tahun 2 digit|05|
|01|bulan 2 digit|05|
|1|bulan 1 digit jika di bawah 10, selain itu 2|5, 12|
|January|nama bulan dalam bahasa inggris|September,August|
|Sep|nama bulan dalam bahasa inggris, 3 huruf|Sep, Aug|
|02|tanggal 2 digit|02|
|2|tanggal 1 digit jika dibawah 10, selain itu 2 digit|8, 12|
|Monday| nama hari dalam bahasa inggris|Saturday, Friday|
|Mon|nama hari dalam bahasa inggris 3 huruf| Sat, Fri|
|15|jam dalam format **24 jam**|18|
|03|jam dalam format **12 jam** 2 digit|05,11|
|3|jam dalam format **12 jam** 1 digit jika dibawah 10, selainnya 2 digit|5,11|
|PM|AM/PM biasa digunakan dengan format jam **12 jam**|PM, AM|
|02|menit 2 digit|08|
|4|menit 1 digit jika dibawah 10, selainnya 2 digit|6,36|
|02|detik 2 digit|06|
|5| detik 1 digit jika dibawah 10, selainnya 2 digit|6,36|
|999999|nano detik|124006|
|MST|lokasi timezone|UTC,WIB,EST|
|Z0700|offset time zone|Z,+0700,-0200|

## parsing dari _time.Time_ ke _string_

untuk memparsing dari _time.Time_ ke _string_ kita bisa menggunakan method `Format()` milik struct `time.Time` dimana kita bisa menentukan layout format yang diinginkan.

```go
package main

import (
  "fmt"
  "time"
)
func main(){
  now := time.Now()
  rfc3339 := now.Format(time.RFC3339)
  fmt.Println("rfc3339",rfc3339)
}
```

## Time, Ticker, dan Scheduler

beberapa function / method yang bisa digunakan untuk menunda atau menjadwalkan eksekusi sebuah statement dalam jeda waktu tertentu.

### function `time.Sleep(d time.Duration)`

function ini bersifat **blocking** statement di bawahnya tidak akan di eksekusi sampai durasi yang telah di tentukan selesai.

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  fmt.Println("Starting...")
  time.Sleep(time.Second * 4)
  fmt.Println("after 4 second")
}
```
saat dijalankan tulisan _after 4 second_ akan muncul 4 detik kemudian setelah tulisan _Starting..._.

### function `time.NewTimer(d time.Duration)`

function ini mereturn _struct_ bertipe _*time.Timer_ yang memiliki properti _C_ bertipe receive only channel. 
cara kerjanya setelah jeda waktu yang diberikan, sebuah data akan di kirim lewat channel _C_ & penggunaan function ini harus diikuti dengan statement penerimaan data dari channel _C_ dimana isi data dari channel _C_ adalah sebuah struct _time.Time_ dimana informasi dateTime nya tepat ketika data dari channel _C_ diterima.

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  timer := time.NewTimer(time.Second * 4)
  fmt.Println("Starting time.NewTimer ...")
  <-timer.C
  fmt.Println("finish after 4 second")
}
```

pada contoh di atas baris code `<-timer.C` menandakan penerimaan data dari channel _C_, karena penerimaan data dari channel bersifat blocking, maka statement setelahnya baru akan di eksekusi setelah 4 detik.


### function `time.AfterFunc(d time.Duration, f func())`

pada function `time.AfterFunc()` parameter kedua yang bertipe function, akan di eksekusi setelah durasi timer yang diberikan oleh parameter pertama habis.

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  ch := make(chan bool)

  // function beriktu bersifat asynchronous (tidak blocking)
  time.AfterFunc(time.Second * 2, func (){
    ch <- bool
    fmt.Println("expired")
  })

  fmt.Println("starting with AfterFunc")
  <-ch // bersifat blocking (synchronous)
  fmt.Println("finish with AfterFunc")


  // menggunakan cara kedua
  time.AfterFunc(time.Second * 2, func(){
    fmt.Println("selesai eksekusi AfterFunc", time.Now().Format(time.RFC822))
  })
  fmt.Println("Hello")
  time.Sleep(time.Second * 3) // bersifat blocking (synchronous)
  fmt.Println("World")
}
```

pada contoh kode di atas, function `AfterFunc` dijalankan dua kali, saat pertama kali dijalankan, setelah statement `fmt.Println("starting with AfterFunc")` dijalankan, selang waktu 2 detik muncul tulisan _"expired"_, 
didalam function callback tersebut terjadi proses pengiriman data lewat channel, menjadikan statement `fmt.Println("finish with AfterFunc")` dijalankan tepat setelah tulisan _"expired"_ muncul.

Lalu saat kedua kali function `AfterFunc` dijalankan, setelah tulisan _"Hello"_ muncul, hasil dari statement `fmt.Println("selesai eksekusi AfterFunc", time.Now().Format(time.RFC822))` muncul.

function `time.AfterFunc` bersifat asynchronous, dimana setelah batas waktu yang telah di tentukan di parameter pertama selesai, callback function dijalankan. 
Jika pada callback function terjadi serah terima data lewat channel, maka function akan tetap berjalan secara asynchronous, hingga dimana statement penerimaan data dilakukan. 
Proses blocking nya terjadi pada statement penerimaan data dari channel.

### function `time.After(d time.Duration)`

kegunaan function `time.After` mirip dengan `time.Sleep`, perbedaannya function ini mengembalikan data channel, sehingga dalam penerapannya perlu menggunakan keyword `<-`

```go
package main

import (
  "fmt"
  "time"
)

func main(){
  now := time.Now()
	h, m, s := now.Clock()
	fmt.Printf("%d:%d:%d:%d\n", h, m, s, now.Nanosecond()/1000000)

	<-time.After(time.Second * 3)

	now = time.Now()
	h, m, s = now.Clock()
	fmt.Printf("%d:%d:%d:%d\n", h, m, s, now.Nanosecond()/1000000)
}
```

### scheduler menggunakan ticker

Selain function function untuk keperluan timer, go juga menyediakan function untuk keperluan scheduler ( ticker ). 
Untuk penggunaan ticker, buat variabel ticker baru dengan menggunakan function `time.NewTicker(d time.Duration)`, dari variabel tersebut kita bisa akses properti _C_ yang merupakan sebuah _channel_, setiap durasi yang sudah ditentukan, variabel bertipe ticker tersebut akan mengirimkan informasi date-time via channel tersebut.

```go
package main

import (
  "fmt"
  "time"
)

func ticker1(){
  ticker := time.NewTicker(time.Second)
  duration, _ := time.ParseDuration("10s")

  for i := time.Second; i < duration; i = i + time.Second {
    fmt.Println(<-ticker.C)

    if i == duration - time.Second {
			ticker.Stop()
			fmt.Println("ticker is stop")
		}
  }
}

func ticker2(){
  ticker := time.NewTicker(time.Second)
  done := make(chan bool)

  go func() {
    time.Sleep(time.Second * 5)
    done <- true
  }()

  for {
    select {
      case <-done :
        ticker.Stop()
        return
      case t := <-ticker.C :
        fmt.Println(t)
    }
  }
}

func main(){
  ticker1()
  ticker2()
}
```

dari 2 contoh di atas terdapat 2 statement yang fungsinya sama, yaitu mematikan ticker menggunakan `ticker.Stop()`. 
By default channel _ticker.C_ akan menerima kiriman data setiap N duration, dimana di contoh di atas ada yang menggunakan 5 second dan 1 second. 
Data yang di kirim ke channel _ticker.C_ berupa date-time kapan event tersebut terjadi


## Time Duration

Tipe `Time.Duration` ini merepresentasikan durasi, contohnya seperti 1 jam 11 menit dst, date denan tipe ini bisa di hasilkan dari operasi pencarian delta atau selisih dari dua buah variabel struct `time.Time` atau juga bisa kita buat sendiri.

Tipe data durasi adalah `time.Duration` yang sebenarnya tipe ini merupakan tipe buatan baru dari `int64`, 
ada beberapa _predefined_ konstanta durasi yang perlu di ketahui :

- `time.Nanosecond` yang nilainya adalah `1`
- `time.Microsecond` yang nilainya adalah `1000`, atau `1000` x `time.Nanosecond`
- `time.Milisecond` yang nilainya adalah `1000000`, atau `1000` x `time.Microsecond`
- `time.Second` yang nilainya adalah `1000000000`, atau `1000` x `time.Milisecond`
- `time.Minute` yang nilainya adalah `1000000000000`, atau `1000` x `time.Second`
- `time.Hour` yang nilainya adalah `1000000000000000`, atau `1000` x `time.Minute`

dari list di atas bisa dicontohkan bahwa sebuah data dengan tipe `time.Duration` yang nilainya 1, maka artinya durasi adalah **1 nanosecond**.

```go
package main

import (
  "fmt"
  "time"
)

func main( ){
  start := time.Now()

  time.Sleep(time.Second * 5)

  duration := time.Since(start)

  fmt.Println(duration)
  fmt.Println("time elapsed in seconds",duration.Seconds())
  fmt.Println("time elapsed in minutes",duration.Minutes())
  fmt.Println("time elapsed in hours",duration.Hours())
}
```

pada contoh code di atas variabel `duration` berisi lama waktu antara kapan variabel `start` di inisialisasi hingga kapan variabel `duration` ini statement nya di eksekusi. 

Cara menghitung durasi bisa dengan menggunakan function `time.Since(t time.Time)`, isi parameter pada function tersebut dengan variabel bertipe `time.Time`, maka durasi antara waktu pada argument vs ketika statement `time.Since()` akan di hitung.

### beberapa method `time.Duration`

tipe durasi `time.Duration` memiliki beberapa mthod yang sangat berguna untuk keperluan menngambil nilai durasinya nya kedalam unit tertentu, misalnya seperti pada contoh di atas, object variabel `duration` diambil nilainya dalam satuan detik menggunakan method `Seconds()`, diambil satuan menit menggunakan method `Minutes()`, dan di ambil dalam satuan jam menggunakan method `Hours()` DLL.