// Advent of Code Day 7: The Treachery of Whales
// Bit of statistics in this one...
package day07

import (
	"aoc/util"
	"fmt"
	"math"
	"sort"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving The Treachery of Whales")
	fmt.Println("-------------------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineCheapestConstantMove(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineCheapestLinearMove(inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

// part 1: constant fuel consumption
func determineCheapestConstantMove(inputFileName string) int {
	input := util.ReadFile(7, inputFileName, ",")
	length := float64(len(input))

	sort.SliceStable(input, func(i, j int) bool {
		return util.ParseInt(input[i]) < util.ParseInt(input[j])
	})

	median := util.ParseInt(input[int(math.Ceil(length/2))])
	totalFuel := 0

	for _, position := range input {
		fuelNeededToRelocate := float64(median - util.ParseInt(position))
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
func determineCheapestLinearMove(inputFileName string) int {
	input := util.ReadFile(7, inputFileName, ",")
	length := float64(len(input))

	sumOfPositions := 0
	for _, position := range input {
		sumOfPositions += util.ParseInt(position)
	}

	mean := int(math.Floor((float64(sumOfPositions)) / (length)))
	totalFuel := 0
	for _, position := range input {
		totalFuel += sumFuelSequence(util.ParseInt(position), mean)
	}

	return totalFuel
}
