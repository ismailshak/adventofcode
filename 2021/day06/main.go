// Advent of Code Day 6: Lanternfish
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("\nSolving Lanternfish")
	fmt.Println("-------------------")
	fmt.Println()

	// Part 1
	p1Start := time.Now()
	p1Result := determineFishCount(80)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineFishCount(256)
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
		panic("Error. Failed to parse int")
	}
	return int(output)
}

// counting the number of fish that pop out a baby on a given day (the day represented by array index)
// note: clamped the array at index 8, because I'll be rotating it as the days go by
func initFishList(input []string) [9]int {
	var fishList [9]int
	for _, fish := range input {
		fishList[parseInt(fish)] += 1
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

func determineFishCount(noOfDays int) int {
	input := readFile()

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
