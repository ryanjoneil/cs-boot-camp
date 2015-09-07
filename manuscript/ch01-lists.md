Chapter 1. Lists
================

Array Lists & Linked Lists
--------------------------

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

### Exercise 0

Create an implementation of a Linked List. It should have the following
structure:

{lang=text,linenos=off}
~~~~~~~
class LinkedListNode:
    attributes:
        next LinkedListNode
        previous LinkedListNode
        value int

class LinkedList:
    methods:
        append(value int)
        get(index int) returns int
        insert(index int, value int)
        remove(index int)
        length() returns int
        traverse(func(int))

    attributes:
        head LinkedListNode
        tail LinkedListNode
        len int
~~~~~~~

If you have access to an underlying array type, implement an Array List with
the same methods, but use an array for storing its data. Give the Array List
a growth factor which determines how much larger it becomes when a call to 
`append(...)` overflows its bounds.


### Exercise 1

Insert the first 10 integers, 0 to 9, into the a list. Traverse the list in
order and print them to standard output.


### Exercise 2

For n = 100, 1000, ..., 100k, time the following functions against your 
Linked List and for an Array or Array List. If you have an Array List 
implementation, try this against a variety of growth factors.

a. Append n integers
b. Get a random index 10000 times.
c. Remove all elements from the front of the list.


D> ## Time Complexity
In evaluating the speed of an algorithm, we are often primarily concerned
with its time complexity. This is a rough measure of the number of 
fundamental operations a computer must perform, given input of a certain 
size or length. Fundamental operations are anything that takes a finite 
amount of time.

- Reading or writing to memory.
- Reading or writing to disk.
- Performing a computation.
- Sending a packet of data over a network.
- Etc.

While these are obviously not all equal in their time impact, we will think 
of them as though they are. Consider the following computation.

{lang=text,linenos=off}
~~~~~~~
    x = (a + 3) * 4
~~~~~~~

This could consist of the following operations on a hypothetical machine.

1. Load the value of the variable a from memory into CPU register (1).
1. Put the value 3 into another CPU register (2).
1. Add the two registers (1) and (2) and store the result in (1).
1. Put the value 4 into CPU register (2).
1. Multiple (1) and (2) and store the result in (1).
1. Copy the value of CPU register (1) into the variable x in memory.

This very simple algorithm thus requires 6 operations, regardless of the
particular value of a. We thus say that the algorithm runs in constant time,
denoted `O(1)`.

Another algorithm might consist of summing the elements of a list.

{lang=text,linenos=off}
~~~~~~~
    sum = 0
    for item in some list:
        sum = sum + item
~~~~~~~

The first statement takes `O(1)` time, as does the nested addition 
statement The loop is executed n times, so it takes `O(1 * n) = O(n)`
time. The whole algorithm then takes `O(1 + n)`. Since `n` grows faster
than `1`, we shorten this to `O(n)`, which is called linear time.

The important point here is that we don't know, a priori, the size of our
input list. The algorithm will take different amounts of time to sum lists
of different lengths. As our input lists get longer, knowing the time 
complexity of the algorithm gives us an estimate of how much longer it will
take to run.

Time complexities are most commonly described using the worst case that an 
algorithm can perform. Consider the case of searching an unordered list to
see if it contains some element.

{lang=text,linenos=off}
~~~~~~~
    contains item = false
        for item in some list:
            if item is what we are looking for:
                contains item = true
                break
~~~~~~~

If the item we are looking for is the first element of some list, then this
algorithm will run in constant time. If it is the last element, then it
runs in linear time. This is what the `O(...)` notation indicates. In this 
case the algorithm runs in `O(n)` time because that is its worst case
behavior.

Let's consider one final algorithm containing a nested loop.

{lang=text,linenos=off}
~~~~~~~
    for i in some list:
        for j in some list:
            print i, j
~~~~~~~

The final line runs in `O(1)` time. It is inside a for loop running `O(n)`
time, which is itself inside another for loop. Thus the time complexity is
`O(1 * n * n) = O(n^2)`, which is known as quadratic time.


Follow-Up Questions
-------------------

- What are the time complexities of the various list method for a Linked List?
- What are the time complexities of the various list method for an Array List?
- Under what circumstances might one use one or the other list type?
- Why is removing the first element of a Linked List faster than an Array List?
- How does growth factor affect the speed of appending to an Array List?
- What is the disadvantage of using a high growth factor?


Solutions
---------

Reference code for the list data structures can be found in the appendix.

<<[Lists: ch01.go](../csbc-bin/ch01.go)
