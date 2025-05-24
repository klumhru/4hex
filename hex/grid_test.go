package hex

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcreteGrid_GetPosition(t *testing.T) {
	assert := assert.New(t)
	grid := NewGrid(Position{Q: 1, R: 2}, "testGrid", 10, 10, nil)
	assert.Equal(Position{Q: 1, R: 2}, grid.GetPosition())
}

func TestConcreteGrid_GetName(t *testing.T) {
	assert := assert.New(t)
	grid := NewGrid(Position{}, "testGrid", 10, 10, nil)
	assert.Equal("testGrid", grid.GetName())
}

func TestConcreteGrid_GetCellAt(t *testing.T) {
	assert := assert.New(t)
	cells := [][]Cell{
		{NewCell(0, 0), NewCell(1, 0)},
		{NewCell(0, 1), NewCell(1, 1)},
	}
	grid := NewGrid(Position{}, "testGrid", 2, 2, cells)

	tests := []struct {
		name    string
		q       int
		r       int
		want    Cell
		wantErr bool
	}{
		{"valid cell", 0, 0, cells[0][0], false},
		{"valid cell", 1, 1, cells[1][1], false},
		{"out of bounds q", 2, 0, nil, true},
		{"out of bounds r", 0, 2, nil, true},
		{"out of bounds negative q", -1, 0, nil, true},
		{"out of bounds negative r", 0, -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grid.GetCellAt(tt.q, tt.r)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
			} else {
				assert.NoError(err)
				assert.Equal(tt.want, got)
			}
		})
	}
}

func TestConcreteGrid_GetCellAtPosition(t *testing.T) {
	assert := assert.New(t)
	cells := [][]Cell{
		{NewCell(0, 0), NewCell(1, 0)},
		{NewCell(0, 1), NewCell(1, 1)},
	}
	grid := NewGrid(Position{}, "testGrid", 2, 2, cells)

	tests := []struct {
		name    string
		pos     Position
		want    Cell
		wantErr bool
	}{
		{"valid cell", Position{Q: 0, R: 0}, cells[0][0], false},
		{"valid cell", Position{Q: 1, R: 1}, cells[1][1], false},
		{"out of bounds q", Position{Q: 2, R: 0}, nil, true},
		{"out of bounds r", Position{Q: 0, R: 2}, nil, true},
		{"out of bounds negative q", Position{Q: -1, R: 0}, nil, true},
		{"out of bounds negative r", Position{Q: 0, R: -1}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grid.GetCellAtPosition(tt.pos)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
			} else {
				assert.NoError(err)
				assert.Equal(tt.want, got)
			}
		})
	}
}

func TestConcreteGrid_GetCellAtIndex(t *testing.T) {
	assert := assert.New(t)
	cells := [][]Cell{
		{NewCell(0, 0), NewCell(1, 0)}, // 0, 1
		{NewCell(0, 1), NewCell(1, 1)}, // 2, 3
	}
	grid := NewGrid(Position{}, "testGrid", 2, 2, cells)

	tests := []struct {
		name    string
		index   int
		want    Cell
		wantErr bool
	}{
		{"valid cell index 0", 0, cells[0][0], false},
		{"valid cell index 1", 1, cells[0][1], false},
		{"valid cell index 2", 2, cells[1][0], false},
		{"valid cell index 3", 3, cells[1][1], false},
		{"out of bounds positive", 4, nil, true},
		{"out of bounds negative", -1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := grid.GetCellAtIndex(tt.index)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
			} else {
				assert.NoError(err)
				assert.Equal(tt.want, got)
			}
		})
	}
}

func TestConcreteGrid_GetCellCount(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name   string
		width  int
		height int
		want   int
	}{
		{"2x2 grid", 2, 2, 4},
		{"3x5 grid", 3, 5, 15},
		{"0x5 grid", 0, 5, 0},
		{"5x0 grid", 5, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := NewGrid(Position{}, "test", tt.width, tt.height, make([][]Cell, tt.height))
			assert.Equal(tt.want, grid.GetCellCount())
		})
	}
}

func TestConcreteGrid_GetWidth(t *testing.T) {
	assert := assert.New(t)
	grid := NewGrid(Position{}, "testGrid", 5, 10, nil)
	assert.Equal(5, grid.GetWidth())
}

func TestConcreteGrid_GetHeight(t *testing.T) {
	assert := assert.New(t)
	grid := NewGrid(Position{}, "testGrid", 5, 10, nil)
	assert.Equal(10, grid.GetHeight())
}

func TestConcreteGrid_GetCells(t *testing.T) {
	assert := assert.New(t)
	cells := [][]Cell{
		{NewCell(0, 0), NewCell(1, 0)},
		{NewCell(0, 1), NewCell(1, 1)},
	}
	grid := NewGrid(Position{}, "testGrid", 2, 2, cells)
	assert.Equal(cells, grid.GetCells())
}

func TestNewGrid(t *testing.T) {
	assert := assert.New(t)
	pos := Position{Q: 1, R: 2}
	name := "myGrid"
	width := 3
	height := 4
	cells := make([][]Cell, height)
	for r := 0; r < height; r++ {
		cells[r] = make([]Cell, width)
		for q := 0; q < width; q++ {
			cells[r][q] = NewCell(q, r)
		}
	}

	grid := NewGrid(pos, name, width, height, cells)

	assert.Equal(pos, grid.GetPosition(), "Position mismatch")
	assert.Equal(name, grid.GetName(), "Name mismatch")
	assert.Equal(width, grid.GetWidth(), "Width mismatch")
	assert.Equal(height, grid.GetHeight(), "Height mismatch")
	assert.Equal(cells, grid.GetCells(), "Cells mismatch")

	// Test that it's a concreteGrid
	_, ok := grid.(*concreteGrid)
	assert.True(ok, "NewGrid should return a *concreteGrid")
}

func TestConcreteGrid_String(t *testing.T) {
	assert := assert.New(t)
	pos := Position{Q: 1, R: 2}
	name := "testGrid"
	width := 5
	height := 10
	grid := NewGrid(pos, name, width, height, nil)
	expectedString := fmt.Sprintf("Grid(name: %s, position: %s, width: %d, height: %d)", name, pos, width, height)
	assert.Equal(expectedString, fmt.Sprintf("%s", grid), "String representation mismatch")
}
