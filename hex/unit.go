package hex

// Unit represents a game unit.
type Unit interface {
	// Move changes the unit's position to the specified position.
	Move(pos Position)
	// Position returns the current position of the unit.
	Position() Position
	// GetName returns the name of the unit.
	GetName() string
}

// concreteUnit implements the Unit interface.
type concreteUnit struct {
	name     string
	position Position
}

// NewUnit creates a new Unit.
func NewUnit(name string) Unit {
	return &concreteUnit{name: name}
}

// Move sets the unit's position.
func (u *concreteUnit) Move(pos Position) {
	u.position = pos
}

// Position returns the unit's current position.
func (u *concreteUnit) Position() Position {
	return u.position
}

// GetName returns the unit's name.
func (u *concreteUnit) GetName() string {
	return u.name
}
