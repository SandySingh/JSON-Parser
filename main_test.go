package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStep1(t *testing.T) {
	dir := "tests/step1"
	files, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal("Cannot read directory")
	}

	for _, file := range files {
		testname := fmt.Sprintf("Running test for file: %s", file.Name())
		t.Run(testname, func(t *testing.T) {
			filePath := dir + "/" + file.Name()
			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("Cannot read file %s", file.Name())
			}
			if !parse(data) {
				t.Errorf("JSON not valid for file %s", file.Name())
			}
		})
	}
}
