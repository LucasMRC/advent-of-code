package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var digitNames = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Run() {
	path := "1/input.txt"
	fmt.Println("Day 1: ", RunOnPath(path))
}

func getSubstring(str string, i, j int) string {
	var subStr string
	if j == len(str) && i < 0 {
		subStr = str
	} else if j == len(str) {
		subStr = str[i:]
	} else if i < 0 {
		subStr = str[:j+1]
	} else {
		subStr = str[i : j+1]
	}
	return subStr
}

func GetDigits(str string) string {
	line := ""
	match := false
	i := 0
	for !match && i < len(str) {
		if _, err := strconv.Atoi(string(str[i])); err == nil {
			line += string(str[i])
			match = true
			break
		}
		j := i + 1
		for !match && j <= i+5 {
			if _, err := strconv.Atoi(string(str[j])); err == nil {
				match = true
				line += string(str[j])
				break
			}
			keyMatch := false
			for key, value := range digitNames {
				if len(key) < j-i {
					continue
				}
				subStr := getSubstring(str, i, j)
				if key == subStr {
					match = true
					line += value
					break
				}
				keySubStr := getSubstring(key, 0, j-i)
				if keySubStr == subStr {
					keyMatch = true
					break
				}
				continue
			}
			if !keyMatch {
				break
			}
			j++
		}
		i++
	}
	i = len(str) - 1
	match = false
	for !match && i >= 0 {
		if _, err := strconv.Atoi(string(str[i])); err == nil {
			line += string(str[i])
			match = true
			break
		}
		j := i - 1
		for !match && j >= i-5 {
			if _, err := strconv.Atoi(string(str[j])); err == nil {
				match = true
				line += string(str[j])
				break
			}
			keyMatch := false
			for key, value := range digitNames {
				if len(key) < i-j {
					continue
				}
				subStr := getSubstring(str, j, i)
				if key == subStr {
					match = true
					line += value
					break
				}
				keySubStr := getSubstring(key, len(key)-i+j-1, len(key))
				if keySubStr == subStr {
					keyMatch = true
					break
				}
				continue
			}
			if !keyMatch {
				break
			}
			j--
		}
		i--
	}
	return line
}

func RunOnPath(path string) int {
	file, err := os.Open(path)
	if err != nil {
		newError := fmt.Errorf("Error reading input: %v", err.Error())
		panic(newError)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		text := scanner.Text()
		digits := GetDigits(text)
		values, err := strconv.Atoi(digits)
		if err != nil {
			errMsg := fmt.Errorf("Error converting from values: %v", err.Error())
			panic(errMsg)
		}
		result += values
	}
	return result
}
