// AdventOfCode Day 1: Sonar Sweep.
// Choosing to stream the data in, versus swallowing the whole input in-memory.
// Choosing to exercise approaches I don't usually take in a professional setting, lol
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("\nSolving Sonar Sweep")
	fmt.Println("---------------------")

	// Part 1
	p1Start := time.Now()
	p1Result := determineDepth()
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineCycledDepth(4)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)

	// Part 1++
	// Trying out solution for part 2 against part 1 prompt
	// - outcome; definitely slower. p1 ~140µs and p1++ ~170µs (p2 ~170µs)
	// - (ran on go v1.17, macos m1 pro)
	p3Start := time.Now()
	p3Result := determineCycledDepth(2)
	p3Duration := time.Since(p3Start)
	fmt.Println("---------------------")
	fmt.Printf("Part 1++: %v (%v)\n", p3Result, p3Duration)
}

type DepthNode struct {
	depth int64
	next  *DepthNode
}

func (node *DepthNode) isGreater(inNode *DepthNode) bool {
	return node.depth > inNode.depth
}

type LinkedList struct {
	length int8
	head   *DepthNode
	tail   *DepthNode
}

// a+b+c < b+c+d == a < d. (more about this below)
func (list *LinkedList) isNewDepthGreater() bool {
	return list.tail.isGreater(list.head)
}

func (list *LinkedList) addNode(value int64) {
	depthNode := DepthNode{
		depth: value,
	}

	if list.length == 0 {
		list.head = &depthNode
		list.tail = &depthNode
	} else {
		list.tail.next = &depthNode
		list.tail = &depthNode
	}

	list.length++
}

func (list *LinkedList) removeHead() {
	if list.head == nil || list.length == 0 {
		panic("Error. Tried to remove head from an empty linked list")
	}

	newHead := list.head.next
	list.head = newHead
	list.length--
}

func openInputFile() *os.File {
	inputFile, err := os.Open("./input.txt")
	if err != nil {
		panic("Error. Failed to read input file.")
	}
	return inputFile
}

// Why does .ParseInt always return int64? :thinking:
func parseInt(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 32)
}

/*
	If we're trying to determine if a+b+c < b+c+d then all we need to check for
	is if a < d (since b and c cancel out there). Using a linked list because it
	fits the whole "i only care about the first and last elements inserted" approach
	I'm going for. If I make a linked list and keep it at length 4, I can cycle
	the "three-measured" depths in and out of the list. So the above structs 'n methods
	are for that. Also, linked list so that I don't have to keep moving/shifting elements to
	the front of a queue/array/slice.
*/
func determineCycledDepth(cycleLimit int8) int32 {
	var numberOfIncreases int32

	inputFile := openInputFile()
	defer inputFile.Close()

	depthList := LinkedList{}
	input := bufio.NewScanner(inputFile)

	for input.Scan() {
		depthValue, _ := parseInt(input.Text())
		depthList.addNode(depthValue)
		if depthList.length == cycleLimit {
			if depthList.isNewDepthGreater() {
				numberOfIncreases++
			}
			depthList.removeHead()
		}
	}

	return numberOfIncreases
}

// original soltution to part 1
func determineDepth() int32 {
	var numberOfIncreases int32

	inputFile := openInputFile()
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	// Scanning once to skip over the first token/input
	input.Scan()
	previous, _ := parseInt(input.Text())

	for input.Scan() {
		current, _ := parseInt(input.Text())
		if current > previous {
			numberOfIncreases++
		}
		previous = current
	}

	return numberOfIncreases
}
