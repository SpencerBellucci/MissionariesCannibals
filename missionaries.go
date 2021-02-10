// Missionaries and Cannibals for CSI 380
// This program solves the classic Missionaries and Cannibals problem

package main

import "fmt"

// A representation of the state of the game
type position struct {
	boatOnWestBank   bool // true is west bank, false is east bank
	westMissionaries int  // west bank missionaries
	westCannibals    int  // west bank cannibals
	eastMissionaries int  // east bank missionaries
	eastCannibals    int  // east bank cannibals
}

// Is this a legal position? In particular, does it have
// more cannibals than missionaries on either bank? Because that is illegal.
func valid(pos position) bool {
	// YOUR CODE GOES HERE
}

// What are all of the next positions we can go to legally from the current position
// Returns nil if there are no valid positions
func (pos position) successors() []position {
	// YOUR CODE GOES HERE
}

// A recursive depth-first search that goes through to find the goal and returns the path to get there
// Returns nil if no solution found
func dfs(start position, goal position, solution []position, visited map[position]bool) []position {
	// YOUR CODE GOES HERE
}

func main() {
	start := position{boatOnWestBank: true, westMissionaries: 3, westCannibals: 3, eastMissionaries: 0, eastCannibals: 0}
	goal := position{boatOnWestBank: false, westMissionaries: 0, westCannibals: 0, eastMissionaries: 3, eastCannibals: 3}
	solution := dfs(start, goal, []position{start}, make(map[position]bool))
	fmt.Println(solution)
}
