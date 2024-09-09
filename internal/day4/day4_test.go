package day4

import (
	"fmt"
	"testing"
)

func TestRunDay4(t *testing.T) {
	tests := []struct {
		path      string
		expected1 int
		expected2 int
	}{
		{
			path:      "test_1.txt",
			expected1: 13,
			expected2: 30,
		},
	}

	for i, test := range tests {
		if result := RunPart1OnPath(test.path); result != test.expected1 {
			fmt.Println("Test 1 failed on case", i+1)
			fmt.Printf("\tExpected %d\n\tGot %d\n", test.expected1, result)
			t.Fail()
		}
		if result := RunPart2OnPath(test.path); result != test.expected2 {
			fmt.Println("Test 2 failed on case", i+1)
			fmt.Printf("\tExpected %d\n\tGot %d\n", test.expected2, result)
			t.Fail()
		}

	}
}
