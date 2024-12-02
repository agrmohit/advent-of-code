package main

import (
	_ "embed"
	"fmt"

	"github.com/agrmohit/advent-of-code/internal/inputs"
	mathinternal "github.com/agrmohit/advent-of-code/internal/math"
)

//go:embed input.txt
var input string

func solvePart1(input string) int {
	rows, err := inputs.ExtractIntRows(input)
	if err != nil {
		panic(err)
	}

	safeCount := 0

	for _, row := range rows {
		isSafe := true
		isIncreasing := row[1] > row[0]

		for i := 1; i < len(row); i++ {
			// If 2 consecutive numbers are same, it's unsafe
			if row[i] == row[i-1] {
				isSafe = false
				break
			}

			// If `isIncreasing` flips anytime during execution, it's unsafe
			if isIncreasing != (row[i] > row[i-1]) {
				isSafe = false
				break
			}

			// If absolute difference between consecutive numbers is not 1, 2 or 3
			if mathinternal.Abs(row[i]-row[i-1]) > 3 {
				isSafe = false
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}

	return safeCount
}

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 02 Part 1 solution:", part1Solution)
}
