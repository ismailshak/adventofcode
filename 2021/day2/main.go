// Advent of Code Day 2: Dive!
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Solving Dive!")
	start := time.Now()
	fmt.Printf("Part 1 Result: %v\n", determinePositionProduct())
	p1Duration := time.Since(start)
	fmt.Printf("Part 2 Result: %v\n", determineAimedPositionProduct())
	p2Duration := time.Since(start)
	fmt.Println("---------------------")
	fmt.Printf("Part 1 took %v. Part 2 took %v\n", p1Duration, p2Duration)
}

func openInputFile() *os.File {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic("Error. Failed to read input file.")
	}
	return inputFile
}

func parseInt(value string) int64 {
	output, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic("Error. Failed to parse int")
	}
	return output
}

func parseMoveDirection(moveValue string) (horizontalScale int64, verticalScale int64) {
	moveParts := strings.Split(moveValue, " ")

	switch moveParts[0] {
	case "forward":
		horizontalScale = parseInt(moveParts[1])
	case "up":
		verticalScale = parseInt(moveParts[1]) * -1 // depth decreases if we move up
	case "down":
		verticalScale = parseInt(moveParts[1])
	default:
		panic("Error. Move direction didn't match enum")
	}
	return
}

func determinePositionProduct() int64 {
	inputFile := openInputFile()
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	var horizontal, depth int64
	for input.Scan() {
		parsedH, parsedV := parseMoveDirection(input.Text())
		horizontal += parsedH
		depth += parsedV
	}

	return horizontal * depth
}

// Reusing the same old interpretation of the submarine commands, I just wrapped the logic a little
func parseAimedMoveDirection(moveValue string, currentAim int64) (int64, int64) {
	hScale, vScale := parseMoveDirection(moveValue)
	return hScale, currentAim + vScale
}

func determineAimedPositionProduct() int64 {
	inputFile := openInputFile()
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	var horizontal, depth, aim int64
	for input.Scan() {
		parsedH, newAim := parseAimedMoveDirection(input.Text(), aim)
		horizontal += parsedH
		aim = newAim
		// only do this when we get a "forward"
		if parsedH > 0 {
			depth += (aim * parsedH)
		}
	}

	return horizontal * depth
}
