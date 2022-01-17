package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("Enter the input string:")
	fmt.Scan(&str)

	maxLength := 1
	start := 0
	end := 0
	strLen := len(str)
	low := 0
	high := 0
	for i := 1; i < strLen; i++ {
		// for even palindrome  abba
		// two center elements i-1, i
		low = i - 1
		high = i

		for low >= 0 && high < strLen && str[low] == str[high]{
			low--
			high++
		}
		// move back to previous valid palindrome value before end of loop
		low++
		high--
		if str[low] == str[high] && maxLength < high-low+1 {
			maxLength = high - low + 1
			start = low
			end = high
		}

		// for odd palindorme  bob
		// only one center element i

		low = i - 1
		high = i + 1

		for low >=0 && high < strLen && str[low] == str[high]  {
			low--
			high++
		}
     
		low++
		high--
		if str[low] == str[high] && maxLength < high-low+1 {
			maxLength = high - low + 1
			start = low
			end = high
		} 
	}

	fmt.Println("max length:", maxLength, "string:", str[start:end+1], "start:", start, "end:", end)
}
