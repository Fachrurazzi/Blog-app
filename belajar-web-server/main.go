package main

import(
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"name": "Razzi Tirta",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}