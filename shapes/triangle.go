package shapes

import "fmt"

// Triangle represents an equilateral triangle with the base horizontal at the bottom.
type Triangle struct {
	name   string
	color  Color
	bounds Bounds
	size   int // side length
	data   map[[2]int]Color
}

// NewTriangle creates an equilateral triangle with the base horizontal at the bottom.
// x, y specify the top vertex; size is the length of each side.
func NewTriangle(x, y, size int, name string) *Triangle {
	data := make(map[[2]int]Color)
	center := size / 2
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if col >= center-row && col <= center+row {
				data[[2]int{x + col, y + row}] = 0 // Default color
			}
		}
	}
	return &Triangle{
		name:   name,
		color:  0,
		bounds: Bounds{X: x, Y: y, Width: size, Height: size},
		size:   size,
		data:   data,
	}
}

// NewIsoscelesTriangle creates a centered isosceles triangle with a base of 2*height-1 and height 'height'.
// The top vertex is at (x+height-1, y), and the base is at row y+height-1, spanning columns x to x+2*height-2.
func NewIsoscelesTriangle(x, y, height int, name string) *Triangle {
	data := make(map[[2]int]Color)
	baseWidth := 2*height - 1
	for row := 0; row < height; row++ {
		left := x + height - 1 - row
		right := x + height - 1 + row
		for col := left; col <= right; col++ {
			data[[2]int{col, y + row}] = 0 // Default color
		}
	}
	return &Triangle{
		name:   name,
		color:  0,
		bounds: Bounds{X: x, Y: y, Width: baseWidth, Height: height},
		size:   height, // For isosceles, size is height
		data:   data,
	}
}

var _ Shape = (*Triangle)(nil)

func (t *Triangle) GetBounds() Bounds {
	return t.bounds
}

func (t *Triangle) SetBounds(b Bounds) {
	t.bounds = b
	t.size = triangleMin(b.Width, b.Height)
}

func (t *Triangle) GetArea() int {
	// Area of equilateral triangle: (sqrt(3)/4) * size^2
	return int(0.433 * float64(t.size*t.size))
}

func (t *Triangle) GetPerimeter() int {
	return 3 * t.size
}

func (t *Triangle) GetPosition() (int, int) {
	return t.bounds.X, t.bounds.Y
}

func (t *Triangle) GetDimensions() (int, int) {
	return t.bounds.Width, t.bounds.Height
}

func (t *Triangle) GetType() string {
	return "triangle"
}

func (t *Triangle) GetName() string {
	return t.name
}

func (t *Triangle) GetColorAt(x, y int) (Color, error) {
	if c, ok := t.data[[2]int{x, y}]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("coordinates out of bounds or outside triangle shape")
}

func (t *Triangle) SetColorAt(x, y int, color Color) error {
	if _, ok := t.data[[2]int{x, y}]; !ok {
		return fmt.Errorf("coordinates out of bounds or outside triangle shape")
	}
	t.data[[2]int{x, y}] = color
	return nil
}

// Rotate90 returns a new Triangle rotated 90 degrees clockwise.
func (t *Triangle) Rotate90() *Triangle {
	rotated := NewTriangle(t.bounds.X, t.bounds.Y, t.size, t.name+"_rot90")
	for y := t.bounds.Y; y < t.bounds.Y+t.size; y++ {
		for x := t.bounds.X; x < t.bounds.X+t.size; x++ {
			if _, err := t.GetColorAt(x, y); err == nil {
				row := y - t.bounds.Y
				col := x - t.bounds.X
				newX := t.bounds.X + t.size - 1 - row
				newY := t.bounds.Y + col
				color, _ := t.GetColorAt(x, y)
				rotated.SetColorAt(newX, newY, color)
			}
		}
	}
	rotated.color = t.color
	return rotated
}

// Flip returns a new Triangle flipped around its central vertical axis.
func (t *Triangle) Flip() *Triangle {
	flipped := NewTriangle(t.bounds.X, t.bounds.Y, t.size, t.name+"_flip")
	for y := t.bounds.Y; y < t.bounds.Y+t.size; y++ {
		for x := t.bounds.X; x < t.bounds.X+t.size; x++ {
			if _, err := t.GetColorAt(x, y); err == nil {
				col := x - t.bounds.X
				newCol := t.size - 1 - col
				newX := t.bounds.X + newCol
				newY := y
				color, _ := t.GetColorAt(x, y)
				flipped.SetColorAt(newX, newY, color)
			}
		}
	}
	flipped.color = t.color
	return flipped
}

func triangleMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
