package hex

import (
	"fmt"

	"github.com/klumhru/4hex/shapes"
)

// AddLayer adds a new layer to the map by generating a grid using the provided function.
func (m *concreteMap) AddLayer(f GenerateGridFunc) error {
	if f == nil {
		return fmt.Errorf("generate function cannot be nil")
	}
	// Create a shape representing the map's bounds
	shape := shapes.NewRectangle(0, 0, m.width, m.height, fmt.Sprintf("Layer_%d", len(m.grids)))
	grid, err := f(shape)
	if err != nil {
		return fmt.Errorf("failed to generate grid: %w", err)
	}
	return m.AddGrid(grid)
}
