package main

import (
	_ "embed"
	"testing"
)

//go:embed test1.txt
var test1 string

func TestSolution(t *testing.T) {
	t.Run("Day 01 part 1", func(t *testing.T) {
		want := 11
		got := solvePart1(test1)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})

	t.Run("Day 01 part 2", func(t *testing.T) {
		want := 31
		got := solvePart2(test1)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})
}
