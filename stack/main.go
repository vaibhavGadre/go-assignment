package main

import "fmt"

type stack []int

// push, pop, peek, isEmpty, length, printStack
func main() {
	var s1 stack
	s1 = s1.Push(4)
	s1 = s1.Push(8)
	s1 = s1.Push(3)
	s1 = s1.Push(7)
	s1 = s1.Push(25)

	fmt.Println("stack:", s1)

	s1 = s1.Pop()
	// s1 = s1.Pop()
	fmt.Println("stack top element:", s1.Seek())
	fmt.Println("Length of stack:", s1.Length())
	fmt.Println("IsEmpty:", s1.IsEmpty())
	fmt.Println("stack:", s1)
}

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() stack {
	if len(s) == 0 {
		panic("The stack is empty")
	}
	fmt.Println("Pop element:", s[len(s)-1])
	return s[:len(s)-1]
}

func (s stack) Seek() int {
	return s[len(s)-1]
}

func (s stack) Length() int {
	return len(s)
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}
