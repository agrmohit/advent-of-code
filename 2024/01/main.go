package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func getInputs(input string) ([]int, []int, error) {
	var left, right []int

	// Check whether input is empty
	if len(input) == 0 {
		return nil, nil, fmt.Errorf("Input is empty")
	}

	// Split input into separate lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		columns := strings.Fields(line)
		if len(columns) != 2 {
			return nil, nil, fmt.Errorf("Invalid line input: %s", line)
		}

		// Parse the left and right columns as numbers
		leftNum, err := strconv.Atoi(columns[0])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid number in left column of line: %s", columns[0])
		}

		rightNum, err := strconv.Atoi(columns[1])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid number in right column of line: %s", columns[1])
		}

		// Append the numbers to respective slices
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	// Check whether the number of elements in left and right slices are the same
	// This check is probably not needed
	if len(left) != len(right) {
		return nil, nil, fmt.Errorf("Length of left and right slices differ: %d, %d", len(left), len(right))
	}

	return left, right, nil
}

func solvePart1(input string) int {
	left, right, err := getInputs(input)
	if err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0

	for i := range left {
		totalDistance += abs(left[i] - right[i])
	}

	return totalDistance
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 01 Part 1 solution:", part1Solution)
}
