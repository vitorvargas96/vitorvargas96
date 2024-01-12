package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(content))
}
