package main

import (
	"fmt"
)

// {}[]()
func main() {
	var input string 
	fmt.Println("Please enter input:")
	fmt.Scan(&input)
	pMap := map[rune]string{'}': "{", ')': "(", ']': "["}

	stack := []string{}

	for _, val := range input {
		if val == '}' || val == ')' || val == ']' {
           if len(stack) >=1 && pMap[val] == stack[len(stack)-1] {
			   stack = stack[:len(stack)-1]
		   } else {
			fmt.Println("not balanced")
			return
		   }
		} else {
			stack = append(stack, string(val))
		}
	}
	if len(stack) == 0 {
		fmt.Println("balanced")	
		return
	}
	fmt.Println("not balanced")
}

