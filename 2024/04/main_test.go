package main

import (
	_ "embed"
	"testing"
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

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

	t.Run("Day 04 part 2 test 1", func(t *testing.T) {
		want := 1
		got := solvePart2(test3)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})

	t.Run("Day 04 part 2 test 2", func(t *testing.T) {
		want := 9
		got := solvePart2(test4)

		if got != want {
			t.Errorf("Incorrect solution, got %d want %d", got, want)
		}
	})
}
