package main

import (
	"fmt"
	// "strconv"
)

func main() {
	// // fungsi strconv.Atoi() string ke int
	// var str ="124"
	// var num, err = strconv.Atoi(str)
	// if err == nil {
	// 	fmt.Println(num)
	// }

	// // fungsi strconv.Itoa() int ke string
	// var number = 123
	// var huruf = strconv.Itoa(number)
	// fmt.Println(huruf)

	// // fungsi strconv.ParseInt()
	// var str1 = "123"
	// var num1, err1 = strconv.ParseInt(str1, 10, 64)
	// if err1 == nil {
	// 	fmt.Println(num1)
	// }

	// // fungsi strconv.FormatInt()
	// var num2 = int64(24)
	// var str2 = strconv.FormatInt(num2, 8)
	// fmt.Println(str2)

	// // fungsi strconv.ParseFloat()
	// var str3 = "12.23"
	// var num3, err2 = strconv.ParseFloat(str3, 32)

	// if err2 == nil {
	// 	fmt.Println(num3)
	// }

	// //fungsi strconv.FormatFloat()
	// var num4 = float64(24.12)
	// var str4 = strconv.FormatFloat(num4, 'f', 6, 64)
	// fmt.Println(str4)

	// // fungsi strconv.ParseBool()
	// var str5 = "true"
	// var bul, err3 = strconv.ParseBool(str5)

	// if err3 == nil {
	// 	fmt.Println(bul)
	// }

	// // fungsi strconv.FormatBool()
	// var bul1 = true
	// var str6 = strconv.FormatBool(bul1)
	// fmt.Println(str6)

	// //konversi data menggunakan casting
	// var a float64 = float64(24)
	// fmt.Println(a)

	// var b int32 = int32(24.00)
	// fmt.Println(b)

	// //casting string <-> byte
	// var text = "hello"
	// var b1 = []byte(text)

	// fmt.Printf("%d %d %d %d\n",b1[0], b1[1], b1[2], b1[3])

	// var byte1 = []byte{104, 97, 108, 111}
	// var s = string(byte1)

	// fmt.Println("%s\n", s)

	// var c int64 = int64('h')
	// fmt.Println(c)

	// var d string = string(104)
	// fmt.Println(d)

	var data = map[string]interface{}{
		"name": "razzi tirta",
		"grade": 2,
		"height": 167.4,
		"isMale": true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["name"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))

		}
	}
}