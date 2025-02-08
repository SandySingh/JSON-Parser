package main

import (
	"fmt"
	"testing"
)

func TestStep1(t *testing.T) {
	validFiles := []string{"tests/step1/valid.json"}
	testValidFiles(t, validFiles...)

	invalidFiles := []string{"tests/step1/invalid.json"}
	testInvalidFiles(t, invalidFiles...)
}

func testValidFiles(t *testing.T, files ...string) {
	for _, file := range files {
		testname := fmt.Sprintf("Running test for file: %s", file)
		t.Run(testname, func(t *testing.T) {
			res, err := readAndParse(file)
			if err != nil {
				t.Fatal("Some error occurred")
			}
			if !res {
				t.Errorf("JSON validation failed for file: %s", file)
			}
		})
	}
}

func testInvalidFiles(t *testing.T, files ...string) {
	for _, file := range files {
		testname := fmt.Sprintf("Running test for file: %s", file)
		t.Run(testname, func(t *testing.T) {
			res, err := readAndParse(file)
			if err != nil {
				t.Fatal("Some error occurred")
			}
			if res {
				t.Errorf("JSON validation failed for file: %s", file)
			}
		})
	}
}
