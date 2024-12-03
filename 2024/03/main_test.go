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
	t.Run("Day 03 part 1", func(t *testing.T) {
		want := 161
		got := solvePart1(test1)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})

	t.Run("Day 03 part 2", func(t *testing.T) {
		want := 48
		got := solvePart2(test2)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})
}
