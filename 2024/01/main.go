// Package main
package main

import (
	_ "embed"
	"fmt"
	"slices"

	"github.com/agrmohit/aoc/internal/inputs"
	"github.com/agrmohit/aoc/internal/mathutils"
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
		totalDistance += mathutils.Abs(left[i] - right[i])
	}

	return totalDistance
}

func solvePart2(input string) int {
	left, right, err := inputs.ExtractIntPairs(input)
	if err != nil {
		panic(err)
	}

	rightCount := make(map[int]int)
	similarityScore := 0

	for _, rnum := range right {
		rightCount[rnum]++
	}

	for _, lnum := range left {
		similarityScore += lnum * rightCount[lnum]
	}

	return similarityScore
}

func main() {
	part1Solution := solvePart1(input)
	part2Solution := solvePart2(input)

	fmt.Println("Day 01 Part 1 solution:", part1Solution)
	fmt.Println("Day 01 Part 2 solution:", part2Solution)
}
