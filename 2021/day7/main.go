// Advent of Code Day 7: The Treachery of Whales
// Bit of statistics in this one...
package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\nSolving The Treachery of Whales")
	fmt.Println("-------------------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineCheapestConstantMove()
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineCheapestLinearMove()
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

func readFile() []string {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), ",")
}

func parseInt(value string) int {
	output, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(output)
}

// part 1: constant fuel consumption
func determineCheapestConstantMove() int {
	input := readFile()
	length := float64(len(input))

	sort.SliceStable(input, func(i, j int) bool {
		return parseInt(input[i]) < parseInt(input[j])
	})

	median := parseInt(input[int(math.Ceil(length/2))])
	totalFuel := 0

	for _, position := range input {
		fuelNeededToRelocate := float64(median - parseInt(position))
		totalFuel += int(math.Abs(fuelNeededToRelocate))
	}

	return totalFuel
}

/*
	arithmetic progression/series because I'm finding the sum of a 1+2+3+4+etc sequence;
				-	SUM = N/2(a+L)

	where N is series length, a is start and L is end
	a will always be 1, and L is always the difference between the crab and target positions
*/
func sumFuelSequence(start, end int) int {
	diff := math.Abs(float64((start - end)))
	return int((diff * 0.5) * (1 + diff))
}

// part 2: linearly increasing fuel consumption
func determineCheapestLinearMove() int {
	input := readFile()
	length := float64(len(input))

	sumOfPositions := 0
	for _, position := range input {
		sumOfPositions += parseInt(position)
	}

	mean := int(math.Floor((float64(sumOfPositions)) / (length)))
	totalFuel := 0
	for _, position := range input {
		newSum := sumFuelSequence(parseInt(position), mean)
		totalFuel += newSum
	}

	return totalFuel
}
