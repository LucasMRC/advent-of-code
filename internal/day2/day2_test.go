package day2

import (
	"fmt"
	"testing"
)

var tests = []struct {
	expected1 int
	expected2 int
	path      string
	cubes     CubesSubset
}{
	{
		expected1: 8,
		expected2: 2286,
		path:      "test_1.txt",
		cubes: CubesSubset{
			Red:   12,
			Green: 13,
			Blue:  14,
		},
	},
}

func TestDay2Part1(t *testing.T) {
	for _, test := range tests {
		result1 := RunOnPathWithCubes(test.path, test.cubes)
		if result1 != test.expected1 {
			fmt.Printf("Failed part 1 with value %d, expected %d\n", result1, test.expected1)
			t.Fail()
		}
		fmt.Println("Part 1: test successful")
	}
}

func TestDay2Part2(t *testing.T) {
	for _, test := range tests {
		result2 := RunOnPathAndGetPowerSum(test.path)
		if result2 != test.expected2 {
			fmt.Printf("Failed part 2 with value %d, expected %d", result2, test.expected2)
			t.Fail()
		}
		fmt.Println("Part 2: test successful")
	}
}
