package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed test1.txt
var input string

// parseInput takes string input and returns two 2d int slicees.
//
// The first returned value is a 2d slice containing a slice with 2 integers,
// the first and second page number. The second value is a list of all page
// numbers to be printed, for each set
func parseInput(input string) ([][2]int, [][]int, error) {
	var pagePairs [][2]int
	var pageNumbersList [][]int

	// pageNumbersStarted is set to true when the first half of input has been parsed
	pageNumbersStarted := false

	// Check whether input is empty
	if len(input) == 0 {
		return nil, nil, fmt.Errorf("Input is empty")
	}

	// Split input into separate lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			pageNumbersStarted = true
			continue
		}

		// Parse the first half of input
		if !pageNumbersStarted {
			pagePairsRegex := regexp.MustCompile(`(\d+)\|(\d+)`)
			match := pagePairsRegex.FindStringSubmatch(line)

			if len(match) > 0 {
				// Ignore the error since regex already does the type matching
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				CurrentLinePagePairs := [2]int{num1, num2}
				pagePairs = append(pagePairs, CurrentLinePagePairs)
			}
		} else {
			// Parse the second half of input
			pageNumbersRegex := regexp.MustCompile(`(\d+)`)
			matches := pageNumbersRegex.FindAllStringSubmatch(line, -1)

			if len(matches) > 0 {
				pageNumbers := []int{}
				for _, match := range matches {
					// Ignore the error since regex already does the type matching
					num, _ := strconv.Atoi(match[1])

					pageNumbers = append(pageNumbers, num)
				}

				pageNumbersList = append(pageNumbersList, pageNumbers)
			}
		}
	}

	return pagePairs, pageNumbersList, nil
}

func getPageMap(pageNumbers []int) map[int]int {
	pageMap := map[int]int{}

	for i, pageNum := range pageNumbers {
		pageMap[pageNum] = i
	}

	return pageMap
}

func checkOrder(pagePairs [][2]int, pageMap map[int]int, inOrder bool) (bool, [2]int) {
	for _, pagePair := range pagePairs {
		page1Index, ok1 := pageMap[pagePair[0]]
		page2Index, ok2 := pageMap[pagePair[1]]

		// ok1 and ok2 are true if corresponding page is not printed,
		// therefore checking the order doesnt make sense
		if !ok1 || !ok2 {
			continue
		}

		inOrder = page1Index < page2Index
		if !inOrder {
			return false, pagePair
		}
	}

	// Return empty slice when they are in order
	return inOrder, [2]int{}
}

func solvePart1(input string) int {
	pagePairs, pageNumbersList, err := parseInput(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	sumOfMiddlePageNumbers := 0

	for _, pageNumbers := range pageNumbersList {
		inOrder := true
		pageMap := getPageMap(pageNumbers)

		inOrder, _ = checkOrder(pagePairs, pageMap, inOrder)

		if inOrder {
			sumOfMiddlePageNumbers += pageNumbers[len(pageNumbers)/2]
		}
	}

	return sumOfMiddlePageNumbers
}

func main() {
	part1Solution := solvePart1(input)

	fmt.Println("Day 05 Part 1 solution:", part1Solution)
}
