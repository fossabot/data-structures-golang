package main

import (
	"fmt"

	"github.com/rwxramon/data-structures-golang/queue"
	"github.com/rwxramon/data-structures-golang/stack"
)

func main() {
	fmt.Println("Stack ===========")
	myStack := stack.Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

	fmt.Println("Queue ===========")
	myQueue := queue.Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
