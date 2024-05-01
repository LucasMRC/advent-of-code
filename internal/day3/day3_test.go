package day3

import (
	"fmt"
	"testing"
)

func TestRunDay3(t *testing.T) {
	tests := []struct {
		path      string
		expected1 int
		expected2 int
	}{
		{
			path:      "test_1.txt",
			expected1: 4361,
			expected2: 467835,
		},
	}

	for _, test := range tests {
		result1 := RunPart1OnPath(test.path)
		if result1 != test.expected1 {
			fmt.Printf("Test failed with result %d, expected %d\n", result1, test.expected1)
			t.Fail()
		}
		result2 := RunPart2OnPath(test.path)
		if result2 != test.expected2 {
			fmt.Printf("Test failed with result %d, expected %d\n", result2, test.expected2)
			t.Fail()
		}
	}
}
