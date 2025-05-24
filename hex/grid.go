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
	// GetCells returns a two-dimensional array of cells representing the grid.
	GetCells() [][]Cell
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
func (g *concreteGrid) GetCells() [][]Cell {
	return g.cells
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
