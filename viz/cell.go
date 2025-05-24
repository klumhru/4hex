package viz

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/klumhru/4hex/hex"
)

// RenderCell renders a cell using its position and charmbracelet/lipgloss
// library for styling. It returns a string representation of the cell's position.
/*
A cell is represented as such:
 _____
/     \
| Q R |
\_____/

The cell's position is displayed in the format "Q R", where Q and R are the
// coordinates of the cell in the grid. The cell is styled using the
*/

const cellTemplate = `
  ___
 /   \
|%s|
 \___/
`

func RenderCell(cell hex.Cell) {
	position := cell.GetPosition()
	// Using lipgloss to style the cell
	cellStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(1, 2).
		Align(lipgloss.Center).
		Background(lipgloss.Color("240")) // Example background color

	cellContent := cellStyle.Render()

	cellContent += "\n" + lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Render(fmt.Sprintf(cellTemplate, fmt.Sprintf("%2d %2d", position.Q, position.R)))
	println(cellContent)
}
