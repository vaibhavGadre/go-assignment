package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	arr := []string{}
	wg.Add(2)
	go print("Hello", &arr)
	go print("World", &arr)
	wg.Wait()
	fmt.Println(arr)
}

func print(str string, arr *[]string) {
	defer wg.Done()
	for i := 0; i < 6; i++ {
		fmt.Println(str)
		mut.Lock()
		*arr = append(*arr, fmt.Sprint(str, i))
		mut.Unlock()
	}	
}
