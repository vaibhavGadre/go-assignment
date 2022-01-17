package main

import (
	"fmt"
	"strings"
)

func main() {
	var a1 string
	var a2 string

	fmt.Println("Enter first array with , separator")
	fmt.Scanln(&a1)

	fmt.Println("Enter second array with , separator")
	fmt.Scanln(&a2)

	m1 := map[string]int{}

	result := []string{}

	for _, val := range strings.Split(a2, ",") {
		if m1[val] == 0 {
			m1[val] = 1
		} else {
			m1[val]++
		}
	}

	for _, val := range strings.Split(a1, ",") {
		if v, ok := m1[val]; ok && v > 0 {
			m1[val]--
		} else {
			result = append(result, val)
		}
	}

	fmt.Println("Following numbers from first array not available in 2nd array:", result)

}
