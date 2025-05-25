package shapes

import "fmt"

// Rectangle implements the Shape interface as an axis-aligned rectangle.
type Rectangle struct {
	name   string
	color  Color
	bounds Bounds
	data   map[[2]int]Color // for per-cell color
}

func NewRectangle(x, y, width, height int, name string) *Rectangle {
	data := make(map[[2]int]Color)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			data[[2]int{x + col, y + row}] = 0 // Default color
		}
	}
	return &Rectangle{
		name:   name,
		color:  0,
		bounds: Bounds{X: x, Y: y, Width: width, Height: height},
		data:   data,
	}
}

var _ Shape = (*Rectangle)(nil)

func (r *Rectangle) GetBounds() Bounds {
	return r.bounds
}

func (r *Rectangle) SetBounds(b Bounds) {
	r.bounds = b
}

func (r *Rectangle) GetArea() int {
	return r.bounds.Width * r.bounds.Height
}

func (r *Rectangle) GetPerimeter() int {
	return 2 * (r.bounds.Width + r.bounds.Height)
}

func (r *Rectangle) GetPosition() (int, int) {
	return r.bounds.X, r.bounds.Y
}

func (r *Rectangle) GetDimensions() (int, int) {
	return r.bounds.Width, r.bounds.Height
}

func (r *Rectangle) GetType() string {
	return "rectangle"
}

func (r *Rectangle) GetName() string {
	return r.name
}

func (r *Rectangle) GetColorAt(x, y int) (Color, error) {
	if c, ok := r.data[[2]int{x, y}]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("coordinates out of bounds")
}

func (r *Rectangle) SetColorAt(x, y int, color Color) error {
	if _, ok := r.data[[2]int{x, y}]; !ok {
		return fmt.Errorf("coordinates out of bounds")
	}
	r.data[[2]int{x, y}] = color
	return nil
}
