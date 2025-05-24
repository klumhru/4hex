package hex

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
