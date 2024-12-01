package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/agrmohit/advent-of-code/internal/inputs"
	mathinternal "github.com/agrmohit/advent-of-code/internal/math"
)

//go:embed input.txt
var input string

func solvePart1(input string) int {
	left, right, err := inputs.ExtractIntPairs(input)
	if err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0

	for i := range left {
		totalDistance += mathinternal.Abs(left[i] - right[i])
	}

	return totalDistance
}

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 01 Part 1 solution:", part1Solution)
}
