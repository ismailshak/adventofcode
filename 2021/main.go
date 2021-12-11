package main

import (
	"flag"
)

func main() {
	puzzleNumber := flag.Int("p", 1, "The puzzle to execute. Day 1 is puzzle 1, etc")
	inputFileName := flag.String("f", "input.txt", "File name for puzzle input, placed in the puzzle dir")
	flag.Parse()

	Run(*puzzleNumber, *inputFileName)
}
