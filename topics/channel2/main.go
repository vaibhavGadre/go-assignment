package main

import (
	"fmt"
	"time"
)

/* func main() {
	// buffered channel will allow values to be sent without receiver until its full
	myCh := make(chan int, 2) // for first 2 channel values no error

	myCh <- 5
	fmt.Println(<-myCh)
} */

func main() {
	myCh := make(chan string)
	go count("Hello", myCh)

	for v := range myCh {
		fmt.Println(v)
	}
}

func count(msg string, c chan string) {
	for i := 0; i < 6; i++ {
		c <- msg
		time.Sleep(400 * time.Millisecond)
	}
	close(c)
}
