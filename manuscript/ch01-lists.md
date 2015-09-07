Chapter 1: Lists
================

In this section we'll explore list structures. There are two fundamental types, a Linked List and an Array List. They differ in how they allocate and reference memory needed for storing data. An Array List stores all its data contiguously, such as:

{lang=text,linenos=off}
~~~~~~~
	| 0 | 1 | 2 | 3 |
~~~~~~~

While a Linked List stores each data item separately in memory. Each data item is part of a LinkedListNode which links to the item before and after it. The list has a Head and a Tail, which are the first and last nodes in the list, respectively.

{lang=text,linenos=off}
~~~~~~~
	| 0 | <-> | 1 | <-> | 2 | <-> | 3 |
~~~~~~~


Exercises
---------

X> ### Exercise 0
X> 
X> Create an implementation of a Linked List. It should have the following
X> structure:
X> 
X> {lang=text,linenos=off}
X> ~~~~~~~
X> class LinkedListNode:
X>     attributes:
X>         next LinkedListNode
X>         previous LinkedListNode
X>         value int
X> 
X> class LinkedList:
X>     methods:
X>         append(value int)
X>         get(index int) returns int
X>         insert(index int, value int)
X>         remove(index int)
X>         length() returns int
X>         traverse(func(int))
X> 
X>     attributes:
X>         head LinkedListNode
X>         tail LinkedListNode
X>         len int
X> ~~~~~~~
X> 
X> If you have access to an underlying array type, implement an Array List with
X> the same methods, but use an array for storing its data. Give the Array List
X> a growth factor which determines how much larger it becomes when a call to 
X> `append(...)` overflows its bounds.


X> ### Exercise 1
X> 
X> Insert the first 10 integers, 0 to 9, into the a list. Traverse the list in
X> order and print them to standard output.

X> ### Exercise 2
X> 
X> For n = 100, 1000, ..., 100k, time the following functions against your 
X> Linked List and for an Array or Array List. If you have an Array List 
X> implementation, try this against a variety of growth factors.
X> 
X> a. Append n integers
X> b. Get a random index 10000 times.
X> c. Remove all elements from the front of the list.


D> ## Time Complexity
D> In evaluating the speed of an algorithm, we are often primarily concerned
D> with its time complexity. This is a rough measure of the number of 
D> fundamental operations a computer must perform, given input of a certain 
D> size or length. Fundamental operations are anything that takes a finite 
D> amount of time.
D> 
D> - Reading or writing to memory.
D> - Reading or writing to disk.
D> - Performing a computation.
D> - Sending a packet of data over a network.
D> - Etc.
D> 
D> While these are obviously not all equal in their time impact, we will think 
D> of them as though they are. Consider the following computation.
D> 
D> 
D> {lang=text,linenos=off}
D> ~~~~~~~
D>     x = (a + 3) * 4
D> ~~~~~~~
D> 
D> This could consist of the following operations on a hypothetical machine.
D> 
D> 1. Load the value of the variable a from memory into CPU register (1).
D> 1. Put the value 3 into another CPU register (2).
D> 1. Add the two registers (1) and (2) and store the result in (1).
D> 1. Put the value 4 into CPU register (2).
D> 1. Multiple (1) and (2) and store the result in (1).
D> 1. Copy the value of CPU register (1) into the variable x in memory.
D> 
D> This very simple algorithm thus requires 6 operations, regardless of the
D> particular value of a. We thus say that the algorithm runs in constant time,
D> denoted `O(1)`.
D> 
D> Another algorithm might consist of summing the elements of a list.
D> 
D> {lang=text,linenos=off}
D> ~~~~~~~
D>     sum = 0
D>     for item in some list:
D>         sum = sum + item
D> ~~~~~~~
D> 
D> The first statement takes `O(1)` time, as does the nested addition 
D> statement The loop is executed n times, so it takes `O(1 * n) = O(n)`
D> time. The whole algorithm then takes `O(1 + n)`. Since `n` grows faster
D> than `1`, we shorten this to `O(n)`, which is called linear time.
D> 
D> The important point here is that we don't know, a priori, the size of our
D> input list. The algorithm will take different amounts of time to sum lists
D> of different lengths. As our input lists get longer, knowing the time 
D> complexity of the algorithm gives us an estimate of how much longer it will
D> take to run.
D> 
D> Time complexities are most commonly described using the worst case that an 
D> algorithm can perform. Consider the case of searching an unordered list to
D> see if it contains some element.
D> 
D> {lang=text,linenos=off}
D> ~~~~~~~
D>     contains item = false
D>         for item in some list:
D>             if item is what we are looking for:
D>                 contains item = true
D>                 break
D> ~~~~~~~
D> 
D> If the item we are looking for is the first element of some list, then this
D> algorithm will run in constant time. If it is the last element, then it
D> runs in linear time. This is what the `O(...)` notation indicates. In this 
D> case the algorithm runs in `O(n)` time because that is its worst case
D> behavior.
D> 
D> Let's consider one final algorithm containing a nested loop.
D> 
D> {lang=text,linenos=off}
D> ~~~~~~~
D>     for i in some list:
D>         for j in some list:
D>             print i, j
D> ~~~~~~~
D> 
D> The final line runs in `O(1)` time. It is inside a for loop running `O(n)`
D> time, which is itself inside another for loop. Thus the time complexity is
D> `O(1 * n * n) = O(n^2)`, which is known as quadratic time.


Follow-Up Questions
-------------------
Q> What are the time complexities of the various list method for a Linked List?
Q> 
Q> What are the time complexities of the various list method for an Array List?
Q> 
Q> Under what circumstances might one use one or the other list type?
Q> 
Q> Why is removing the first element of a Linked List faster than an Array List?
Q> 
Q> How does growth factor affect the speed of appending to an Array List?
Q> 
Q> What is the disadvantage of using a high growth factor?


Solutions
---------
<<[Lists: Main Method](../csbc-bin/DataStructures_01_Lists.go)
<<[Lists: List Interface](../csbc-bin/DataStructures_01_Lists.go)
<<[Lists: Array List](../csbc-bin/DataStructures_01_Lists.go)
<<[Lists: Linked List](../csbc-bin/DataStructures_01_Lists.go)
