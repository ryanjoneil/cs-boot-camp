Chapter 2. Queues & Stacks
==========================

Discussion
----------
There are two special purpose list structures that are typically based on
Linked Lists: Queues & Stacks. These are useful as simple containers that have
set orders for adding and removing items.

Queues add new items to the tails of their lists and remove them from the
beginning. This is known as First-In-First-Out (FIFO) order. In Queueing
Theory, this is referred to as the queue discipline.

Stacks also add items to the tails, but remove them from the tail. Stacks are
essentially Queues with a Last-In-First-Out (LIFO) discipline.


Exercises
---------

### Exercise 0
Use your Linked List implementation to create Queue and Stack containers.
They should have the following methods:

{lang=text,linenos=off}
~~~~~~~
	class Queue:
		methods:
			enqueue(value int)
			dequeue() return int
			empty() returns bool

	class Stack:
		methods:
			push(value int)
			pop() return int
			empty() returns bool
~~~~~~~


### Exercise 1
Insert the first 10 integers, 0 to 9, into the a Queue. Remove them in order
and print them to standard output.


### Exercise 2
Do the same thing with a Stack.


Follow-Up Questions
-------------------
- Why are these data structures typically based on Linked Lists?
- What would be the impact of using an Array List for a Queue?
- What would be the impact of using an Array List for a Stack?

{pagebreak}

Solutions
---------
<<[Lists: ch02.go](../csbc-bin/ch02.go)
