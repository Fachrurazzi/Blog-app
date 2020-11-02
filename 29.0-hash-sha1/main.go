package main

import (
	"fmt"
	"crypto/sha1"
	"time"
)


// penerapan hash SHA1
// func main() {
// 	text := "this is secret"
// 	sha := sha1.New()
// 	sha.Write([]byte(text))
// 	encrypted := sha.Sum(nil)
// 	encryptedString := fmt.Sprintf("%x", encrypted)

// 	fmt.Println(encryptedString)
// }

// metode salting pada hash SHA1
func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
func main() {
	text := "this is secret"
	fmt.Printf("original : %s \n\n", text)

	hashed1, salt1 := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	hashed2, salt2 := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed2)

	hashed3, salt3 := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed3)

	_,_,_ = salt1, salt2, salt3
}
