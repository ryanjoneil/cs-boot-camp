Chapter 5. Binary Trees
=======================

Discussion
----------
Binary trees are similar to Linked Lists in that they are composed of nodes,
each of which contains references to other nodes. Instead of a head node,
these start with a root node. Each node can have up to two child nodes (left
and right) and may have a reference to its parent.

Nodes without children are known as leaf nodes. In the tree pictured below,
the nodes containing values 1, 3, and 9 are leaf nodes. The root node contains
the value 5.

{lang=text,linenos=off}
~~~~~~~
	      5
	     / \
	    4   7
	   /   / \
	  2   6   8
	 / \       \
	1   3       9
~~~~~~~

A binary tree is balanced if there is no more than 1 level difference between
any of its leaf nodes, where the level is the number of hops from the root.
The tree shown above is an example of a balanced tree, while the one below
is not.

{lang=text,linenos=off}
~~~~~~~
	      5
	     / \
	    4   7
	   /   
	  2   
	 / \  
	1   3 
~~~~~~~

In these exercises, we shall concern ourselves with binary search trees. These
are binary trees where all elements to the left of a node have value less than
the node value, and all elements to the right of a node have value greater. We
are not going to concern ourselves with balanced trees at this time.

To add a value to an existing tree, we simply start at the root and follow down
left and right node references depending on their value until we find a leaf.
We then add a new leaf to the tree.

For instance, say we start with the following tree.

{lang=text,linenos=off}
~~~~~~~
	  5
	 / \
	1   8
~~~~~~~

If we add 6 to it, we now have this structure.

{lang=text,linenos=off}
~~~~~~~
	  5
	 / \
	1   8
	   /
	  6
~~~~~~~

Adding a node with value 7 then another one with value 8 yields:

{lang=text,linenos=off}
~~~~~~~
	  5
	 / \
	1   8
	   / \
	  6   9
	   \
	    7
~~~~~~~

To search this tree for the value 7, we follow the path 5, 8, 6, 7. To search
for the value 10, we follow the path 5, 8, 9. As 9 is a leaf node and does not
have the value we are searching for, we conclude that the tree does not contain
a node with value 10.


Exercises
---------

### Exercise 0
Create an implementation of a Binary Tree with the following structure:

{lang=text,linenos=off}
~~~~~~~
	class BinaryTreeNode:
		attributes:
			left BinaryTreeNode
			right BinaryTreeNode
			parent BinaryTreeNode
			value int

	class BinaryTree:
		methods:
			insert(value int)
			search(index int) returns BinaryTreeNode
			traverse(func(int))
		
		attributes:
			root BinaryTreeNode
~~~~~~~

The insert function should place a new leaf node in the tree such that it is
less than everything it is to the left of and greater than everything it is to
to the right of. Searching should either return a node instance or null if the
given value cannot be found.

Traversing the tree amounts to running a given function in a depth-first order.
Starting at the root node, do the following:

1. Traverse the left side of the current node.
1. Run the given function on the current node.
1. Traverse the right side of the node.

This should amount to running the function on the nodes of the tree in order.


### Exercise 1
Instantiate a new Binary Tree. Insert all even numbers from 0 to 100, 
[0, 2, 4, ...], into the tree in random order.


### Exercise 2
Search the tree for every number from 0 to 100. Ensure that all even values
are found in the tree and all odd values are not.


### Exercise 3
Use the `traverse(...)` method to print out the values of the tree in order.


Discussion: Insertion Order
---------------------------
When creating a tree, insertion order affects its final structure. The 
following two trees are functionally equivalent, despite having different
structure.

{lang=text,linenos=off}
~~~~~~~
	  5           6
	 / \         / \
	1   8       5   8
	   /       /
	  6       1 
~~~~~~~


Follow-Up Questions
-------------------
- What structure would the tree have if we inserted numbers in order?
- What other data structure would that resemble?
- What are the costs and benefits of implementing `traverse(...)` recursively?


Solutions
---------
<<[Binary Trees: ch05.go](../csbc-bin/ch05.go)
