package main

import (
	"fmt"
	"github.com/ryanjoneil/cslib"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("Exercise 1: 10,000 random Linked List searches")
	for length := 10; length <= 1000000; length *= 10 {
		fmt.Printf("List Length % 8d: ", length)

		list := &cslib.LinkedList{}
		for i := 0; i < length; i++ {
			list.Append(2 * i)
		}

		before := time.Now()
		for j := 0; j < 10000; j++ {
			value := rand.Intn(list.Length() * 2)
			node := list.Search(value)
			if node != nil && node.Value != value {
				fmt.Printf("%d != %d\n", node.Value, value)
			}
		}
		millis := int((time.Now().Sub(before)) / time.Millisecond)
		fmt.Printf("% 6d ms\n", millis)
	}

	fmt.Println("\nExercise 2: 10,000 random Array List searches")
	for length := 10; length <= 1000000; length *= 10 {
		fmt.Printf("List Length % 8d: ", length)

		list := &cslib.ArrayList{Growth: 2.0}
		for i := 0; i < length; i++ {
			list.Append(2 * i)
		}

		before := time.Now()
		for j := 0; j < 10000; j++ {
			value := rand.Intn(list.Length() * 2)
			index := list.Search(value)

			// Sanity check.
			if index != -1 && list.Get(index) != value {
				fmt.Printf("list[%d] == %d, != %d\n", index, list.Get(index), value)
			}
		}
		millis := int((time.Now().Sub(before)) / time.Millisecond)
		fmt.Printf("% 6d ms\n", millis)
	}

	fmt.Println("\nExercise 3: 10,000 random Binary Tree searches")
	for length := 10; length <= 1000000; length *= 10 {
		fmt.Printf("Tree Size % 8d: ", length)

		tree := &cslib.BinaryTree{}
		for i := 0; i < length; i++ {
			tree.Insert(rand.Intn(length * 2))
		}

		before := time.Now()
		for j := 0; j < 10000; j++ {
			value := rand.Intn(length * 2)
			node := tree.Search(value)

			// Sanity check.
			if node != nil && node.Value != value {
				fmt.Printf("%d, != %d\n", node.Value, value)
			}
		}
		millis := int((time.Now().Sub(before)) / time.Millisecond)
		fmt.Printf("% 6d ms\n", millis)
	}
}
