package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var _, b, _, _ = runtime.Caller(0)
var base_path = filepath.Dir(b)

type card struct {
	winningNs     map[string]bool
	hand          []string
	multiplyTimes int
}

func Test() {
	// cmd := exec.Command(fmt.Sprintf("go test %s/", base_path))
	// var out strings.Builder
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(out.String())
}

func Run() {
	path := base_path + "/input.txt"

	fmt.Println("=================")
	fmt.Println("Day 4, part 1:", RunPart1OnPath(path))
	fmt.Println("Day 4, part 2:", RunPart2OnPath(path))
}

func RunPart1OnPath(path string) int {
	cards := parseInput(path)
	result := calculatePoints(cards)

	return result
}

func divideByteSlice(s []byte) [][]byte {
	result := [][]byte{}
	for i := 0; ; i += 3 {
		result = append(result, []byte{s[i], s[i+1]})
		if i == (len(s) - 2) {
			break
		}
	}

	return result
}

func parseInput(path string) []card {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	result := []card{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Bytes()

		colonIndex := bytes.Index(line, []byte{58}) + 2
		groups := bytes.Split(line[colonIndex:], []byte{32, 124, 32})

		wn := map[string]bool{}
		for _, n := range divideByteSlice(groups[0]) {
			key := strings.Trim(string(n), " ")
			if err != nil {
				panic(err)
			}
			wn[key] = false
		}
		h := []string{}
		for _, n := range divideByteSlice(groups[1]) {
			key := strings.Replace(string(n), " ", "", 1)
			if err != nil {
				panic(err)
			}
			h = append(h, key)
			if _, ok := wn[key]; ok {
				wn[key] = true
			}
		}
		card := card{
			winningNs:     wn,
			hand:          h,
			multiplyTimes: 1,
		}
		result = append(result, card)
	}

	return result
}

func calculatePoints(cards []card) int {
	result := 0
	for _, c := range cards {
		points := 0
		for _, won := range c.winningNs {
			if won {
				if points == 0 {
					points = 1
				} else {
					points = points << 1
				}
			}
		}
		result += points
	}
	return result
}

func RunPart2OnPath(path string) int {
	cards := parseInput(path)
	result := calculateCards(cards)

	return result
}

func calculateCards(cards []card) int {
	result := len(cards)
	for i, c := range cards {
		j := 0
		for _, w := range c.winningNs {
			if w {
				j++
				cards[i+j].multiplyTimes += cards[i].multiplyTimes
			}
		}
		result += j * c.multiplyTimes
	}

	return result
}
