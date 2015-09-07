Algorithms 01: Searching
========================

Discussion
----------
One of the most fundamental questions we can ask about data is: Does a set 
of items contain some known item or not?

We have [already seen](https://github.com/ryanjoneil/csbin/blob/master/DataStructures_03_BinaryTrees.md) 
an example of how one might answer this 
question against data organized into a binary tree. However, it is much more 
common to ask this about data that is organized in lists. The speed with which 
we can answer this question depends greatly on the list type we are using.

The simplest way to search any list is to start at its beginning, check each
element in order, and stop if particular item is found. If the item is at the
start of the list, we are in luck. The time complexity of our search in this
best case scenario is constant, or O(1). But if the item isn't in the list, we
are forced to look at every before deciding that it isn't found. This worst
case scenario exhibits time complexity that grows linearly with the size of
our list, in O(n).

We have explored [searching lists](https://github.com/ryanjoneil/csbin/blob/master/DataStructures_01_Lists.md) 
generally. Now we examine it based on the type of list. If we implement this
naive search technique on an Array List, our algorithm may look like this.

	contains item = false
	i = first index of the list
	while i <= last index of the list:
		if index i of the list is what we are looking for:
			contains item = true
			break

If we are using a Linked List, then we start at the head and check every
node until we reach the end.

	contains item = false
	current node = list head
	while current node is not null:
		if current node contains what we are looking for:
			contains item = true
			break
		else:
			current node = next node

Despite their different implementations, both of these algorithms have the same
time complexity. 

If we are using an Array List and our data is sorted, we can do much better.
Instead of checking every node in order, we can start in the middle. If the
middle node is greater than the item we are looking for, we examine the middle
of the list to its left. If it is less than the item we are looking for, we
examine the middle of the list to its right. We repeat this procedure with
sublists of smaller and smaller size until we either find the item or determine
that it is not contained in our list. This approach is known as divide and
conquer. 

For instance, say we are searching for the number 8 in the following list.

	| 0 | 1 | 2 | 4 | 5 | 6 | 7 | 8 | 10 |

We start at the middle, which is 5. Since 5 is less than 8, we search the
sublist to its right.

	                    | 6 | 7 | 8 | 10 |

This list doesn't have an exact middle, so we can look at either the second
or third element. Let's say we pick the second, and get 7. 7 is less than 8,
so we now search the sublist to its right.

	                            | 8 | 10 |

Again, we have no middle, so here we end up looking at the first element. This
is what we are looking for, so we're done. Note that if it didn't happen to be
what we were searching for, we would still be done since we can't further
subdivide the list.

Divide and conquer reduces the size of the list we need to search by one half
with each iteration. The maximum number of nodes we have to examine on sorted 
data is the base-2 logarithm of the list length. Using an Array List means
we have O(1)-time retrievals, so our search time in this case is O(log n).


Exercise 1
----------
For n = 10, 100, ..., 1 million, created a sorted Linked List of the first
n even numbers. Time 10,000 random searches against these lists.


Exercise 2
----------
For n = 10, 100, ..., 1 million, created a sorted Array List of the first
n even numbers. Time 10,000 random searches against these lists.


Exercise 3
----------
For n = 10, 100, ..., 1 million, created a Binary Tree and insert n random
integers into it. Time 10,000 random searches against these trees.


Follow-Up Questions
-------------------
* Does it make a difference if the data is sorted for the Linked List?
* What is the time complexity of the Binary Tree search if the tree
  is balanced?
* If data is inserted into a Binary Tree in sorted order, it resembles a
  Linked List is structure. What does this mean for its search time?
