package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/agrmohit/aoc/internal/inputs"
)

//go:embed test1.txt
var input string

// checkAndUpdatePosition checks the next guard movement following the given rules.
// It updates the grid with new guard position and also records the new position in
// the positionsGrid by incrementing the corresping position count. We return boolean
// true if updates position lies within the grid and false otherwise.
//
// This functions only does one action in one run, either move one space or turn.
// It returns the validity for updated position, not the next position
func checkAndUpdatePosition(grid [][]byte, positionsGrid [][]int) bool {
	currentPosition := [2]int{}
	currentDirection := ' '
	newPosition := [2]int{}

	// Find current and new position
outerLoop:
	for i, row := range grid {
		for j, cell := range row {
			currentPosition[0] = i
			currentPosition[1] = j

			switch cell {
			case '^':
				newPosition[0] = i - 1
				newPosition[1] = j
				currentDirection = '^'
			case '>':
				newPosition[0] = i
				newPosition[1] = j + 1
				currentDirection = '>'
			case 'v':
				newPosition[0] = i + 1
				newPosition[1] = j
				currentDirection = 'v'
			case '<':
				newPosition[0] = i
				newPosition[1] = j - 1
				currentDirection = '<'
			default:
				continue
			}

			// Increment the count of current position in positionsGrid
			positionsGrid[currentPosition[0]][currentPosition[1]]++

			// Set current position to default '.'
			// Between this point in code and where new position is marked,
			// the guard position disappears from the grid. But it is fine as
			// there are no returns in between unless guard moves off grid in
			// which case it is the expected result
			grid[currentPosition[0]][currentPosition[1]] = '.'

			break outerLoop
		}
	}

	// Check if new position is out of grid
	outOfGridVertically := newPosition[0] < 0 || newPosition[0] >= len(grid)
	outOfGridHorizontally := newPosition[1] < 0 || newPosition[1] >= len(grid[0])
	if outOfGridVertically || outOfGridHorizontally {
		return false
	}

	// Check if new position is an obstacle
	if grid[newPosition[0]][newPosition[1]] == '#' {
		// Decrement the count of current position in positionsGrid since
		// it will get incremented again next run
		positionsGrid[currentPosition[0]][currentPosition[1]]--

		switch currentDirection {
		case '^':
			grid[currentPosition[0]][currentPosition[1]] = '>'
		case '>':
			grid[currentPosition[0]][currentPosition[1]] = 'v'
		case 'v':
			grid[currentPosition[0]][currentPosition[1]] = '<'
		case '<':
			grid[currentPosition[0]][currentPosition[1]] = '^'
		}
	} else {
		grid[newPosition[0]][newPosition[1]] = byte(currentDirection)
	}

	return true
}

func countDistinctPositions(positionsGrid [][]int) int {
	count := 0

	for _, row := range positionsGrid {
		for _, cell := range row {
			if cell > 0 {
				count++
			}
		}
	}

	return count
}

func solvePart1(input string) (int, error) {
	grid, err := inputs.ExtractCharacterGrid(input)
	if err != nil {
		return -1, err
	}

	// positionsGrid keeps track of positions visited by the guard
	positionsGrid := make([][]int, len(grid))
	for i := range positionsGrid {
		positionsGrid[i] = make([]int, len(grid[i]))
	}

	currentPositionIsValid := true
	for currentPositionIsValid {
		currentPositionIsValid = checkAndUpdatePosition(grid, positionsGrid)
	}

	return countDistinctPositions(positionsGrid), nil
}

func main() {
	part1Solution, err := solvePart1(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	fmt.Println("Day 06 Part 1 solution:", part1Solution)
}
