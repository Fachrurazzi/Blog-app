package main
// decode array json ke array objek
import(
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	jsonString := `[
		{"Name": "Fachrurazzi", "Age": 22},
		{"Name": "Marhamah", "Age": 22}
	]`

	var data []User

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1 :", data[0].FullName)
	fmt.Println("user 2 :", data[1].FullName)

	// encode objek ke json string

	object := []User{{"Fachrurazzi", 22}, {"Marhamah", 22}}
	jsonData, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString1 = string(jsonData)
	fmt.Println(jsonString1)
}