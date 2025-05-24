package hex

import "fmt"

// Grid represents a hexagonal grid. It is a two-dimensional array of hexagonal cells.
// Use grids as layers to create complex maps.
type Grid interface {
	// GetPosition returns the position of the grid in the map.
	GetPosition() Position
	// GetName returns the name of the grid.
	GetName() string
	// GetCellAt returns the cell at the specified q and r coordinates.
	GetCellAt(q, r int) (Cell, error)
	// GetCellAtPosition returns the cell at the specified position.
	GetCellAtPosition(pos Position) (Cell, error)
	// GetCellAtIndex returns the cell at the specified index.
	GetCellAtIndex(index int) (Cell, error)
	// GetCellCount returns the total number of cells in the grid.
	GetCellCount() int
	// GetWidth and GetHeight return the dimensions of the grid.
	GetWidth() int
	GetHeight() int
	// CopyCellsTo copies the grid's cells into the destination slice.
	// It returns an error if the destination slice has incorrect dimensions.
	CopyCellsTo(destination [][]Cell) error
}

// concreteGrid implements the Grid interface.
type concreteGrid struct {
	position Position
	name     string
	width    int
	height   int
	cells    [][]Cell
}

func (g *concreteGrid) GetPosition() Position {
	return g.position
}
func (g *concreteGrid) GetName() string {
	return g.name
}
func (g *concreteGrid) GetCellAt(q, r int) (Cell, error) {
	if q < 0 || r < 0 || q >= g.width || r >= g.height {
		return nil, fmt.Errorf("cell at (%d, %d) is out of bounds", q, r)
	}
	return g.cells[r][q], nil
}
func (g *concreteGrid) GetCellAtPosition(pos Position) (Cell, error) {
	if pos.Q < 0 || pos.R < 0 || pos.Q >= g.width || pos.R >= g.height {
		return nil, fmt.Errorf("cell at position %s is out of bounds", pos)
	}
	return g.GetCellAt(pos.Q, pos.R)
}
func (g *concreteGrid) GetCellAtIndex(index int) (Cell, error) {
	if index < 0 || index >= g.GetCellCount() {
		return nil, fmt.Errorf("cell at index %d is out of bounds", index)
	}
	q := index % g.width
	r := index / g.width
	return g.GetCellAt(q, r)
}
func (g *concreteGrid) GetCellCount() int {
	return g.width * g.height
}
func (g *concreteGrid) GetWidth() int {
	return g.width
}
func (g *concreteGrid) GetHeight() int {
	return g.height
}

// CopyCellsTo copies the grid's cells into the destination slice.
// It returns an error if the destination slice has incorrect dimensions or is nil.
func (g *concreteGrid) CopyCellsTo(destination [][]Cell) error {
	if destination == nil {
		return fmt.Errorf("destination slice cannot be nil")
	}
	if len(destination) != g.height {
		return fmt.Errorf("destination has %d rows, expected %d", len(destination), g.height)
	}
	for r := 0; r < g.height; r++ {
		if destination[r] == nil {
			// Or alternatively, initialize it: destination[r] = make([]Cell, g.width)
			// For now, requiring pre-initialized rows.
			return fmt.Errorf("destination row %d cannot be nil, expected a pre-initialized row of width %d", r, g.width)
		}
		if len(destination[r]) != g.width {
			return fmt.Errorf("destination row %d has %d columns, expected %d", r, len(destination[r]), g.width)
		}
		for q := 0; q < g.width; q++ {
			destination[r][q] = g.cells[r][q]
		}
	}
	return nil
}

// NewGrid creates a new Grid with the specified position, name, width, height, and cells.
func NewGrid(position Position, name string, width, height int, cells [][]Cell) Grid {
	return &concreteGrid{
		position: position,
		name:     name,
		width:    width,
		height:   height,
		cells:    cells,
	}
}

// String implements the Stringer interface for Grid.
func (g *concreteGrid) String() string {
	return fmt.Sprintf("Grid(name: %s, position: %s, width: %d, height: %d)", g.name, g.position, g.width, g.height)
}

// GenerateGridFunc is a function type that generates a Grid.
type GenerateGridFunc func(position Position, name string, width, height int) (Grid, error)
