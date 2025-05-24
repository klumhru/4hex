package hex

type Cell interface {
	// GetPosition returns the position of the cell in the grid.
	GetPosition() Position
}

type concreteCell struct {
	position Position
}

func (c *concreteCell) GetPosition() Position {
	return c.position
}

// NewCell creates a new Cell with the specified position.
func NewCell(q, r int) Cell {
	return &concreteCell{
		position: Position{Q: q, R: r},
	}
}
