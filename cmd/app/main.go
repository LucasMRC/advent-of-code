package main

import (
	"os"

	"github.com/LucasMRC/advent-of-code/internal/day1"
	"github.com/LucasMRC/advent-of-code/internal/day2"
	"github.com/LucasMRC/advent-of-code/internal/day3"
	"github.com/LucasMRC/advent-of-code/internal/day4"
)

func main() {
	argCount := len(os.Args)

	var command, day string
	if argCount == 1 {
		command = "run"
		day = "all"
	} else {
		command = os.Args[1]
		day = os.Args[2]
	}

	if command == "run" {
		switch day {
		case "1":
			day1.Run()
		case "2":
			day2.Run()
		case "3":
			day3.Run()
		case "4":
			day4.Run()
		case "all":
			day1.Run()
			day2.Run()
			day3.Run()
			day4.Run()
		}
	} else if command == "test" {
		switch day {
		case "1":
			day1.Run()
		case "2":
			day2.Run()
		case "3":
			day3.Run()
		case "4":
			day4.Test()
		case "all":
			day1.Run()
			day2.Run()
			day3.Run()
			day4.Run()
		}

	}
}
