package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcreteCell_GetPosition(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		cell     Cell
		expected Position
	}{
		{
			name:     "Test case 1",
			cell:     NewCell(0, 0),
			expected: Position{Q: 0, R: 0},
		},
		{
			name:     "Test case 2",
			cell:     NewCell(1, 2),
			expected: Position{Q: 1, R: 2},
		},
		{
			name:     "Test case 3",
			cell:     NewCell(-1, -2),
			expected: Position{Q: -1, R: -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.expected, tt.cell.GetPosition(), "concreteCell.GetPosition() mismatch")
		})
	}
}

func TestNewCell(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		q        int
		r        int
		expected Position
	}{
		{
			name:     "Test case 1",
			q:        0,
			r:        0,
			expected: Position{Q: 0, R: 0},
		},
		{
			name:     "Test case 2",
			q:        1,
			r:        2,
			expected: Position{Q: 1, R: 2},
		},
		{
			name:     "Test case 3",
			q:        -1,
			r:        -2,
			expected: Position{Q: -1, R: -2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cell := NewCell(tt.q, tt.r)
			assert.Equal(tt.expected, cell.GetPosition(), "NewCell().GetPosition() mismatch")
		})
	}
}
