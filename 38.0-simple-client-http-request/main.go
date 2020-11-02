package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

var baseUrl = "http://localhost:8080"

type student struct {
	ID string
	Name string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseUrl+"/users", nil)
	if err !=  nil {
		fmt.Println(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data,nil
}
func main()  {
	var users, err = fetchUsers()
	if err != nil {
		return
	}
		fmt.Println("Error!", err.Error())

	for _, each := range users {
		fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	}
}