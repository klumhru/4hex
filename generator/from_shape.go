package generator

import (
	"github.com/klumhru/4hex/hex"
	"github.com/klumhru/4hex/shapes"
)

// GridFromShape implements hex.GenerateGridFunc. It creates a grid from a shapes.Shape parameter.
func GridFromShape(shape shapes.Shape) (hex.Grid, error) {
	b := shape.GetBounds()
	name := shape.GetName()
	gridPos := hex.Position{Q: b.X, R: b.Y}

	cells := make([][]hex.Cell, b.Height)
	for r := 0; r < b.Height; r++ {
		cells[r] = make([]hex.Cell, b.Width)
		for q := 0; q < b.Width; q++ {
			if _, err := shape.GetColorAt(b.X+q, b.Y+r); err == nil {
				cells[r][q] = hex.NewCell(q, r)
			} else {
				cells[r][q] = nil
			}
		}
	}

	return hex.NewGrid(gridPos, name, b.Width, b.Height, cells), nil
}
