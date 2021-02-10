// Missionaries and Cannibals for CSI 380
// This program solves the classic Missionaries and Cannibals problem
// Spencer Bellucci
// CSI 380
// 2/10/21

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
	westSide := false
	eastSide := false

	// make a negative false
	if pos.westMissionaries < 0 || pos.eastMissionaries < 0 || pos.westCannibals < 0 || pos.eastCannibals < 0 {
		return false
	}

	// make anything more than 3 false
	if pos.westMissionaries > 3 || pos.eastMissionaries > 3 || pos.westCannibals > 3 || pos.eastCannibals > 3 {
		return false
	}

	// check if west side is valid
	if pos.westMissionaries >= pos.westCannibals || pos.westMissionaries == 0 {
		westSide = true
	} else {
		westSide = false
	}

	//check if east side is valid
	if pos.eastMissionaries >= pos.eastCannibals || pos.eastMissionaries == 0 {
		eastSide = true
	} else {
		eastSide = false
	}

	//if both sides are valid return true
	if westSide && eastSide == true {
		return true
	} else {
		return false
	}
}

// What are all of the next positions we can go to legally from the current position
// Returns nil if there are no valid positions
func (pos position) successors() []position {
	var allPositions []position

	// if boat is on west bank
	if pos.boatOnWestBank == true {
		// check both sides with - 1 missionary, then - 2
		for i := 1; i < 3; i++ {
			pos.westMissionaries = pos.westMissionaries - i
			pos.eastMissionaries = pos.eastMissionaries + i
			pos.boatOnWestBank = false
			if valid(pos) {
				newPos := pos
				allPositions = append(allPositions, newPos)
			}
			// set back to normal
			pos.eastMissionaries = pos.eastMissionaries - i
			pos.westMissionaries = pos.westMissionaries + i
			pos.boatOnWestBank = true
		}

		// do the same thing with cannibals
		for i := 1; i < 3; i++ {
			pos.westCannibals = pos.westCannibals - i
			pos.eastCannibals = pos.eastCannibals + i
			pos.boatOnWestBank = false
			if valid(pos) {
				newPos := pos
				allPositions = append(allPositions, newPos)
			}
			pos.eastCannibals = pos.eastCannibals - i
			pos.westCannibals = pos.westCannibals + i
			pos.boatOnWestBank = true
		}

		// check with one missionary and one cannibal
		pos.westCannibals = pos.westCannibals - 1
		pos.eastCannibals = pos.eastCannibals + 1
		pos.westMissionaries = pos.westMissionaries - 1
		pos.eastMissionaries = pos.eastMissionaries + 1
		pos.boatOnWestBank = false
		if valid(pos) {
			newPos := pos
			allPositions = append(allPositions, newPos)
		}
		pos.eastCannibals = pos.eastCannibals - 1
		pos.westCannibals = pos.westCannibals + 1
		pos.westMissionaries = pos.westMissionaries + 1
		pos.eastMissionaries = pos.eastMissionaries - 1
		pos.boatOnWestBank = true

	} else if pos.boatOnWestBank == false {
		// check both sides with - 1 missionary, then - 2
		for i := 1; i < 3; i++ {
			pos.eastMissionaries = pos.eastMissionaries - i
			pos.westMissionaries = pos.westMissionaries + i
			pos.boatOnWestBank = true
			if valid(pos) {
				newPos := pos
				allPositions = append(allPositions, newPos)
			}
			// set back to normal
			pos.westMissionaries = pos.westMissionaries - i
			pos.eastMissionaries = pos.eastMissionaries + i
			pos.boatOnWestBank = false
		}

		// do the same thing with cannibals
		for i := 1; i < 3; i++ {
			pos.eastCannibals = pos.eastCannibals - i
			pos.westCannibals = pos.westCannibals + i
			pos.boatOnWestBank = true
			if valid(pos) {
				newPos := pos
				allPositions = append(allPositions, newPos)
			}
			pos.westCannibals = pos.westCannibals - i
			pos.eastCannibals = pos.eastCannibals + i
			pos.boatOnWestBank = false
		}

		// check with one missionary and one cannibal
		pos.eastCannibals = pos.eastCannibals - 1
		pos.westCannibals = pos.westCannibals + 1
		pos.eastMissionaries = pos.eastMissionaries - 1
		pos.westMissionaries = pos.westMissionaries + 1
		pos.boatOnWestBank = true
		if valid(pos) {
			newPos := pos
			allPositions = append(allPositions, newPos)
		}
		pos.westCannibals = pos.westCannibals - 1
		pos.eastCannibals = pos.eastCannibals + 1
		pos.eastMissionaries = pos.eastMissionaries + 1
		pos.westMissionaries = pos.westMissionaries - 1
		pos.boatOnWestBank = false

	}

	// return array if there were any valid positions
	if len(allPositions) != 0 {
		return allPositions
	} else {
		return nil
	}

}

// A recursive depth-first search that goes through to find the goal and returns the path to get there
// Returns nil if no solution found
func dfs(start position, goal position, solution []position, visited map[position]bool) []position {
	// mark solution as visited
	visited[start] = true
	// add current move to solution
	newSolution := append(solution, start)

	// if this state != goal
	if start != goal {
		// find all valid moves from this state
		var branches = start.successors()

		// if no valid moves
		if branches == nil {
			// go back to previous move
			newSolution = solution[:len(solution)-1]
			return dfs(start, goal, newSolution, visited)
		}

		// check for a move that has not been visited
		for i := 0; i < len(branches); i++ {
			_, ok := visited[branches[i]]
			if !ok {
				// call function with new move
				return dfs(branches[i], goal, newSolution, visited)
			}
		}

		// go back to previous move
		newSolution = solution[:len(solution)-1]
		return dfs(solution[len(solution)-1], goal, newSolution, visited)

	} else {
		// return solution
		newSolution = solution[1 : len(solution)+1]
		return newSolution
	}
}

func main() {
	start := position{boatOnWestBank: true, westMissionaries: 3, westCannibals: 3, eastMissionaries: 0, eastCannibals: 0}
	goal := position{boatOnWestBank: false, westMissionaries: 0, westCannibals: 0, eastMissionaries: 3, eastCannibals: 3}
	solution := dfs(start, goal, []position{start}, make(map[position]bool))
	fmt.Println(solution)
}
