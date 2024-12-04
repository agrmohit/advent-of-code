package main

import (
	_ "embed"
	"testing"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

func TestSolution(t *testing.T) {
	t.Run("Day 04 part 1 test 1", func(t *testing.T) {
		want := 4
		got := solvePart1(test1)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})

	t.Run("Day 04 part 1 test 2", func(t *testing.T) {
		want := 18
		got := solvePart1(test2)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})
}
