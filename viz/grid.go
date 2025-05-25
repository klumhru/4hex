package viz

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/klumhru/4hex/hex"
)

var (
	// Styles

	// cellStyle is the style used for rendering cells in the grid.
	cellStyle = lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Center) // Example background color
)

func RenderGrid(grid hex.Grid) {
	doc := strings.Builder{}
	rows := make([]string, grid.GetWidth())

	for q := range grid.GetWidth() {
		cols := make([]string, 0)
		for r := range grid.GetHeight() {
			cell, err := grid.GetCellAt(q, r)
			if err != nil {
				// Handle the error, e.g., log it or skip rendering this cell.
				println("Error getting cell at position:", q, r, "-", err.Error())
				continue
			}
			if cell != nil {
				// Render the cell and add it to the row.
				cols = append(cols, cellStyle.Render(fmt.Sprintf(cellTemplate, cell.GetPosition().Q, cell.GetPosition().R)))
			} else {
				// If the cell is nil, render placeholder content.
				cols = append(cols, cellStyle.Render(strings.Repeat(" ", len(strings.Split(cellTemplate, "\n")[2]))))
			}
		}

		if q%2 == 0 {
			// If the column index is odd, add an empty string to align the first column.
			cols = append([]string{"     "}, cols...)
		}

		rows[q] = lipgloss.JoinHorizontal(lipgloss.Center, cols...)
	}
	doc.WriteString(lipgloss.JoinVertical(lipgloss.Top, rows...))

	fmt.Println(doc.String())
}
