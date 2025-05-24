package hex

import (
	"fmt"
)

// --- Interfaces ---

// Position represents a coordinate on the hex grid.
type Position struct {
	Q int
	R int
}

// Unit represents a game unit.
type Unit interface {
	Move(pos Position)
	Position() Position
	GetName() string
}

// --- Implementations ---

// String provides a string representation for Position.
func (p Position) String() string {
	return fmt.Sprintf("Pos(q:%d, r:%d)", p.Q, p.R)
}

// --- Factory Functions ---

// NewPosition creates a new Position.
func NewPosition(q, r int) Position {
	return Position{Q: q, R: r}
}

// NewMap creates a new Map.
func NewMap(width, height int) Map {
	return &concreteMap{width: width, height: height}
}
