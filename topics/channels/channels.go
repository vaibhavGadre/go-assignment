package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	myCh := make(chan int) // unbuffered channel
	wg.Add(2)

	// send-only channel,  going into box
	go func(myCh chan<- int, wg *sync.WaitGroup) {
		myCh <- 5 // blocking waits until someone ready to receive
		close(myCh)
		wg.Done()
	}(myCh, wg)

	// receive-only channel, going out box
	go func(myCh <-chan int, wg *sync.WaitGroup) {
		val, isOpen := <-myCh
		fmt.Println(val, isOpen)
		val, isOpen = <-myCh
		fmt.Println(val, isOpen)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
