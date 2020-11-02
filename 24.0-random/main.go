package main

import (
	"fmt"
	// "math/rand"
	"time"
)

// // random biasa yang nilai sudah bisa diketahui
// func main() {
// // 	rand.Seed(10)
// // 	fmt.Println("random ke-1", rand.Int())
// // 	fmt.Println("random ke-2", rand.Int())
// // 	fmt.Println("random ke-3", rand.Int())

// //random unique
// // rand.Seed(time.Now().UTC().UnixNano())
// // fmt.Println(rand.Int())
// // fmt.Println(rand.Int())
// // fmt.Println(rand.Int())
// 	 alphabet := randomString(10)
// 	 fmt.Println(alphabet)
// }

// var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randomString(length int) string {
// 	b := make([]rune, length)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(len(letters))]
// 	}

// 	return string(b)
// }

//time. parsing time & format time
func main() {
	time1 := time.Now()

	fmt.Printf("%v\n", time1)

	time2 := time.Date(2016, 06, 15, 8,0,0 ,0, time.Local)
	fmt.Printf("%v\n", time2)

	now := time.Now()
	fmt.Println("year :", now.Year(), "month :", now.Month())

	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2018-05-14 12:00:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "13/08/2015 WITA"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	var tanggal, _ = time.Parse(time.RFC822, "02 Aug 20 08:00 WITA")
	fmt.Println(tanggal.String())
}