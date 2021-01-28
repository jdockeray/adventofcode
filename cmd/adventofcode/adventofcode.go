package main

import (
	"fmt"
	"github.com/jdockeray/adventofcode/internal/day1"
	"github.com/jdockeray/adventofcode/internal/day10"
	"github.com/jdockeray/adventofcode/internal/day2"
	"github.com/jdockeray/adventofcode/internal/day3"
	"github.com/jdockeray/adventofcode/internal/day4"
	"github.com/jdockeray/adventofcode/internal/day5"
	"github.com/jdockeray/adventofcode/internal/day6"
	"github.com/jdockeray/adventofcode/internal/day7"
	"github.com/jdockeray/adventofcode/internal/day8"
	"github.com/jdockeray/adventofcode/internal/day9"
	"os"
)

func main() {
	day := os.Getenv("DAY")

	switch day {
	case "day1":
		day1.Day1()
	case "day2":
		day2.Day2()
	case "day3":
		day3.Day3()
	case "day4":
		day4.Day4()
	case "day5":
		day5.Day5()
	case "day6":
		day6.Day6()
	case "day7":
		day7.Day7()
	case "day8":
		day8.Day8()
	case "day9":
		day9.Day9()
	case "day10":
		day10.Day10()
	default:
		fmt.Printf("unrecognised argument: %s", day)
	}

}
