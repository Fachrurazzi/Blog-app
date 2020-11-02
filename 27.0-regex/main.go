package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

	// Method MatchString()
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	// Method FindString()
	var str = regex.FindString(text)
	fmt.Println(str)

	//Method FindStringIndex()
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)

	var str1 = text[0:6]
	fmt.Println(str1)

	// Method FindAllString()
	var str2 = regex.FindAllString(text, -1)
	fmt.Println(str2)

	var str3 = regex.FindAllString(text, 1)
	fmt.Println(str3)

	// Method ReplaceAllString()
	var str4 = regex.ReplaceAllString(text, "potato")
	fmt.Println(str4)

	//Method ReplaceAllStringFunc()
	var str5 = regex.ReplaceAllStringFunc(text, func (each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str5)

	//Method Split()
	var regex1,_ = regexp.Compile(`[a-b]+`)

	var str6 = regex1.Split(text, -1)
	fmt.Print("%#v \n", str6)


}