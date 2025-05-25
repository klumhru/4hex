package shapes

import "fmt"

// Square implements the Shape interface as an axis-aligned rectangle with equal width and height.
type Square struct {
	name   string
	color  Color
	bounds Bounds
	data   map[[2]int]Color // for per-cell color
}

func NewSquare(x, y, size int, name string) *Square {
	data := make(map[[2]int]Color)
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			data[[2]int{x + col, y + row}] = 0 // Default color
		}
	}
	return &Square{
		name:   name,
		color:  0,
		bounds: Bounds{X: x, Y: y, Width: size, Height: size},
		data:   data,
	}
}

// Ensure Square implements Shape interface
var _ Shape = (*Square)(nil)

func (s *Square) GetBounds() Bounds {
	return s.bounds
}

func (s *Square) SetBounds(b Bounds) {
	s.bounds = b
}

func (s *Square) GetArea() int {
	return s.bounds.Width * s.bounds.Height
}

func (s *Square) GetPerimeter() int {
	return 4 * s.bounds.Width
}

func (s *Square) GetPosition() (int, int) {
	return s.bounds.X, s.bounds.Y
}

func (s *Square) GetDimensions() (int, int) {
	return s.bounds.Width, s.bounds.Height
}

func (s *Square) GetType() string {
	return "square"
}

func (s *Square) GetName() string {
	return s.name
}

func (s *Square) GetColorAt(x, y int) (Color, error) {
	if c, ok := s.data[[2]int{x, y}]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("coordinates out of bounds")
}

func (s *Square) SetColorAt(x, y int, color Color) error {
	if _, ok := s.data[[2]int{x, y}]; !ok {
		return fmt.Errorf("coordinates out of bounds")
	}
	s.data[[2]int{x, y}] = color
	return nil
}
