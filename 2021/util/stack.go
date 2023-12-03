package util

import "fmt"

// generics not yet supported in Go v1.17 (coming in 1.18)
// but the below will work when i upgrade, for now duplicating code for the day 12 usage :(

//import "fmt"

//type Stack[T any] []T

//func (s Stack[T]) Push(v T) Stack[T] {
//	return append(s, v)
//}

//func (s Stack[T]) Pop() (Stack[T], T) {
//	if len(s) == 0 {
//		panic("Can't pop an empty stack")
//	}

//	l := len(s)
//	return s[:l-1], s[l-1]
//}

//func (s Stack[T]) Peek() T {
//	if len(s) == 0 {
//		return ""
//	}

//	l := len(s)
//	return s[l-1]
//}

//const (
//	CYAN  = "\033[36m"
//	RESET = "\033[0m"
//)

//// hard to read the default print because it wraps the output in `[ ]`
//func (s Stack) Print() {
//	fmt.Print("| ")
//	for _, token := range s {
//		fmt.Print(CYAN, token, RESET, ", ")
//	}
//	fmt.Print(" |\n")
//}

type Stack [][]string

func (s Stack) Push(v []string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, []string) {
	if len(s) == 0 {
		panic("Can't pop an empty stack")
	}

	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Peek() []string {
	if len(s) == 0 {
		return nil
	}

	l := len(s)
	return s[l-1]
}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

const (
	CYAN  = "\033[36m"
	RESET = "\033[0m"
)

// hard to read the default print because it wraps the output in `[ ]`
func (s Stack) Print() {
	fmt.Print("| ")
	for _, token := range s {
		fmt.Print(CYAN, token, RESET, ", ")
	}
	fmt.Print(" |\n")
}
