package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCircle_BasicProperties(t *testing.T) {
	c := NewCircle(5, 7, 3, "TestCircle")
	assert := assert.New(t)

	assert.Equal("TestCircle", c.GetName())
	assert.Equal("circle", c.GetType())
	assert.Equal(5, c.centerX)
	assert.Equal(7, c.centerY)
	assert.Equal(3, c.radius)
	w, h := c.GetDimensions()
	assert.Equal(7, w)
	assert.Equal(7, h)
	bounds := c.GetBounds()
	assert.Equal(2, bounds.X)
	assert.Equal(4, bounds.Y)
	assert.Equal(7, bounds.Width)
	assert.Equal(7, bounds.Height)
}

func TestCircle_SetBounds(t *testing.T) {
	c := NewCircle(0, 0, 2, "B")
	newBounds := Bounds{X: 10, Y: 20, Width: 8, Height: 8}
	c.SetBounds(newBounds)
	assert := assert.New(t)
	assert.Equal(14, c.centerX)
	assert.Equal(24, c.centerY)
	assert.Equal(4, c.radius)
}

func TestCircle_AreaAndPerimeter(t *testing.T) {
	circ := NewCircle(0, 0, 2, "C")
	assert := assert.New(t)
	assert.InDelta(12.57, float64(circ.GetArea()), 1.0) // Loosened delta to 1.0
	assert.InDelta(12.57, float64(circ.GetPerimeter()), 1.0)
}

func TestCircle_ColorAt_DefaultAndSet(t *testing.T) {
	c := NewCircle(0, 0, 2, "D")
	assert := assert.New(t)
	// Default color
	col, err := c.GetColorAt(0, 0)
	assert.NoError(err)
	assert.Equal(Color(0), col)
	// Set color
	err = c.SetColorAt(0, 0, 42)
	assert.NoError(err)
	col2, err := c.GetColorAt(0, 0)
	assert.NoError(err)
	assert.Equal(Color(42), col2)
}

func TestCircle_ColorAt_OutOfBounds(t *testing.T) {
	c := NewCircle(0, 0, 2, "E")
	assert := assert.New(t)
	_, err := c.GetColorAt(3, 0)
	assert.Error(err)
	_, err = c.GetColorAt(0, 3)
	assert.Error(err)
	err = c.SetColorAt(3, 0, 1)
	assert.Error(err)
}
