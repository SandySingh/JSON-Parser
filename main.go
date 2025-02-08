package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]
	res, err := readAndParse(filePath)
	if err != nil {
		os.Exit(1)
	}
	if res {
		fmt.Println("JSON is valid")
	} else {
		fmt.Println("JSON is not valid")
		os.Exit(1)
	}
}

func readAndParse(filePath string) (bool, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file")
		return false, err
	}
	return parse(file), nil
}
