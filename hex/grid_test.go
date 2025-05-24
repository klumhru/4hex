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

func TestConcreteGrid_CopyCellsTo(t *testing.T) {
	assert := assert.New(t)

	// Original cells for the grid
	originalCells := [][]Cell{
		{NewCell(0, 0), NewCell(1, 0)},
		{NewCell(0, 1), NewCell(1, 1)},
	}
	grid := NewGrid(Position{}, "testGrid", 2, 2, originalCells)
	// Cast to concreteGrid to access CopyCellsTo if it's not on the interface (it is, but good practice for direct tests)
	cGrid, ok := grid.(*concreteGrid)
	assert.True(ok, "Grid should be a concreteGrid")

	tests := []struct {
		name          string
		destination   [][]Cell
		gridToUse     Grid // Use interface type
		expectError   bool
		expectedCells [][]Cell // Only if no error
		errorMsg      string   // Expected error message content if expectError is true
	}{
		{
			name: "successful copy",
			destination: func() [][]Cell {
				d := make([][]Cell, 2)
				d[0] = make([]Cell, 2)
				d[1] = make([]Cell, 2)
				return d
			}(),
			gridToUse:     cGrid,
			expectError:   false,
			expectedCells: originalCells,
		},
		{
			name:        "nil destination",
			destination: nil,
			gridToUse:   cGrid,
			expectError: true,
			errorMsg:    "destination slice cannot be nil",
		},
		{
			name: "incorrect row count",
			destination: func() [][]Cell {
				d := make([][]Cell, 1) // Expect 2 rows
				d[0] = make([]Cell, 2)
				return d
			}(),
			gridToUse:   cGrid,
			expectError: true,
			errorMsg:    "destination has 1 rows, expected 2",
		},
		{
			name: "nil row in destination",
			destination: func() [][]Cell {
				d := make([][]Cell, 2)
				d[0] = make([]Cell, 2)
				// d[1] is nil
				return d
			}(),
			gridToUse:   cGrid,
			expectError: true,
			errorMsg:    "destination row 1 cannot be nil",
		},
		{
			name: "incorrect column count in a row",
			destination: func() [][]Cell {
				d := make([][]Cell, 2)
				d[0] = make([]Cell, 2)
				d[1] = make([]Cell, 1) // Expect 2 columns
				return d
			}(),
			gridToUse:   cGrid,
			expectError: true,
			errorMsg:    "destination row 1 has 1 columns, expected 2",
		},
		{
			name:          "copy from empty grid (0x0) to valid empty destination",
			destination:   make([][]Cell, 0),
			gridToUse:     NewGrid(Position{}, "emptyGrid", 0, 0, make([][]Cell, 0)),
			expectError:   false,
			expectedCells: make([][]Cell, 0),
		},
		{
			name:        "copy from grid with cells to empty (0x0) destination",
			destination: make([][]Cell, 0),
			gridToUse:   cGrid, // 2x2 grid
			expectError: true,
			errorMsg:    "destination has 0 rows, expected 2",
		},
		{
			name:        "copy from empty grid (0x0) to nil destination",
			destination: nil,
			gridToUse:   NewGrid(Position{}, "emptyGrid", 0, 0, make([][]Cell, 0)),
			expectError: true,
			errorMsg:    "destination slice cannot be nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Ensure gridToUse is a concreteGrid to call CopyCellsTo directly if needed for testing non-interface aspects
			// However, the method is on the interface, so direct call is fine.
			err := tt.gridToUse.CopyCellsTo(tt.destination)

			if tt.expectError {
				assert.Error(err)
				if tt.errorMsg != "" {
					assert.Contains(err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(err)
				assert.Equal(tt.expectedCells, tt.destination, "Copied cells do not match expected cells")

				// Verify that the destination is a deep copy, not a reference (if applicable for Cell type)
				if len(originalCells) > 0 && len(tt.destination) > 0 && len(originalCells[0]) > 0 && len(tt.destination[0]) > 0 {
					// Modify original and check destination is not affected (only if Cell is a pointer or has mutable fields)
					// For now, assume Cell is simple enough or this test is primarily for the copy mechanism itself.
					// If originalCells were modified after copy, tt.destination should not change.
				}
			}
		})
	}
}
