package main

import(
	"fmt"
	"encoding/json"
)

// decode json ke variabel objek struct

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "razzi tirta", "Age": 22}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user ", data.FullName)
	fmt.Println("age ", data.Age)

	// decode json ke map[string]interface{} & interface{}

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodeData = data2.(map[string]interface{})

	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age :", decodeData["Age"])
}