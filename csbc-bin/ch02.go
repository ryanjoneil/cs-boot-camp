package main

import (
	"fmt"

	"github.com/ryanjoneil/cs-boot-camp/csbc-lib"
)

func main() {
	// Queue exercise
	fmt.Println("Exercise 1:")
	queue := &cslib.Queue{}
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
	}
	for !queue.Empty() {
		fmt.Printf("%d ", queue.Dequeue())
	}
	fmt.Println()

	// Stack exercise
	fmt.Println("\nExercise 2:")
	stack := &cslib.Stack{}
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for !stack.Empty() {
		fmt.Printf("%d ", stack.Pop())
	}
	fmt.Println()
}
