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

// checkIsSafe takes an int row and returns whether it is safe
func checkIsSafe(row []int) bool {
	isIncreasing := row[1] > row[0]

	for i := 1; i < len(row); i++ {
		// If 2 consecutive numbers are same, it's unsafe
		if row[i] == row[i-1] {
			return false
		}

		// If `isIncreasing` flips anytime during execution, it's unsafe
		if isIncreasing != (row[i] > row[i-1]) {
			return false
		}

		// If absolute difference between consecutive numbers is not 1, 2 or 3
		if mathutils.Abs(row[i]-row[i-1]) > 3 {
			return false
		}
	}

	return true
}

func solvePart1(input string) int {
	rows, err := inputs.ExtractIntRows(input)
	if err != nil {
		panic(err)
	}

	safeCount := 0

	for _, row := range rows {
		if checkIsSafe(row) {
			safeCount++
		}
	}

	return safeCount
}

func solvePart2(input string) int {
	rows, err := inputs.ExtractIntRows(input)
	if err != nil {
		panic(err)
	}

	safeCount := 0

	for _, row := range rows {
		// Check whether it is safe without removing any element
		if checkIsSafe(row) {
			safeCount++
		} else {
			// Remove each element one by one and check whether it is safe
			for i := range row {
				if checkIsSafe(slices.Concat(row[:i], row[i+1:])) {
					safeCount++
					break
				}
			}
		}
	}

	return safeCount
}

func main() {
	part1Solution := solvePart1(input)
	part2Solution := solvePart2(input)

	fmt.Println("Day 01 Part 1 solution:", part1Solution)
	fmt.Println("Day 01 Part 2 solution:", part2Solution)
}
