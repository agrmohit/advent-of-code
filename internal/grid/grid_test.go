package grid

import (
	"testing"

	"github.com/agrmohit/aoc/internal/inputs"
)

func TestFindWordsInGrid(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	want := 18

	characterGrid, _ := inputs.ExtractCharacterGrid(input)
	directions := []SearchDirection{
		HorizontalForward,
		HorizontalBackward,
		VerticalDown,
		VerticalUp,
		DiagonalUpwardsLeft,
		DiagonalUpwardsRight,
		DiagonalDownwardsLeft,
		DiagonalDownwardsRight,
	}
	WordLocationList := FindWordsInGrid(characterGrid, []string{"XMAS"}, directions)
	got := len(WordLocationList)

	if got != want {
		t.Errorf("got %v, want%v", got, want)
	}
}
