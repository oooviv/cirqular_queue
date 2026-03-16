package main

import "fmt"

const size int = 4

type cirqularQueue struct {
	front int
	rear  int
	array [size]int
}

func (cq *cirqularQueue) isEmpty() bool {
	return cq.front == -1
}

func (cq *cirqularQueue) isFull() bool {
	return (cq.rear+1)&(size-1) == cq.front
}

func (cq *cirqularQueue) Enqueue(v int) bool {
	if cq.isFull() {
		return false
	}
	if cq.front == -1 {
		cq.front = 0
	}
	cq.rear = (cq.rear + 1) & (size - 1)
	cq.array[cq.rear] = v

	return true
}

func (cq *cirqularQueue) Dequeue() (int, bool) {
	value := 0
	if cq.isEmpty() {
		return 0, false
	}
	value = cq.array[cq.front]
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) & (size - 1)
	}
	return value, true
}

func (cq *cirqularQueue) PeekQueue() [size]int {
	array := [size]int{}
	if cq.isEmpty() {
		return array
	}
	i := cq.front
	j := 0
	for {
		array[j] = cq.array[i]
		if i == cq.rear {
			break
		}
		i = (i + 1) & (size - 1)
		j++
	}
	return array
}

func main() {
	cq := cirqularQueue{
		front: -1,
		rear:  -1,
		array: [size]int{},
	}
	_ = cq

	// Enqueue
	if cq.Enqueue(1) {
		fmt.Println("Added:", 1)
	}
	if cq.Enqueue(2) {
		fmt.Println("Added:", 2)
	}
	if cq.Enqueue(3) {
		fmt.Println("Added:", 3)
	}
	if cq.Enqueue(4) {
		fmt.Println("Added:", 4)
	}
	if cq.Enqueue(5) {
		fmt.Println("Added:", 5)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	if j, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", j)
	}
	if j, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", j)
	}
	if j, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", j)
	}
	// Enqueue
	if cq.Enqueue(5) {
		fmt.Println("Added:", 5)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()
}
