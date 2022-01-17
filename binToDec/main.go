package main

import (
	"fmt"
)

// i/p: 1010
// o/p: 10
func main() {
	var x int

	fmt.Println("Enter binary number:")
	fmt.Scan(&x)

	var result int
	var rem int

	i := 0
	m := 1

	for x != 0 {
		rem = x % 10
		x /= 10
		result = result + (rem * m)
		i++
		m = m * 2
	}

	fmt.Println("The decimal value:", result)
}
