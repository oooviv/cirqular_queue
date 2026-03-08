package main

import "fmt"

const size = 4

type CirqularQueue struct {
	array [size]int
	front int
	rear  int
}

func (cq *CirqularQueue) IsEmpty() bool {
	return cq.front == -1
}

func (cq *CirqularQueue) IsFull() bool {
	return (cq.rear+1)%size == cq.front
}

func (cq *CirqularQueue) Enqueue(v int) bool {
	if cq.IsFull() {
		return false
	}
	if cq.front == -1 {
		cq.front = 0
		cq.rear = 0
		cq.array[cq.front] = v
	} else {
		cq.rear = (cq.rear + 1) % size
		cq.array[cq.rear] = v
	}
	return true
}

func (cq *CirqularQueue) Dequeue() (int, bool) {
	if cq.IsEmpty() {
		return 0, false
	}
	value := cq.array[cq.front]
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
	} else {
		cq.front = (cq.front + 1) % size
	}
	return value, true
}

func (cq *CirqularQueue) PeekQueue() [size]int {
	result := [size]int{}
	if cq.IsEmpty() {
		return result
	}
	i := cq.front
	j := 0
	for {
		// v [1 2 3 4]
		//0
		//f _
		// [4 1 2 3]
		//r _
		//0
		value := cq.array[i]
		result[j] = value
		if i == cq.rear {
			break
		}
		i = (i + 1) % size
		j++
	}

	return result
}

func main() {
	cq := CirqularQueue{
		front: -1,
		rear:  -1,
	}

	// Enqueue
	fmt.Println()
	fmt.Println(cq.PeekQueue())
	if cq.Enqueue(0) {
		fmt.Println("Added:", 0)
	}
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
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	// Dequeue
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	// Enqueue
	fmt.Println()
	if cq.Enqueue(1) {
		fmt.Println("Added:", 1)
	}
	if cq.Enqueue(2) {
		fmt.Println("Added:", 2)
	}
	if cq.Enqueue(3) {
		fmt.Println("Added:", 3)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	// Dequeue
	if v, ok := cq.Dequeue(); ok {
		fmt.Println("Deleted:", v)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	// Enqueue
	if cq.Enqueue(3) {
		fmt.Println("Added:", 3)
	}
	if cq.Enqueue(5) {
		fmt.Println("Added:", 5)
	}
	fmt.Println(cq.PeekQueue())
	fmt.Println()

	/*

	[0 0 0 0]
	Added: 0
	Added: 1
	Added: 2
	Added: 3
	[0 1 2 3]

	Deleted: 0
	Deleted: 1
	Deleted: 2
	Deleted: 3
	[0 0 0 0]


	Added: 1
	Added: 2
	Added: 3
	[1 2 3 0]

	Deleted: 1
	[2 3 0 0]

	Added: 3
	Added: 5
	[2 3 3 5]
	*/
}
