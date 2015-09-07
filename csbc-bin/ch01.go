package main

import (
	"fmt"
	"github.com/ryanjoneil/cs-boot-camp/csbc-lib"
	"math"
	"math/rand"
	"time"
)

func timeFunc(what string, f func()) {
	before := time.Now()
	f()
	millis := int((time.Now().Sub(before)) / time.Millisecond)
	fmt.Printf("\t%s:\t% 5d ms\n", what, millis)
}

func testList(listType string, n int, list cslib.List) {
	// a. Append n integers.
	timeFunc(fmt.Sprintf("%s Append", listType), func() {
		for j := 0; j < n; j++ {
			list.Append(j)
		}
	})

	// b. Get a random index 10000 times.
	timeFunc(fmt.Sprintf("%s Get 10K", listType), func() {
		for j := 0; j < 10000; j++ {
			list.Get(rand.Intn(list.Length()))
		}
	})

	// c. Remove all elements from the front of the list.
	timeFunc(fmt.Sprintf("%s Remove", listType), func() {
		for list.Length() > 0 {
			list.Remove(0)
		}
	})

	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var list cslib.List

	// 1. Insert 10 integers and print them out in order.
	fmt.Println("Exercise 1:")

	list = &cslib.LinkedList{}
	for i := 0; i < 10; i++ {
		list.Append(i)
	}
	list.Traverse(func(x int) { fmt.Printf("%d ", x) })
	fmt.Println()

	// The remaining exercises we run for a variety of different list types.

	// 2. For n = 100, ..., X, time the following operations:
	fmt.Println("\nExercise 2:")
	for i := 2; i <= 5; i++ {
		// Size of list.
		n := int(math.Pow(10, float64(i)))
		fmt.Printf("n = %d\n", n)

		list = &cslib.LinkedList{}
		testList("LinkedList", n, list)

		list = &cslib.ArrayList{Growth: 1.1}
		testList("ArrayList(1.1)", n, list)

		list = &cslib.ArrayList{Growth: 1.5}
		testList("ArrayList(1.5)", n, list)

		list = &cslib.ArrayList{Growth: 2.0}
		testList("ArrayList(2.0)", n, list)
	}
}
