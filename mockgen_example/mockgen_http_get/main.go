package main

import (
	"fmt"
)

func main() {
	response, err := GetGoogle()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(response))
}
