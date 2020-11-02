package main

import(
	"fmt"
	"encoding/base64"
)

func main() {
	data := "razzy tirta"

	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	decodeByte, _ := base64.StdEncoding.DecodeString(encodedString)
	decodedString := string(decodeByte)
	fmt.Println(decodedString)

	// penerapan fungsi encode() dan decode()
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	encodedString1 := string(encoded)
	fmt.Println(encodedString1)

	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		panic(err.Error())
	}

	decodedString2 := string(decoded)
	fmt.Println(decodedString2)

	// encode dan decode pada data url

	url := "https://kalipare.com/"

	encodedString3 := base64.StdEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedString3)

	decodeByte1, _ := base64.StdEncoding.DecodeString(encodedString3)
	decodedString3 := string(decodeByte1)
	fmt.Println(decodedString3)


}

