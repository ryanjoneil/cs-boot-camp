Techniques 01: Recursion & Iteration
====================================

Discussion
----------
There are two types of control flow structures available to us for expressing
loops: recursion and iteration. While the circumstances that might lead one to
choose one or the other differ, they are functionally equivalent. Any loop
expressed in one can also be expressed in the other.

Consider the following function that computes the sum of a list of numbers.

	function sum_iter(list of numbers):
		sum = 0
		for number in list of numbers:
			sum = sum + number
		return sum

If we call this function on the list [1, 2, 3, 4], we should expect to see the 
following execution profile, roughly.

	sum_iter([1, 2, 3, 4]):
	|-- sum = 0
	|-- sum = 0 + 1 = 1
	|-- sum = 1 + 2 = 3
	|-- sum = 3 + 3 = 6
	|-- sum = 6 + 4 = 10
	|-- return 10

As expected, the time complexity of this algorithm is O(n). It is also notable
that the spatial complexity is O(1) -- the function runs in a constant amount
of memory.

Assume we have two functions, head(...) which returns the first element of a 
list, and tail(...) which returns the remainder of the list. One way we could
write the same function recursively is as follows.

	function sum_recur(list of numbers):
		if list of numbers is empty:
			return 0
		else:
			return head(list of numbers) + sum_recur(tail(list of numbers))

This is a common way of expressing recursive functions: begin with a condition
that will terminate the looping behavior, and recurse if it is not met. Calling
this function with the same list produces this profile.

	sum_recur([1, 2, 3, 4]):
	|-- return 1 + sum_recur([2, 3, 4])
	|-- |-- return 2 + sum_recur([3, 4])
	|-- |-- |-- return 3 + sum_recur([4])
	|-- |-- |-- |-- return 4 + sum_recur([])
	|-- |-- |-- |-- |-- return 0
	|-- |-- |-- |-- return 4 + 0 = 4
	|-- |-- |-- return 3 + 4 = 7
	|-- |-- return 2 + 7 = 9
	|-- return 1 + 9 = 10

This also runs in O(n) time, but each recursive call builds up a deferred
operation that runs only once the recursion is complete. These deferred 
operations must be stored in memory, typically on a stack. In the case of
this function, the storage requirements are also O(n), so it is actually
less efficient memory-wise than the iterative version.

Recursive functions that do not build up deferred operations on successive
calls are known as tail recursive. These have the desirable feature of running
in constant memory. We can rewrite our recursive summation function above to
be tail recursive by using an accumulator variable, essentially the same as 
the sum variable in our iterative version, and an extra function call.

	function sum_tail_recur(list of numbers):
		function sum_inner(sum, numbers):
			if numbers is empty:
				return sum
			else:
				return sum_inner(sum + head(numbers), tail(numbers))
		return sum_inner(0, list of numbers)

This one executes in linear time and constant space, like the iterative one.

	sum_tail_recur([1, 2, 3, 4]):
	|-- return sum_inner(0, [1, 2, 3, 4])
	|-- return sum_inner(1, [2, 3, 4])
	|-- return sum_inner(3, [3, 4])
	|-- return sum_inner(6, [4])
	|-- return sum_inner(10, [])
	|-- return 10


Exercise 1
----------
The factorial function, usually denoted n!, is defined as follows.

	n! = 1 if n < 2
	n! = n * (n-1)! otherwise

Write recursive, iterative, and tail recursive versions of this function.
Compute the value of 10!.


Exercise 2
----------
The Fibonacci sequence is defined as follows.

	fib(0) = 0
	fib(1) = 1
	fib(n) = fib(n-1) + fib(n-2) if n > 1

This gives the sequence 0, 1, 1, 2, 3, 5, 8, etc.

Write recursive, iterative, and tail recursive versions of this function.
Compute the value of fib(25).


Discussion: List Processing
---------------------------
Two of the first high level programming languages which emerged in the 1950s
were FORTRAN at IBM and LISP at MIT. These were condensed from the names 
Formula Translating System and List Processing, respectively. In FORTRAN, due
to its mathematical nature, list data were typically stored in array 
structures. Lists in LISP, on the other, were stored as linked lists.

In LISP, the head and tail functions are known as car and cdr. In LISP's
prefix notation, these work as follows.

	(car [1, 2, 3]) = 1
	(cdr [1, 2, 3]) = [2, 3]
	(cdr [1])       = []

Most list processing in LISP involved removing the head of a list, processing
it, and recursing over the tail of the list. This pattern can be seen in 
modern functional languages like Erlang or Haskell.

These functions could be chained using successive application of a and d in
their invocations. Thus, cadr would return the head of the tail of a list.

	(cadr [1, 2, 3]) = (car (cdr [1, 2, 3]))
	                 = (car [2, 3])
	                 = 2


Follow-Up Questions
-------------------
- How would one implement head(...) and tail(...)?
- Is it better to use a Linked or an Array List with head(...) and tail(...)?
- What are the time and space complexities of the three factorial functions?
- What are the time and space complexities of the three Fibonacci functions?
