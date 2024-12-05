package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
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

// getPageMap generates a map with page number as key and its index as value
func getPageMap(pageNumbers []int) map[int]int {
	pageMap := map[int]int{}

	for i, pageNum := range pageNumbers {
		pageMap[pageNum] = i
	}

	return pageMap
}

// checkOrder checks whether the given page numbers are in order following given rules
//
// If they are in order, it returns boolean true and and empty int slice.
// If they are not in order, it returns boolean false along with an int slice
// containing the page numbers that violate rules.
func checkOrder(pagePairs [][2]int, pageMap map[int]int) (bool, [2]int) {
	inOrder := true

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
		pageMap := getPageMap(pageNumbers)

		inOrder, _ := checkOrder(pagePairs, pageMap)

		if inOrder {
			sumOfMiddlePageNumbers += pageNumbers[len(pageNumbers)/2]
		}
	}

	return sumOfMiddlePageNumbers
}

// swap does in-place arithmetic swap on given slice for given indices and returns new slice
func swap(arr []int, a, b int) []int {
	arr[a] = arr[a] + arr[b]
	arr[b] = arr[a] - arr[b]
	arr[a] = arr[a] - arr[b]

	return arr
}

// fixOrder fixes the order of given page numbers following provided rules.
//
// It is essentially a sorting algorithms with custom rules. It takes sorting
// rules as a 2d slice in pagePairs and list of page numbers as a slice in
// pagePairs and returns the sorted/fixed order page numbers.
func fixOrder(pagePairs [][2]int, pageNumbers []int) []int {
	for i := 0; i < len(pagePairs); i++ {
		pageMap := getPageMap(pageNumbers)
		inOrder, pagePairsNotInOrder := checkOrder(pagePairs, pageMap)

		// Return page numbers when they are in order
		if inOrder {
			return pageNumbers
		}

		// If page numbers are not in order, swap the offending pair of pages
		a := pageMap[pagePairsNotInOrder[0]]
		b := pageMap[pagePairsNotInOrder[1]]
		pageNumbers = swap(pageNumbers, a, b)
	}

	return []int{}
}

func solvePart2(input string) int {
	pagePairs, pageNumbersList, err := parseInput(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	sumOfMiddlePageNumbers := 0

	for _, pageNumbers := range pageNumbersList {
		pageMap := getPageMap(pageNumbers)

		inOrder, _ := checkOrder(pagePairs, pageMap)

		if !inOrder {
			fixedPageNumbers := fixOrder(pagePairs, pageNumbers)
			if len(fixedPageNumbers) == 0 {
				log.Fatalf("ERROR: %v", fmt.Errorf("Unable to fix page numbers"))
			}
			sumOfMiddlePageNumbers += fixedPageNumbers[len(fixedPageNumbers)/2]
		}
	}

	return sumOfMiddlePageNumbers
}

func main() {
	part1Solution := solvePart1(input)
	part2Solution := solvePart2(input)

	fmt.Println("Day 05 Part 1 solution:", part1Solution)
	fmt.Println("Day 05 Part 2 solution:", part2Solution)
}
