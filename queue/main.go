package main

import "fmt"

type queue []int

// enqueue, dequeue, front, isEmpty, length, printStack

func main() {
	var q1 queue
	q1 = q1.Enqueue(4)
	q1 = q1.Enqueue(9)
	q1 = q1.Enqueue(25)
	q1 = q1.Enqueue(65)

	fmt.Println("Queue:", q1)

	q1 = q1.Dequeue()

	fmt.Println("Queue Front element:", q1.Front())
	fmt.Println("Length of stack:", q1.Length())
	fmt.Println("IsEmpty:", q1.IsEmpty())
	fmt.Println("Queue:", q1)
}

func (q queue) Enqueue(v int) queue {
	return append([]int{v}, q...)
}

func (q queue) Dequeue() queue {
	if len(q) == 0 {
		panic("queue is empty")
	}
	fmt.Println("Dequeue element:", q[0])
	return q[1:]
}

func (q queue) Front() int {
	return q[0]
}
func (q queue) Length() int {
	return len(q)
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}
