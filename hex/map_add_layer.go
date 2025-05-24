package hex

import "fmt"

// AddLayer adds a new layer to the map by generating a grid using the provided function.
func (m *concreteMap) AddLayer(f GenerateGridFunc) error {
	if f == nil {
		return fmt.Errorf("generate function cannot be nil")
	}
	grid, err := f(Position{}, fmt.Sprintf("Layer_%d", len(m.grids)), m.width, m.height)
	if err != nil {
		return fmt.Errorf("failed to generate grid: %w", err)
	}
	return m.AddGrid(grid)
}
