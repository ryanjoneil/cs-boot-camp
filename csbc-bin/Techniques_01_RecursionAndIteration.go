package main

import "fmt"

// Recursive version of factorial function that builds deferred operations.
func factorial1(n int) int {
	if n < 2 {
		return 1
	} else {
		return n * factorial1(n-1)
	}
}

// Iterative version of factorial that runs in constant space.
func factorial2(n int) int {
	acc := 1
	for ; n > 1; n-- {
		acc *= n
	}
	return acc
}

// Tail recursive version of factorial.
func fact3Inner(acc int, x int) int {
	if x < 2 {
		return acc
	} else {
		return fact3Inner(acc*x, x-1)
	}
}

func factorial3(n int) int {
	return fact3Inner(1, n)
}

// Recursive version of Fibonacci function.
func fibonacci1(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci1(n-1) + fibonacci1(n-2)
	}
}

// Iterative version of the Fibonacci function.
func fibonacci2(n int) int {
	x, y := 0, 1
	for ; n > 0; n-- {
		x, y = y, x+y
	}
	return x
}

// Tail recursive version of the Fibonacci function.
func fib3Inner(x, y, n int) int {
	if n > 0 {
		return fib3Inner(y, x+y, n-1)
	} else {
		return x
	}
}

func fibonacci3(n int) int {
	return fib3Inner(0, 1, n)
}

func main() {
	fmt.Printf("factorial1(10) = %d\n", factorial1(10))
	fmt.Printf("factorial2(10) = %d\n", factorial2(10))
	fmt.Printf("factorial3(10) = %d\n", factorial3(10))

	fmt.Printf("fibonacci1(25) = %d\n", fibonacci1(25))
	fmt.Printf("fibonacci2(25) = %d\n", fibonacci2(25))
	fmt.Printf("fibonacci3(25) = %d\n", fibonacci3(25))
}
