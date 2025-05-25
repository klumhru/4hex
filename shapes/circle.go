package shapes

import (
	"fmt"
	"math"
)

type Circle struct {
	name    string
	color   Color
	centerX int
	centerY int
	radius  int
	data    map[[2]int]Color // for per-cell color
}

func NewCircle(x, y, radius int, name string) *Circle {
	data := make(map[[2]int]Color)
	for row := -radius; row <= radius; row++ {
		for col := -radius; col <= radius; col++ {
			if col*col+row*row <= radius*radius {
				data[[2]int{x + col, y + row}] = 0 // Default color
			}
		}
	}
	return &Circle{
		name:    name,
		color:   0,
		centerX: x,
		centerY: y,
		radius:  radius,
		data:    data,
	}
}

var _ Shape = (*Circle)(nil)

func (c *Circle) GetBounds() Bounds {
	return Bounds{
		X:      c.centerX - c.radius,
		Y:      c.centerY - c.radius,
		Width:  2*c.radius + 1,
		Height: 2*c.radius + 1,
	}
}

func (c *Circle) SetBounds(b Bounds) {
	// Adjust center and radius to fit new bounds (best effort)
	c.centerX = b.X + b.Width/2
	c.centerY = b.Y + b.Height/2
	c.radius = min(b.Width, b.Height) / 2
}

func (c *Circle) GetArea() int {
	return int(math.Round(math.Pi * float64(c.radius) * float64(c.radius)))
}

func (c *Circle) GetPerimeter() int {
	return int(math.Round(2 * math.Pi * float64(c.radius)))
}

func (c *Circle) GetPosition() (int, int) {
	return c.centerX, c.centerY
}

func (c *Circle) GetDimensions() (int, int) {
	return 2*c.radius + 1, 2*c.radius + 1
}

func (c *Circle) GetType() string {
	return "circle"
}

func (c *Circle) GetName() string {
	return c.name
}

func (c *Circle) GetColorAt(x, y int) (Color, error) {
	if col, ok := c.data[[2]int{x, y}]; ok {
		return col, nil
	}
	return 0, fmt.Errorf("coordinates out of bounds or outside circle shape")
}

func (c *Circle) SetColorAt(x, y int, color Color) error {
	if _, ok := c.data[[2]int{x, y}]; !ok {
		return fmt.Errorf("coordinates out of bounds or outside circle shape")
	}
	c.data[[2]int{x, y}] = color
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
