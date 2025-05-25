package shapes

type Color int16

type Bounds struct {
	X      int `json:"x" description:"X coordinate of the shape's top-left corner"`
	Y      int `json:"y" description:"Y coordinate of the shape's top-left corner"`
	Width  int `json:"width" description:"Width of the shape"`
	Height int `json:"height" description:"Height of the shape"`
}

type Shape interface {
	GetBounds() Bounds // Returns the bounding box of the shape
	SetBounds(Bounds)  // Sets the bounding box of the shape
	GetArea() int
	GetPerimeter() int
	GetPosition() (int, int)                // Returns the top-left corner position (x, y)
	GetDimensions() (int, int)              // Returns the width and height of the shape
	GetType() string                        // Returns the type of the shape (e.g., "rectangle", "circle")
	GetName() string                        // Returns the name of the shape
	GetColorAt(x, y int) (Color, error)     // Returns the color at a specific point (x, y)
	SetColorAt(x, y int, color Color) error // Sets the color at a specific point (x, y)
}
