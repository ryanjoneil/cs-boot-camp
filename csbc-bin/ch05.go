package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ryanjoneil/cs-boot-camp/csbc-lib"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// 1. Insert all even numbers from 0 to 100 into a tree in random order.
	// Create list of even numbers from 0 to 100.
	numbers := make([]int, 50)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = 2 * i
	}

	// Shuffle that list.
	for i := 0; i < len(numbers); i++ {
		j := rand.Intn(len(numbers))
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	tree := &cslib.BinaryTree{}
	for _, i := range numbers {
		tree.Insert(i)
	}

	// 2. Search for every integer from 0 to 100. Ensure that only
	//    even integers are found within the tree.
	for i := 0; i < 100; i++ {
		node := tree.Search(i)
		if i%2 != 0 {
			if node != nil {
				fmt.Println("Fail: found odd number", i, "in tree")
			}
		} else {
			if node == nil {
				fmt.Println("Fail: could not find", i)
			} else if node.Value != i {
				fmt.Println("Fail:", i, "!=", node.Value)
			}
		}
	}

	// 3. Print out the tree nodes in order using Depth-First Search.
	tree.Traverse(func(x int) { fmt.Print(x, " ") })
	fmt.Println()
}
