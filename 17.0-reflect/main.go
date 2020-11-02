package main

import f "fmt"
import "reflect"

// func main() {
// 	// mencari tipe data & value menggunakan reflect
// 	var number = 23
// 	var reflectValue = reflect.ValueOf(number)

// 	f.Println("tipe variabel :", reflectValue.Type())

// 	if reflectValue.Kind() == reflect.Int {
// 		f.Println("nilai variabel :", reflectValue.Int())
// 		// pengaksesan nilai dalam bentuk inteface{}
// 	}
// 		f.Println("nilai variabel :", reflectValue.Interface())
// }

// pengaksesan informasi property varibel objek
type student struct {
	Name string
	Grade int
}

// func (s *student) getPropertyInfo() {
// 	var reflectValue = reflect.ValueOf(s)

// 	if reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}

// 	var reflectType = reflectValue.Type()

// 	for i := 0; i < reflectValue.NumField(); i++ {
// 		f.Println("nama		:", reflectType.Field(i).Name)
// 		f.Println("tipe data:", reflectType.Field(i).Type)
// 		f.Println("nilai	:", reflectValue.Field(i).Interface())
// 	}
// }

// func main() {
// 	var s = &student{Name: "Razzi", Grade : 2}
// 	s.getPropertyInfo()
// }

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "Razzi Tirta", Grade: 2}
	f.Println("nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Tirta"),
	})

	f.Println("nama :", s1.Name)
}