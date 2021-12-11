// Advent of Code Day 3: Binary Diagnostics
package day03

import (
	"aoc/util"
	"bufio"
	"fmt"
	"strconv"
	"time"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Binary Diagnostics")
	fmt.Println("--------------------------")

	// Part 1
	p1Start := time.Now()
	power, gamma, epsilon := determinePower(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1: (%v)\n", p1Duration)
	fmt.Printf("\tPower - %v\n\tGamma - %v\n\tEpsilon - %v\n", power, gamma, epsilon)

	// Part 2
	p2Start := time.Now()
	lifeSupport, oxygen, cO2 := determineLifeSupport(inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2: (%v)\n", p2Duration)
	fmt.Printf("\tLife Support - %v\n\tOxygen - %v\n\tCO2 - %v\n", lifeSupport, oxygen, cO2)
}

func binaryStringToInt(binValue string) int64 {
	intValue, err := strconv.ParseInt(binValue, 2, 32)
	if err != nil {
		panic(err)
	}

	return intValue
}

// janky bitwise NOT (TODO: how to bit manip. in Go tho?)
func flipBits(value string) string {
	var flippedBinary string
	for _, bit := range value {
		if bit == '1' {
			flippedBinary += "0"
		} else {
			flippedBinary += "1"
		}
	}
	return flippedBinary
}

// If we see a '1', increase the frequency at current index by 1
// If we see a '0', decrease the frequency at current index by 1
func accumulateBits(binaryValue string, bitFrequency *[12]int32) {
	for index, char := range binaryValue {
		if char == '1' {
			bitFrequency[index]++
		} else {
			bitFrequency[index]--
		}
	}
}

// If element at index is positive, '1' was the most common bit
// If element at index is negative, '0' was the most common bit
func frequencyToBinaryString(bitFrequency *[12]int32) string {
	var commonBitsBinary string
	for _, freq := range bitFrequency {
		if freq > 0 {
			commonBitsBinary += "1"
		} else {
			commonBitsBinary += "0"
		}
	}

	return commonBitsBinary
}

func determinePower(inputFileName string) (power int64, gamma, epsilon string) {
	inputFile := util.OpenInputFile(3, inputFileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)

	// Stores frequency in cumulative form
	bitFrequency := [12]int32{}

	for input.Scan() {
		binaryValue := input.Text()
		accumulateBits(binaryValue, &bitFrequency)
	}

	gamma = frequencyToBinaryString(&bitFrequency)
	// if we're looking for the opposite of gamma, then I'm gonna flip it
	epsilon = flipBits(gamma)
	power = binaryStringToInt(gamma) * binaryStringToInt(epsilon)
	return
}

// ew. (TODO: come back and clean this up)
func getBitMatch(bitFrequencyValue int32, isOxygen bool) byte {
	if isOxygen {
		if bitFrequencyValue >= 0 {
			return '1'
		} else {
			return '0'
		}
	} else {
		if bitFrequencyValue >= 0 {
			return '0'
		} else {
			return '1'
		}
	}
}

// Will return the final reading that matches the bit criteria
// Would probably look cleaner and perform better if I was better at Go, lol
// (TODO: get comfortable with slice trickery and replace duplicate arrays)
func filterReadings(input []string, isOxygen bool, index int) string {
	if len(input) == 1 {
		return input[0]
	}

	bitFrequency := [12]int32{}

	for _, reading := range input {
		accumulateBits(reading, &bitFrequency)
	}

	bitMatch := getBitMatch(bitFrequency[index], isOxygen)

	var filteredList []string
	for _, reading := range input {
		if reading[index] == bitMatch {
			filteredList = append(filteredList, reading)
		}
	}

	return filterReadings(filteredList, isOxygen, index+1)
}

func determineLifeSupport(inputFileName string) (lifeSupport int64, oxygen, cO2 string) {
	input := util.ReadFile(3, inputFileName, "\n")

	oxygen = filterReadings(input, true, 0)
	cO2 = filterReadings(input, false, 0)
	lifeSupport = binaryStringToInt(oxygen) * binaryStringToInt(cO2)
	return
}
