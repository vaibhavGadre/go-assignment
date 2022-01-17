package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	var input string
    fmt.Println("Enter the input:")
	fmt.Scan(&input) 
	
	// ops := [6]string{"+", "-", "*", "/", "%", "^"}

	var postfix string
	var stack []string

    for _, val := range strings.Split(input, "") {
		if isAlphaNum(val) {
			postfix += val
		} else if val == "(" {
          stack = append(stack, val)
		} else if val == "^" {
			stack = append(stack, val)
		} else if val == ")" {
			for len(stack) != 0 && stack[len(stack) - 1] != "("{
				postfix += stack[len(stack) - 1] //store and pop until ( has found
				stack = stack[:len(stack) - 1]
			}
			stack = stack[:len(stack) - 1] //remove the '(' from stack
		} else {
           if preced(val) > preced(stack[len(stack) - 1]) {
			stack = append(stack, val) //push if precedence is high
		   } else {
			   for len(stack) !=0 && preced(val) <= preced(stack[len(stack) - 1]) {
				postfix += stack[len(stack) - 1]  //store and pop until higher precedence is found
				stack = stack[:len(stack) - 1]
			   }
			   stack = append(stack, val) 
		   }
		}
	}

	for len(stack) != 0{
		postfix += stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
	}
    
	fmt.Println("postfix string:", postfix, stack)

}

func preced(x string) int {
	if x == "+" || x == "-" {
		return 1
	} else if x == "*" || x == "/" || x == "%"{
		return 2
	} else if x == "^"{
      return 3
	}else{
		return 0
	}
}
/* func contains(arr [6]string, x string) bool {
	for _, v:= range arr {
		if v == x {
			return true
		}
	}
	return false
} */

func isAlphaNum(x string) bool {
	result, _ :=  regexp.MatchString(`[a-zA-Z0-9]`, x)
	return result
}