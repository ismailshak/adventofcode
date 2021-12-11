package main

import (
	"aoc/day01"
	"aoc/day02"
	"aoc/day03"
	"aoc/day04"
	"aoc/day05"
	"aoc/day06"
	"aoc/day07"
	"aoc/day08"
	"aoc/day09"
	"aoc/day10"
	"aoc/day11"
	"fmt"
)

func Run(puzzleNumber int, inputFileName string) {
	switch puzzleNumber {
	default:
		fmt.Println("There are 25 puzzles. (1-25)")
	case 1:
		day01.Solve(inputFileName)
	case 2:
		day02.Solve(inputFileName)
	case 3:
		day03.Solve(inputFileName)
	case 4:
		day04.Solve(inputFileName)
	case 5:
		day05.Solve(inputFileName)
	case 6:
		day06.Solve(inputFileName)
	case 7:
		day07.Solve(inputFileName)
	case 8:
		day08.Solve(inputFileName)
	case 9:
		day09.Solve(inputFileName)
	case 10:
		day10.Solve(inputFileName)
	case 11:
		day11.Solve(inputFileName)
	case 12:
		fmt.Println("No puzzle for this day yet")
	case 13:
		fmt.Println("No puzzle for this day yet")
	case 14:
		fmt.Println("No puzzle for this day yet")
	case 15:
		fmt.Println("No puzzle for this day yet")
	case 16:
		fmt.Println("No puzzle for this day yet")
	case 17:
		fmt.Println("No puzzle for this day yet")
	case 18:
		fmt.Println("No puzzle for this day yet")
	case 19:
		fmt.Println("No puzzle for this day yet")
	case 20:
		fmt.Println("No puzzle for this day yet")
	case 21:
		fmt.Println("No puzzle for this day yet")
	case 22:
		fmt.Println("No puzzle for this day yet")
	case 23:
		fmt.Println("No puzzle for this day yet")
	case 24:
		fmt.Println("No puzzle for this day yet")
	case 25:
		fmt.Println("No puzzle for this day yet")
	}
}
