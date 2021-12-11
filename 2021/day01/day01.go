// AdventOfCode Day 1: Sonar Sweep.
// Choosing to stream the data in, versus swallowing the whole input in-memory.
// Choosing to exercise approaches I don't usually take in a professional setting, lol
package day01

import (
	"bufio"
	"fmt"
	"time"

	"aoc/util"
)

func Solve(inputFileName string) {
	fmt.Println("\nSolving Sonar Sweep")
	fmt.Println("---------------------")

	// Part 1
	p1Start := time.Now()
	p1Result := determineDepth(inputFileName)
	p1Duration := time.Since(p1Start)
	fmt.Printf("Part 1 Result: %v (%v)\n", p1Result, p1Duration)

	// Part 2
	p2Start := time.Now()
	p2Result := determineCycledDepth(4, inputFileName)
	p2Duration := time.Since(p2Start)
	fmt.Printf("Part 2 Result: %v (%v)\n", p2Result, p2Duration)

	// Part 1++
	// Trying out solution for part 2 against part 1 prompt
	// - outcome; definitely slower. p1 ~140µs and p1++ ~170µs (p2 ~170µs)
	// - (ran on go v1.17, macos m1 pro)
	p3Start := time.Now()
	p3Result := determineCycledDepth(2, inputFileName)
	p3Duration := time.Since(p3Start)
	fmt.Println("---------------------")
	fmt.Printf("Part 1++: %v (%v)\n", p3Result, p3Duration)
}

type DepthNode struct {
	depth int
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

func (list *LinkedList) addNode(value int) {
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

/*
	If we're trying to determine if a+b+c < b+c+d then all we need to check for
	is if a < d (since b and c cancel out there). Using a linked list because it
	fits the whole "i only care about the first and last elements inserted" approach
	I'm going for. If I make a linked list and keep it at length 4, I can cycle
	the "three-measured" depths in and out of the list. So the above structs 'n methods
	are for that. Also, linked list so that I don't have to keep moving/shifting elements to
	the front of a queue/array/slice.
*/
func determineCycledDepth(cycleLimit int8, fileName string) int32 {
	var numberOfIncreases int32

	inputFile := util.OpenInputFile(1, fileName)
	defer inputFile.Close()

	depthList := LinkedList{}
	input := bufio.NewScanner(inputFile)

	for input.Scan() {
		depthValue := util.ParseInt(input.Text())
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
func determineDepth(fileName string) int32 {
	var numberOfIncreases int32

	inputFile := util.OpenInputFile(1, fileName)
	defer inputFile.Close()

	input := bufio.NewScanner(inputFile)
	// Scanning once to skip over the first token/input
	input.Scan()
	previous := util.ParseInt(input.Text())

	for input.Scan() {
		current := util.ParseInt(input.Text())
		if current > previous {
			numberOfIncreases++
		}
		previous = current
	}

	return numberOfIncreases
}
