package main

import (
	"fmt"
)

func main() {
	x := []int{19, 2, 7, 19, 3, 19, 1}
	// sort.Sort(sort.Reverse(sort.IntSlice(x)))

	second := -890809098
	largest := -890808098

	for _, val := range x {
		if val > largest {
			second = largest
			largest = val
		} else if val != largest && val > second {
			second = val
		}
	}

	if second == 890809098 {
		fmt.Println("There is no second highest number in an integer array.")
	}
	
	fmt.Println("Second highest number in an integer array:", second)
	fmt.Println("Top two maximum number in an array:", largest, second)
	
}
