package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSquare_BasicProperties(t *testing.T) {
	sq := NewSquare(2, 3, 5, "TestSquare")
	assert := assert.New(t)

	assert.Equal("TestSquare", sq.GetName())
	assert.Equal("square", sq.GetType())
	assert.Equal(25, sq.GetArea())
	assert.Equal(20, sq.GetPerimeter())
	x, y := sq.GetPosition()
	assert.Equal(2, x)
	assert.Equal(3, y)
	w, h := sq.GetDimensions()
	assert.Equal(5, w)
	assert.Equal(5, h)
	bounds := sq.GetBounds()
	assert.Equal(2, bounds.X)
	assert.Equal(3, bounds.Y)
	assert.Equal(5, bounds.Width)
	assert.Equal(5, bounds.Height)
}

func TestSquare_SetBounds(t *testing.T) {
	sq := NewSquare(0, 0, 4, "B")
	newBounds := Bounds{X: 1, Y: 2, Width: 3, Height: 3}
	sq.SetBounds(newBounds)
	assert := assert.New(t)
	assert.Equal(newBounds, sq.GetBounds())
}

func TestSquare_ColorAt_DefaultAndSet(t *testing.T) {
	sq := NewSquare(0, 0, 3, "C")
	assert := assert.New(t)
	// Default color
	c, err := sq.GetColorAt(1, 1)
	assert.NoError(err)
	assert.Equal(Color(0), c)
	// Set color
	err = sq.SetColorAt(1, 1, 42)
	assert.NoError(err)
	c2, err := sq.GetColorAt(1, 1)
	assert.NoError(err)
	assert.Equal(Color(42), c2)
}

func TestSquare_ColorAt_OutOfBounds(t *testing.T) {
	sq := NewSquare(0, 0, 2, "D")
	assert := assert.New(t)
	_, err := sq.GetColorAt(-1, 0)
	assert.Error(err)
	_, err = sq.GetColorAt(0, -1)
	assert.Error(err)
	_, err = sq.GetColorAt(2, 0)
	assert.Error(err)
	_, err = sq.GetColorAt(0, 2)
	assert.Error(err)
	// SetColorAt out of bounds
	err = sq.SetColorAt(2, 2, 1)
	assert.Error(err)
}
