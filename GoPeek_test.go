package main

import (
	"testing"
	"fmt"
)

// The file that will be used for testing
var fileTestPath = "Test/InputFiles/TestFile.go"

func TestFindGoFile(t *testing.T) {
	expected := fileTestPath

	testDir := "Test/InputFiles"
	testFile := "TestFile.go"

	result := findGoFile(testDir, testFile)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestGetCommentCount(t *testing.T) {
	expected := 7

	result := getCommentCount(fileTestPath)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestIsFuncPrivate(t *testing.T) {
	tests := []struct {
		flag   	  string
		expected  bool
	} {
		// Test inputs
		{"openFile(file string)", true},
		{"readFile()", true},
		{"ShowUI()", false},
		{"AddNum(a int, b int)", false},
	}

	// Test all
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.flag), func(t *testing.T) {

			// True - private
			// False - public
			result := isFuncPrivate(test.flag)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}

		})
	}
}