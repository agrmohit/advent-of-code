package grid

import (
	"testing"

	"github.com/agrmohit/aoc/internal/inputs"
)

func TestFindSubgridsInGrid(t *testing.T) {
	input := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`
	want := 9

	characterGrid, _ := inputs.ExtractCharacterGrid(input)

	subgrids := [][][]byte{}
	subgridCount := 0

	subgrids = append(subgrids, [][]byte{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	})
	subgrids = append(subgrids, [][]byte{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	})
	subgrids = append(subgrids, [][]byte{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	})
	subgrids = append(subgrids, [][]byte{
		{'S', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'M'},
	})

	for _, subgrid := range subgrids {
		subgridLocationList := FindSubgridsInGrid(characterGrid, subgrid, '.')
		subgridCount += len(subgridLocationList)
	}
	got := subgridCount

	if got != want {
		t.Errorf("got %v, want%v", got, want)
	}
}
