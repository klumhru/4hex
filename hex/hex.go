package hex

import (
	"fmt"
)

// --- Interfaces ---

// Position represents a coordinate on the hex grid.
type Position struct {
	// Q and R are the axial coordinates of the hexagonal grid.
	Q int
	R int
}

// --- Implementations ---

// String implements the Stringer interface for Position.
func (p Position) String() string {
	return fmt.Sprintf("Pos(q:%d, r:%d)", p.Q, p.R)
}

// --- Factory Functions ---

// NewPosition creates a new Position.
func NewPosition(q, r int) Position {
	return Position{Q: q, R: r}
}
