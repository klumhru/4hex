package hex

// Map represents the game map.
type Map interface {
	GetDimensions() (width int, height int)
}

// concreteMap implements the Map interface.
type concreteMap struct {
	width  int
	height int
}

func (m *concreteMap) GetDimensions() (int, int) {
	return m.width, m.height
}
