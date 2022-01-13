package main

import "fmt"

// 1,1,2,3,5,8,13
func main() {
	n := 7

	fmt.Printf("%vth number in fibonacci series: %v\n", n, fibSequence(n))
}

func fibSequence(n int) int {
	if n <= 1 {
		return n
	}

	return fibSequence(n-1) + fibSequence(n-2)
}
