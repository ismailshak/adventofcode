// Advent of Code Day 6: Lanternfish
package day06

import (
	"aoc/util"
	"fmt"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Lanternfish")
	fmt.Println("-------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineFishCount(80, inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineFishCount(256, inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)
}

// counting the number of fish that pop out a baby on a given day (the day represented by array index)
// note: clamped the array at index 8, because I'll be rotating it as the days go by
func initFishList(input []string) [9]int {
	var fishList [9]int
	for _, fish := range input {
		fishList[util.ParseInt(fish)] += 1
	}
	return fishList
}

func sum(list [9]int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
}

func determineFishCount(noOfDays int, inputFileName string) int {
	input := util.ReadFile(6, inputFileName, ",")

	// sort of pivoting the data
	fishList := initFishList(input)

	for day := 0; day < noOfDays; day++ {
		// basically, if we "cycle" 0 -> 8, using [index mod 9] we can keep incrementing
		// the target day indirectly (knowing generation span) and store the bare minimum amount
		// of data to sum the total fish count
		fishList[(day+7)%9] += fishList[day%9]
	}

	return sum(fishList)
}
