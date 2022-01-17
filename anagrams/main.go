package main

import (
	"fmt"
	"strings"
)




func main()  {
	var ip1 string
	var ip2 string

	fmt.Println("Enter first string:")
	fmt.Scanln(&ip1)
	fmt.Println("Enter second string:")
	fmt.Scanln(&ip2)

	fmt.Println("values:", ip1, ip2)

	if len(ip1) != len(ip2) {
		fmt.Println("Both strings are NOT anagrams")
		return 
	}

	m1 := make(map[string]int, 0)
	
	for _, val := range strings.Split(ip1, "") {
		if m1[val] == 0 {
			m1[val] = 1
		} else {
			m1[val]++
		}
	}
	fmt.Println("Character map:", m1)
	for _, val := range strings.Split(ip2, ""){
		if v, ok := m1[val]; ok && v > 0 {
			m1[val]-- 
		} else {
			fmt.Println("Both strings are NOT anagrams")
			return
		}
	}
	// if !contains(ip2, val){
	//    fmt.Println("Both strings are NOT anagrams")
	//    return 
	// }
	fmt.Println("Both strings are anagrams")
}


// func contains(str string, c string) bool{
// 	for _, val := range strings.Split(str, "") {
// 		if val == c{
//           return true
// 		}
// 	}

// 	return false
// } 