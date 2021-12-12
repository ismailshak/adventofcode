package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OpenInputFile(puzzleNumber int, name string) *os.File {
	inputFile, err := os.Open(BuildPuzzlePath(puzzleNumber, name))
	if err != nil {
		panic(err)
	}
	return inputFile
}

func ReadFile(puzzleNumber int, name, splitBy string) []string {
	file, err := os.ReadFile(BuildPuzzlePath(puzzleNumber, name))
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), splitBy)
}

func ParseInt(value string) int {
	output, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(output)
}

func BuildPuzzlePath(puzzleNumber int, file string) string {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v/day%02d/%v", rootDir, puzzleNumber, file)
}
