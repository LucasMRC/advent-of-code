package day1

import (
	"testing"
)

func TestDay1(t *testing.T) {
	tests := []struct {
		expected int
		path     string
	}{
		{
			expected: 142,
			path:     "test_1.txt",
		},
		{
			expected: 281,
			path:     "test_2.txt",
		},
	}

	for _, test := range tests {
		result := RunOnPath(test.path)
		if result != test.expected {
			t.Fail()
		}
	}
}
