// Package grid contains functions used for grid (2d slice) operations
package grid

import (
	"github.com/agrmohit/aoc/internal/strutils"
)

// SearchDirection represents possible directions for word searching in a grid
type SearchDirection int

const (
	HorizontalForward      SearchDirection = iota // HorizontalForward searches from left to right
	HorizontalBackward                            // HorizontalBackward searches from right to left
	VerticalDown                                  // VerticalDown searches from top to bottom
	VerticalUp                                    // VerticalUp searches from bottom to top
	DiagonalUpwardsLeft                           // DiagonalForward searches from bottom-right to top-left
	DiagonalUpwardsRight                          // DiagonalForward searches from bottom-left to top-right
	DiagonalDownwardsLeft                         // DiagonalBackward searches from top-right to bottom-left
	DiagonalDownwardsRight                        // DiagonalBackward searches from top-left to bottom-right
)

// WordLocation represents the coordinates of a found word in the grid
type WordLocation struct {
	Word   string // Word is the word that was found in the grid
	StartX int    // StartX is the starting x-coordinate of the word in the grid
	StartY int    // StartY is the starting y-coordinate of the word in the grid
	EndX   int    // EndX is the ending x-coordinate of the word in the grid
	EndY   int    // EndY is the ending y-coordinate of the word in the grid
}

// FindWordsInGrid searches for words in a given grid in specified directions
//
// BUG: This may not work with characters that don't fit in one byte
func FindWordsInGrid(grid [][]byte, words []string, directions []SearchDirection) []WordLocation {
	var result []WordLocation

	for _, word := range words {
		for i, row := range grid {
			for j := range row {
				for _, direction := range directions {
					wl := checkIndexForWord(grid, i, j, word, direction)
					if wl.Word != "" {
						result = append(result, wl)
					}
				}
			}
		}
	}

	return result
}

func checkIndexForWord(grid [][]byte, i, j int, word string, direction SearchDirection) WordLocation {
	var result WordLocation

	found := true
	ogWord := word

	// Check whether the word needs to be reversed (so that we dont duplicate logic)
	switch direction {
	case HorizontalBackward, VerticalUp, DiagonalUpwardsLeft, DiagonalUpwardsRight:
		word = strutils.Reverse(word)
	}

	switch direction {
	case HorizontalForward, HorizontalBackward:
		if !(j+len(word) > len(grid[i])) {
			for charIndex, char := range word {
				if grid[i][j+charIndex] != byte(char) {
					found = false
					break
				}
			}

			if found {
				result.Word = ogWord
				result.StartX = j
				result.StartY = i
				result.EndX = j + len(word) - 1
				result.EndY = i
			}
		}

	case VerticalDown, VerticalUp:
		if !(i+len(word) > len(grid)) {
			for charIndex, char := range word {
				if grid[i+charIndex][j] != byte(char) {
					found = false
					break
				}
			}

			if found {
				result.Word = ogWord
				result.StartX = j
				result.StartY = i
				result.EndX = j
				result.EndY = i + len(word) - 1
			}
		}

	case DiagonalDownwardsRight, DiagonalUpwardsLeft:
		if !(j+len(word) > len(grid[i])) && !(i+len(word) > len(grid)) {
			for charIndex, char := range word {
				if grid[i+charIndex][j+charIndex] != byte(char) {
					found = false
					break
				}
			}

			if found {
				result.Word = ogWord
				result.StartX = j
				result.StartY = i
				result.EndX = j + len(word) - 1
				result.EndY = i + len(word) - 1
			}
		}

	case DiagonalDownwardsLeft, DiagonalUpwardsRight:
		if !(j-len(word)+1 < 0) && !(i+len(word) > len(grid)) {
			for charIndex, char := range word {
				if grid[i+charIndex][j-charIndex] != byte(char) {
					found = false
					break
				}
			}

			if found {
				result.Word = ogWord
				result.StartX = j
				result.StartY = i
				result.EndX = j - len(word) + 1
				result.EndY = i + len(word) - 1
			}
		}
	}

	return result
}
