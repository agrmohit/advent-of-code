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

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 04 Part 1 solution:", part1Solution)
}
