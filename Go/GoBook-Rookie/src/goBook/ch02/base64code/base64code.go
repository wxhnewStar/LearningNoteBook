package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "Away keyboard. https://golang.org/"

	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))

	fmt.Println(encodedMessage)
	fmt.Printf("len:%d\n", len([]byte(encodedMessage)))

	data, err := base64.StdEncoding.DecodeString(encodedMessage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
		fmt.Printf("len:%d\n", len([]byte(data)))
	}
}
