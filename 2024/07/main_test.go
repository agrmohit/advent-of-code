package main

import (
	_ "embed"
	"testing"
)

//go:embed test1.txt
var test1 string

func TestSolution(t *testing.T) {
	t.Run("Day 07 part 1 test 1", func(t *testing.T) {
		want := 3749
		got := solvePart1(test1)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})
}
