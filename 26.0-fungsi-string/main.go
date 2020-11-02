package main

import(
	"fmt"
	"strings"
)

func main() {
	// fungsi strings.Contains()
	var isExists = strings.Contains("razzy tirta", "tirta")
	fmt.Println(isExists)

	// fungsi strings.HasPrefix()
	var isPrefix1 = strings.HasPrefix("razzy", "ra")
	fmt.Println(isPrefix1)

	var isPrefix2 = strings.HasPrefix("tirta", "ra")
	fmt.Println(isPrefix2)

	// fungsi strings.HasSuffix()
	var isSuffix1 = strings.HasSuffix("tirta", "ti")
	fmt.Println(isSuffix1)

	var isSuffix2 = strings.HasSuffix("razzy", "zy")
	fmt.Println(isSuffix2)

	// fungsi strings.Count()
	var howMany = strings.Count("razzy tirta", "a")
	fmt.Println(howMany)

	// fungsi strings.Index()
	var index1 = strings.Index("bambang suyono", "n")
	fmt.Println(index1)

	// fungsi strings.Replace()
	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1)
	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2)
	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3)

	// fungsi strings.Repeat()
	var str = strings.Repeat("na", 4)
	fmt.Println(str)

	// fungsi strings.Split()
	var string1 = strings.Split("we are young", " ")
	fmt.Println(string1)

	var string2 = strings.Split("batman", "")
	fmt.Println(string2)

	// fungsi strings.Join()
	var data = []string{"banana", "mango", "grape"}
	var str1 = strings.Join(data, "-")
	fmt.Println(str1)

	// fungsi strings.ToLower()
	var str2 = strings.ToLower("aLAY")
	fmt.Println(str2)

	// fungsi strings.ToUpper()
	var str3 = strings.ToUpper("eat!")
	fmt.Println(str3)
}