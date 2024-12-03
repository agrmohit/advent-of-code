package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func solvePart1(input string) int {
	rexp := regexp.MustCompile(`mul\((\d{1,3})\,(\d{1,3})\)`)
	matches := rexp.FindAllStringSubmatch(input, -1)

	result := 0

	for _, match := range matches {
		// Ignore the error since regex already does the type matching
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])

		result += num1 * num2
	}

	return result
}

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 03 Part 1 solution:", part1Solution)
}
