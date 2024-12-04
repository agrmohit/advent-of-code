package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/agrmohit/aoc/internal/grid"
	"github.com/agrmohit/aoc/internal/inputs"
)

//go:embed input.txt
var input string

func solvePart1(input string) int {
	characterGrid, err := inputs.ExtractCharacterGrid(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	directions := []grid.SearchDirection{
		grid.HorizontalForward,
		grid.HorizontalBackward,
		grid.VerticalDown,
		grid.VerticalUp,
		grid.DiagonalUpwardsLeft,
		grid.DiagonalUpwardsRight,
		grid.DiagonalDownwardsLeft,
		grid.DiagonalDownwardsRight,
	}
	words := grid.FindWordsInGrid(characterGrid, []string{"XMAS"}, directions)

	return len(words)
}

func solvePart2(input string) int {
	characterGrid, err := inputs.ExtractCharacterGrid(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	// Thats wayyy too many nested slices but it workd :-)
	subgrids := [][][]byte{}
	subgridCount := 0

	subgrids = append(subgrids, [][]byte{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	})
	subgrids = append(subgrids, [][]byte{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	})
	subgrids = append(subgrids, [][]byte{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	})
	subgrids = append(subgrids, [][]byte{
		{'S', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'M'},
	})

	for _, subgrid := range subgrids {
		subgridLocationList := grid.FindSubgridsInGrid(characterGrid, subgrid, '.')
		subgridCount += len(subgridLocationList)
	}

	return subgridCount
}

func main() {
	part1Solution := solvePart1(input)
	part2Solution := solvePart2(input)

	fmt.Println("Day 04 Part 1 solution:", part1Solution)
	fmt.Println("Day 04 Part 2 solution:", part2Solution)
}
