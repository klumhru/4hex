package viz

import (
	"encoding/json"

	"github.com/klumhru/4hex/generator"
	"github.com/klumhru/4hex/shapes"
)

func Run(opts *Options) {
	// Initialize the visualization with the provided options.
	// This function will set up the necessary components for the visualization.
	// For example, it might initialize a graphical window, load resources, etc.

	width, height := opts.Positional.Width, opts.Positional.Height

	// Placeholder for actual implementation
	s, err := json.Marshal(opts)
	if err != nil {
		println("Error marshalling options:", err)
		return
	}
	println("Visualization options:", string(s))

	var shape shapes.Shape
	size := width
	if height < width {
		size = height
	}
	switch opts.Positional.Shape {
	case "hexagonal":
		shape = shapes.NewSquare(0, 0, size, "TestHexGrid")
	case "circular":
		shape = shapes.NewCircle(0, 0, size, "TestCircularGrid")
	case "square":
		shape = shapes.NewSquare(0, 0, size, "TestSquareGrid")
	case "triangular":
		shape = shapes.NewTriangle(0, 0, size, "TestTriangularGrid")
	case "isoceles":
		shape = shapes.NewIsoscelesTriangle(0, 0, size, "TestIsocelesGrid")
	default:
		println("Unsupported shape:", opts.Positional.Shape)
		return
	}

	cells, err := generator.GridFromShape(shape)
	if err != nil {
		println("Error generating grid from shape:", err)
		return
	}
	// Render the generated grid.
	RenderGrid(cells)
}
