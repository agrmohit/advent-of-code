package grid

// SubgridLocation represents the coordinates of a found subgrid in the grid
type SubgridLocation struct {
	Subgrid [][]byte // Subgrid is the subgrid that was found in the grid
	StartX  int      // StartX is the starting x-coordinate of the subgrid in the grid
	StartY  int      // StartY is the starting y-coordinate of the subgrid in the grid
	EndX    int      // EndX is the ending x-coordinate of the subgrid in the grid
	EndY    int      // EndY is the ending y-coordinate of the subgrid in the grid
}

// FindSubgridsInGrid searches for subgrids in a given grid
//
// '.' could be any character in the grid. All other characters need an exact match.
// wildcardCharacter in subgrid can match any character in grid
func FindSubgridsInGrid(grid, subgrid [][]byte, wildcardCharacter byte) []SubgridLocation {
	var result []SubgridLocation

	for i, row := range grid {
		for j := range row {
			sgl := checkIndexForSubgrid(i, j, grid, subgrid, wildcardCharacter)
			if len(sgl.Subgrid) != 0 {
				result = append(result, sgl)
			}
		}
	}
	return result
}

func checkIndexForSubgrid(i, j int, grid, subgrid [][]byte, wildcardCharacter byte) SubgridLocation {
	var result SubgridLocation

	for si, row := range subgrid {
		for sj := range row {
			// Check whether subgrid exceeds the bounds of grid
			if i+len(subgrid) > len(grid) || j+len(subgrid[0]) > len(grid[0]) {
				return result
			}

			// Check whether the character matches or is the wildcard character
			if !(subgrid[si][sj] == wildcardCharacter || subgrid[si][sj] == grid[i+si][j+sj]) {
				return result
			}
		}
	}

	result.Subgrid = subgrid
	result.StartX = j
	result.StartY = i
	result.EndX = j + len(subgrid[0]) - 1
	result.EndY = i + len(subgrid) - 1

	return result
}
