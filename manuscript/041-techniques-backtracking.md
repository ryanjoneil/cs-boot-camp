Techniques 02: Backtracking
===========================

Discussion
----------
Backtracking is a computational technique in which one incrementally builds
solutions to a given problem. If a solution under consideration is found to be
not viable, one backtracks to a known acceptable point and attempts to build
another candidate solution.

Backtracking often works well in a recursive setting, as some problems may be
stated such that solutions are composed of smaller solutions. Consider the 
N-Queens problem. In this problem, one seeks to place N queens on a square,
N-by-N chess board such that no queen threatens any other queen. That is,
no two queens are in the same row, column, or are diagonal to each other.

For a chess board of size 4, the following would be an acceptable solution.

	- Q - - 
	- - - Q 
	Q - - - 
	- - Q - 

An algorithm for solving N-Queens might look something like this. Here, we
place a queen in the first row, then add queens recursively to the remaining
rows. If a queen can be placed in each row without conflicts, then the
algorithm terminates. Otherwise, it backtracks and moves the conflicting
queen to the next column. It returns true if all queens can be placed, false
if they cannot.

	function place_queen(row int):
		if row > N:
			return true

		for column in 1 to N:
			put the queen for this row into this column

			if this queen does not conflict with another:
				return place_queen(row + 1)
		
		return false



Exercise 1
----------
Solve the N-Queens problem recursively. Print out a placement of 8 queens
on a chess board.


Follow-Up Questions
-------------------
- What does the final return statement do?
- How does one test for conflicts between queens?
- What is the worst case time complexity of this algorithm?
- How does that take into account testing for conflicts?
- Is this algorithm tail recursive? What is its space complexity?
