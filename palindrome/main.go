package main

import (
	"fmt"
)

func main() {
	// x := "oyo"
	var strX string

	fmt.Println("Enter input:")
	fmt.Scan(&strX)
	
	length := len(strX)
	fmt.Printf("value is %v and length is %v\n", strX, length)
	isPalindrome(length, strX)
}


func isPalindrome(length int, strX string) bool {
	mid := length / 2
	i := 0
	for i <= mid {
		if strX[i] != strX[length-1-i] {
			fmt.Println("input is not palindrome")
			return false
		} else {
			i++
		}
	}
	fmt.Println("input is palindrome")
	return true
}
