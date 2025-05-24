package generator

import (
	"fmt"
	"testing"

	"github.com/klumhru/4hex/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// expectedHexCells calculates the number of cells in a regular hexagon of a given radius.
// A hexagon of radius N has 3*N*(N+1) + 1 cells.
func expectedHexCells(radius int) int {
	if radius < 0 {
		return 0 // Should ideally not happen with valid radius calculation
	}
	return 3*radius*(radius+1) + 1
}

func TestGenerateHexagonalGrid(t *testing.T) {
	testCases := []struct {
		desc                string
		inputGridPos        hex.Position
		inputName           string
		inputWidth          int
		inputHeight         int
		expectedNonNilCells int
		expectedError       bool
		// For specific checks on cell presence/absence if needed later
		// cellChecks          []struct{ q, r int; expectNil bool }
	}{
		{
			desc:                "5x5_radius2",
			inputGridPos:        hex.Position{Q: 1, R: 2},
			inputName:           "TestHexGrid_5x5",
			inputWidth:          5,
			inputHeight:         5,
			expectedNonNilCells: expectedHexCells(2), // minDim=5, radius=(5-1)/2=2. Cells=19
			expectedError:       false,
		},
		{
			desc:                "1x1_radius0_single_cell",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "TestHexGrid_1x1",
			inputWidth:          1,
			inputHeight:         1,
			expectedNonNilCells: expectedHexCells(0), // minDim=1, radius=(1-1)/2=0. Cells=1
			expectedError:       false,
		},
		{
			desc:                "2x2_radius0_zero_cells_due_to_centering",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "TestHexGrid_2x2",
			inputWidth:          2,
			inputHeight:         2,
			expectedNonNilCells: 0, // minDim=2, radius=(2-1)/2=0. Centering on 0.5,0.5 means no integer cell hits for radius 0.
			expectedError:       false,
		},
		{
			desc:                "3x3_radius1",
			inputGridPos:        hex.Position{Q: -1, R: -1},
			inputName:           "TestHexGrid_3x3",
			inputWidth:          3,
			inputHeight:         3,
			expectedNonNilCells: expectedHexCells(1), // minDim=3, radius=(3-1)/2=1. Cells=7
			expectedError:       false,
		},
		{
			desc:                "7x4_radius1_height_limits",
			inputGridPos:        hex.Position{Q: 5, R: 5},
			inputName:           "TestHexGrid_7x4",
			inputWidth:          7,
			inputHeight:         4,
			expectedNonNilCells: expectedHexCells(1), // minDim=4, radius=(4-1)/2=1. Cells=7
			expectedError:       false,
		},
		{
			desc:                "4x7_radius1_width_limits",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "TestHexGrid_4x7",
			inputWidth:          4,
			inputHeight:         7,
			expectedNonNilCells: expectedHexCells(1), // minDim=4, radius=(4-1)/2=1. Cells=7
			expectedError:       false,
		},
		{
			desc:                "2x1_radius0_zero_cells_due_to_centering",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "TestHexGrid_2x1",
			inputWidth:          2,
			inputHeight:         1,
			expectedNonNilCells: 0, // minDim=1, radius=0. Width 2 means centerQ=0.5. No cells.
			expectedError:       false,
		},
		{
			desc:                "1x2_radius0_zero_cells_due_to_centering",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "TestHexGrid_1x2",
			inputWidth:          1,
			inputHeight:         2,
			expectedNonNilCells: 0, // minDim=1, radius=0. Height 2 means centerR=0.5. No cells.
			expectedError:       false,
		},
		{
			desc:                "invalid_width_zero",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "InvalidGrid",
			inputWidth:          0,
			inputHeight:         5,
			expectedNonNilCells: 0,
			expectedError:       true,
		},
		{
			desc:                "invalid_height_negative",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "InvalidGrid",
			inputWidth:          5,
			inputHeight:         -1,
			expectedNonNilCells: 0,
			expectedError:       true,
		},
		{
			desc:                "invalid_width_and_height_zero",
			inputGridPos:        hex.Position{Q: 0, R: 0},
			inputName:           "InvalidGrid",
			inputWidth:          0,
			inputHeight:         0,
			expectedNonNilCells: 0,
			expectedError:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			grid, err := GenerateHexagonalGrid(tc.inputGridPos, tc.inputName, tc.inputWidth, tc.inputHeight)

			if tc.expectedError {
				assert.Error(err, "Expected an error for invalid input")
				assert.Nil(grid, "Grid should be nil on error")
			} else {
				assert.NoError(err, "GenerateHexagonalGrid returned an unexpected error")
				require.NotNil(grid, "Grid should not be nil for valid input")

				assert.Equal(tc.inputName, grid.GetName(), "Grid name mismatch")
				assert.Equal(tc.inputGridPos, grid.GetPosition(), "Grid position mismatch")
				assert.Equal(tc.inputWidth, grid.GetWidth(), "Grid width mismatch")
				assert.Equal(tc.inputHeight, grid.GetHeight(), "Grid height mismatch")

				actualNonNilCells := 0
				if grid.GetWidth() > 0 && grid.GetHeight() > 0 {
					destCells := make([][]hex.Cell, grid.GetHeight())
					for i := range destCells {
						destCells[i] = make([]hex.Cell, grid.GetWidth())
					}

					copyErr := grid.CopyCellsTo(destCells)
					require.NoError(copyErr, "CopyCellsTo should succeed on a valid grid")

					for rIdx := 0; rIdx < grid.GetHeight(); rIdx++ {
						for qIdx := 0; qIdx < grid.GetWidth(); qIdx++ {
							cell := destCells[rIdx][qIdx]
							if cell != nil {
								actualNonNilCells++
								cellPos := cell.GetPosition()
								assert.Equal(qIdx, cellPos.Q, fmt.Sprintf("Cell at grid[%d][%d] has incorrect Q position. Expected %d, got %d", rIdx, qIdx, qIdx, cellPos.Q))
								assert.Equal(rIdx, cellPos.R, fmt.Sprintf("Cell at grid[%d][%d] has incorrect R position. Expected %d, got %d", rIdx, qIdx, rIdx, cellPos.R))
							}
						}
					}
				}
				assert.Equal(tc.expectedNonNilCells, actualNonNilCells, "Mismatch in number of non-nil cells")
			}
		})
	}
}
