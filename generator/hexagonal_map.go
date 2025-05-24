package generator

import (
	"fmt"
	"math"

	"github.com/klumhru/4hex/hex" // Corrected module path
)

// GenerateHexagonalGrid creates a grid with cells forming a hexagonal shape.
// The hexagon is centered within the given width and height.
// Cells outside the hexagon are nil. This function implements the hex.GenerateGridFunc interface.
func GenerateHexagonalGrid(gridPos hex.Position, name string, width, height int) (hex.Grid, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width (%d) and height (%d) must be positive", width, height)
	}

	cells := make([][]hex.Cell, height)
	for r := 0; r < height; r++ {
		cells[r] = make([]hex.Cell, width)
	}

	// Determine the radius of the hexagon.
	// A hexagon of radius N has (2*N + 1) cells across its widest part.
	// So, 2*N + 1 <= min(width, height)  => N <= (min(width, height) - 1) / 2.
	minDimension := math.Min(float64(width), float64(height))
	hexagonRadiusInt := int((minDimension - 1) / 2)

	if hexagonRadiusInt < 0 {
		hexagonRadiusInt = 0 // Smallest hexagon is a single cell (radius 0)
	}
	hexagonRadiusFloat := float64(hexagonRadiusInt)

	// Center of the grid array (using floating point for precision with even/odd dimensions)
	centerQ := float64(width-1) / 2.0
	centerR := float64(height-1) / 2.0

	for rIdx := 0; rIdx < height; rIdx++ {
		for qIdx := 0; qIdx < width; qIdx++ {
			// Current cell's axial coordinates (q_idx, r_idx)
			qFloat := float64(qIdx)
			rFloat := float64(rIdx)

			// Convert to coordinates relative to the hexagon's center
			relQ := qFloat - centerQ
			relR := rFloat - centerR
			// The third cube coordinate, s = -q-r. So relS = -(relQ + relR)
			relS := -relQ - relR

			// Check if the cell is within the hexagon using cube coordinates.
			// A cell is in a hexagon of radius N if max(|relQ|, |relR|, |relS|) <= N.
			// Using a small epsilon for floating point comparisons might be more robust,
			// but direct comparison should work for this geometry.
			if math.Max(math.Abs(relQ), math.Max(math.Abs(relR), math.Abs(relS))) <= hexagonRadiusFloat {
				// Pass qIdx and rIdx directly to NewCell, as per its definition.
				cells[rIdx][qIdx] = hex.NewCell(qIdx, rIdx)
			} else {
				// Cell is outside the hexagon, leave it as nil
				cells[rIdx][qIdx] = nil
			}
		}
	}

	return hex.NewGrid(gridPos, name, width, height, cells), nil
}
