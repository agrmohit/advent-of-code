// Package inputs contains functions used to parse input data
package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

// ExtractIntPairs takes 2 column input and returns 2 int slices
func ExtractIntPairs(input string) ([]int, []int, error) {
	var left, right []int

	// Check whether input is empty
	if len(input) == 0 {
		return nil, nil, fmt.Errorf("Input is empty")
	}

	// Split input into separate lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		// Split each line into columns separated by whitespace
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

	return left, right, nil
}

// ExtractIntRows takes multiple lines of whitespace separated numbers and
// parses it into a 2d int slice
func ExtractIntRows(input string) ([][]int, error) {
	var result [][]int

	// Check whether input is empty
	if len(input) == 0 {
		return nil, fmt.Errorf("Input is empty")
	}

	// Split input into separate lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		// Split each line into columns separated by whitespace
		columns := strings.Fields(line)

		// Create an int slice with length equal to number of columns
		row := make([]int, len(columns))

		// Convert the strings to int and store in `row` slice
		for i, column := range columns {
			// Parse each column as numbers
			num, err := strconv.Atoi(column)
			if err != nil {
				return nil, fmt.Errorf("invalid number %q in row %q", column, row)
			}

			row[i] = num
		}

		// Append the int row to result
		result = append(result, row)
	}

	return result, nil
}

func ExtractCharacterGrid(input string) ([][]byte, error) {
	var result [][]byte

	// Check whether input is empty
	if len(input) == 0 {
		return nil, fmt.Errorf("Input is empty")
	}

	// Split input into separate lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	expectedLineLength := len(lines[0])
	for _, line := range lines {
		if len(line) != expectedLineLength {
			return nil, fmt.Errorf("Grid doesn't have consistent line lengths")
		}

		result = append(result, []byte(line))
	}

	return result, nil
}
