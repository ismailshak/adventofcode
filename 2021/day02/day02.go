// Advent of Code Day 2: Dive!
package day02

import (
	"aoc/util"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Dive!")
	fmt.Println("-------------")

	// Part 1
	p1Start := time.Now()
	p1Result := determinePositionProduct(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineAimedPositionProduct(inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

func parseMoveDirection(moveValue string) (horizontalScale int, verticalScale int) {
	moveParts := strings.Split(moveValue, " ")

	switch moveParts[0] {
	case "forward":
		horizontalScale = util.ParseInt(moveParts[1])
	case "up":
		verticalScale = util.ParseInt(moveParts[1]) * -1 // depth decreases if we move up
	case "down":
		verticalScale = util.ParseInt(moveParts[1])
	default:
		panic("Error. Move direction didn't match enum")
	}
	return
}

func determinePositionProduct(fileName string) int {
	inputFile := util.OpenInputFile(2, fileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	var horizontal, depth int
	for input.Scan() {
		parsedH, parsedV := parseMoveDirection(input.Text())
		horizontal += parsedH
		depth += parsedV
	}

	return horizontal * depth
}

// Reusing the same old interpretation of the submarine commands, I just wrapped the logic a little
func parseAimedMoveDirection(moveValue string, currentAim int) (int, int) {
	hScale, vScale := parseMoveDirection(moveValue)
	return hScale, currentAim + vScale
}

func determineAimedPositionProduct(inputFileName string) int {
	inputFile := util.OpenInputFile(2, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	var horizontal, depth, aim int
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
