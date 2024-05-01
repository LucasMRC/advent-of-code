package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	path := "internal/day3/input.txt"

	fmt.Println("=================")
	fmt.Println("Day 3, part 1:", RunPart1OnPath(path))
	fmt.Println("Day 3, part 2:", RunPart2OnPath(path))
}

func RunPart1OnPath(path string) int {
	result := 0
	file, err := os.ReadFile(path)
	if err != nil {
		errMsg := fmt.Errorf("Error opening input file on day 3: %s", err.Error())
		panic(errMsg)
	}
	numberRegx := regexp.MustCompile("\\d+")
	symbolRegx := regexp.MustCompile("[^0-9\\.]")
	fileLines := strings.Split(string(file), "\n")
	for i, line := range fileLines {
		numbers := numberRegx.FindAllStringIndex(line, -1)
		hasSymbolBefore := false
		hasSymbolAfter := false
		hasSymbolAbove := false
		hasSymbolBelow := false

		for _, number := range numbers {
			start := number[0]
			end := number[1]
			checkPrevLine := i != 0
			checkNextLine := i < len(fileLines)-1 && len(fileLines[i+1]) != 0
			checkPrevIndex := start != 0
			checkNextIndex := end < len(line)-1

			if checkPrevIndex {
				hasSymbolBefore = symbolRegx.Match([]byte{line[start-1]})
			}
			if checkNextIndex {
				hasSymbolAfter = symbolRegx.Match([]byte{line[end]})
			}
			if checkPrevLine {
				var substr string
				prevLine := fileLines[i-1]
				if checkPrevIndex {
					if end < len(prevLine) {
						substr = prevLine[start-1 : end+1]
					} else {
						substr = prevLine[start-1 : end-1]
					}
				} else {
					if end < len(prevLine) {
						substr = prevLine[start : end+1]
					} else {
						substr = prevLine[start : end-1]
					}
				}
				hasSymbolAbove = symbolRegx.Match([]byte(substr))
			}
			if checkNextLine {
				var substr string
				nextLine := fileLines[i+1]
				if checkPrevIndex {
					if end < len(nextLine) {
						substr = nextLine[start-1 : end+1]
					} else {
						substr = nextLine[start-1 : end]
					}
				} else {
					if end < len(nextLine) {
						substr = nextLine[start : end+1]
					} else {
						substr = nextLine[start:end]
					}
				}
				hasSymbolBelow = symbolRegx.Match([]byte(substr))
			}
			if hasSymbolAbove || hasSymbolBefore || hasSymbolAfter || hasSymbolBelow {
				n, err := strconv.Atoi(line[start:end])
				if err == nil {
					result += n
				}
			}
		}
	}

	return result
}

func RunPart2OnPath(path string) int {
	result := 0
	file, err := os.ReadFile(path)
	if err != nil {
		errMsg := fmt.Errorf("Error opening input file on day 3: %s", err.Error())
		panic(errMsg)
	}
	gearRegx := regexp.MustCompile("\\*")
	fileLines := strings.Split(string(file), "\n")
	for i, line := range fileLines {
		gearMatches := gearRegx.FindAllStringIndex(line, -1)
		for _, match := range gearMatches {
			index := match[0]
			var adjacentNumbers []int
			checkPrevLine := i != 0
			checkNextLine := i < len(fileLines)-1 && len(fileLines[i+1]) != 0
			checkLineForGears(line, index, &adjacentNumbers)
			if checkPrevLine {
				checkLineForGears(fileLines[i-1], index, &adjacentNumbers)
			}
			if checkNextLine {
				checkLineForGears(fileLines[i+1], index, &adjacentNumbers)
			}
			if len(adjacentNumbers) == 2 {
				result += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	return result
}

func checkLineForGears(line string, i int, acc *[]int) {
	numberRegx := regexp.MustCompile("\\d+")
	numberMatches := numberRegx.FindAllStringIndex(line, -1)
	for _, numberMatch := range numberMatches {
		if i >= numberMatch[0]-1 && i <= numberMatch[1] {
			n, err := strconv.Atoi(line[numberMatch[0]:numberMatch[1]])
			if err == nil {
				*acc = append(*acc, n)
			}
		}
	}
}
