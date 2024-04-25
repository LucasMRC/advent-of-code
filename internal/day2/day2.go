package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubesSubset struct {
	Green int
	Blue  int
	Red   int
}

type Game map[int][]CubesSubset

func Run() {
	path := "internal/day2/input.txt"
	cubes := CubesSubset{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	fmt.Println("=================")
	fmt.Println("Day 2, part 1: ", RunOnPathWithCubes(path, cubes))
	fmt.Println("Day 2, part 2: ", RunOnPathAndGetPowerSum(path))
}

func RunOnPathWithCubes(path string, cubes CubesSubset) int {
	file, err := os.Open(path)
	if err != nil {
		errMsg := fmt.Errorf("Error reading input from day 2: %s", err.Error())
		panic(errMsg)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	game := Game{}

	for scanner.Scan() {
		text := scanner.Text()
		partial, err := parseInputLine(text)
		if err != nil {
			panic(err)
		}
		for id, cubes := range partial {
			game[id] = cubes
		}
	}

	result := getPossibleSubset(game, cubes)
	return result
}

func RunOnPathAndGetPowerSum(path string) int {
	file, err := os.Open(path)
	if err != nil {
		errMsg := fmt.Errorf("Error reading input from day 2: %s", err.Error())
		panic(errMsg)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		text := scanner.Text()
		partial, err := parseInputLine(text)
		if err != nil {
			panic(err)
		}
		minCubes := CubesSubset{}
		for _, games := range partial {
			for _, cubes := range games {
				if cubes.Blue > minCubes.Blue {
					minCubes.Blue = cubes.Blue
				}
				if cubes.Green > minCubes.Green {
					minCubes.Green = cubes.Green
				}
				if cubes.Red > minCubes.Red {
					minCubes.Red = cubes.Red
				}
			}
		}
		power := minCubes.Red * minCubes.Blue * minCubes.Green
		result += power
	}
	return result
}

func parseInputLine(line string) (Game, error) {
	values := strings.Split(strings.Replace(line, "Game ", "", 1), ": ")
	id, err := strconv.Atoi(values[0])
	game := Game{
		id: []CubesSubset{},
	}
	if err != nil {
		return nil, err
	}
	ss := strings.SplitN(values[1], "; ", -1)
	for _, s := range ss {
		var subset CubesSubset
		for _, pair := range strings.SplitN(s, ", ", -1) {
			keyValue := strings.Split(pair, " ")
			key := keyValue[1]
			value, err := strconv.Atoi(keyValue[0])
			if err != nil {
				panic(err)
			}
			if key == "green" {
				subset.Green = value
			} else if key == "blue" {
				subset.Blue = value
			} else if key == "red" {
				subset.Red = value
			}
		}
		game[id] = append(game[id], subset)
	}
	return game, nil
}

func getPossibleSubset(game Game, cubes CubesSubset) int {
	result := 0
	for id, css := range game {
		gameIsPossible := true
		for _, c := range css {
			if c.Green > cubes.Green || c.Blue > cubes.Blue || c.Red > cubes.Red {
				gameIsPossible = false
			}
		}
		if gameIsPossible {
			result += id
		}
	}

	return result
}
