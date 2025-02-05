package main

import (
	"fmt"
	"os"
)

func main() {
	test1, err := os.ReadFile("tests/step1/valid.json")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	if parse(test1) {
		fmt.Println("JSON is valid")
	} else {
		fmt.Println("JSON is not valid")
		os.Exit(1)
	}
}
