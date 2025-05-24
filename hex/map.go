package hex

import "fmt"

// Map represents the game map.
type Map interface {
	GetDimensions() (width int, height int)
	GetGrids() []Grid
	GetGridByName(name string) (Grid, error)
	GetGridByIndex(index int) (Grid, error)
	// AddGrid adds a new grid to the map.
	AddGrid(grid Grid) error
	// RemoveGrid removes a grid from the map by name.
	RemoveGrid(name string) error
	// RemoveGridByIndex removes a grid from the map by index.
	RemoveGridByIndex(index int) error
}

// concreteMap implements the Map interface.
type concreteMap struct {
	width  int
	height int
	grids  []Grid
}

func (m *concreteMap) GetDimensions() (int, int) {
	return m.width, m.height
}

func (m *concreteMap) GetGrids() []Grid {
	return m.grids
}

func (m *concreteMap) GetGridByName(name string) (Grid, error) {
	for _, grid := range m.grids {
		if grid.GetName() == name {
			return grid, nil
		}
	}
	return nil, fmt.Errorf("grid with name %s not found", name)
}

func (m *concreteMap) GetGridByIndex(index int) (Grid, error) {
	if index < 0 || index >= len(m.grids) {
		return nil, fmt.Errorf("index %d out of bounds", index)
	}
	return m.grids[index], nil
}

func (m *concreteMap) AddGrid(grid Grid) error {
	if grid == nil {
		return fmt.Errorf("cannot add nil grid")
	}
	m.grids = append(m.grids, grid)
	return nil
}

func (m *concreteMap) RemoveGrid(name string) error {
	for i, grid := range m.grids {
		if grid.GetName() == name {
			m.grids = append(m.grids[:i], m.grids[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("grid with name %s not found", name)
}

func (m *concreteMap) RemoveGridByIndex(index int) error {
	if index < 0 || index >= len(m.grids) {
		return fmt.Errorf("index %d out of bounds", index)
	}
	m.grids = append(m.grids[:index], m.grids[index+1:]...)
	return nil
}

// NewMap creates a new Map with the specified dimensions.
func NewMap(width, height int) Map {
	return &concreteMap{
		width:  width,
		height: height,
		grids:  make([]Grid, 0),
	}
}

// String implements the Stringer interface for concreteMap.
func (m *concreteMap) String() string {
	return fmt.Sprintf("Map(width: %d, height: %d, grids: %d)", m.width, m.height, len(m.grids))
}
