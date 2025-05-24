package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a new concreteGrid for testing
func newTestGrid(name string) Grid {
	// For map tests, we often don't need a fully populated grid.
	// Adjust width, height, and cells as needed if specific grid interactions are tested.
	return NewGrid(Position{Q: 0, R: 0}, name, 1, 1, nil) // Removed S field
}

func TestNewMap(t *testing.T) {
	assert := assert.New(t)
	width, height := 10, 20
	m := NewMap(width, height)

	assert.NotNil(t, m, "NewMap should return a non-nil map.")
	w, h := m.GetDimensions()
	assert.Equal(width, w, "Width should match.")
	assert.Equal(height, h, "Height should match.")
	assert.Empty(m.GetGrids(), "Newly created map should have no grids.")
}

func TestConcreteMap_GetDimensions(t *testing.T) {
	assert := assert.New(t)
	width, height := 5, 15
	m := NewMap(width, height)
	w, h := m.GetDimensions()
	assert.Equal(width, w)
	assert.Equal(height, h)
}

func TestConcreteMap_AddGrid(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(10, 10)
	grid1 := newTestGrid("grid1")
	grid2 := newTestGrid("grid2")

	err := m.AddGrid(grid1)
	assert.NoError(err)
	assert.Len(m.GetGrids(), 1, "Map should have 1 grid after adding one.")
	assert.Equal(grid1, m.GetGrids()[0], "The added grid should be in the map.")

	err = m.AddGrid(grid2)
	assert.NoError(err)
	assert.Len(m.GetGrids(), 2, "Map should have 2 grids after adding another.")

	// Test adding nil grid
	err = m.AddGrid(nil)
	assert.Error(err, "Adding a nil grid should return an error.")
	assert.Contains(err.Error(), "cannot add nil grid", "Error message should indicate nil grid.")
	assert.Len(m.GetGrids(), 2, "Map should still have 2 grids after attempting to add nil.")
}

func TestConcreteMap_GetGrids(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(5, 5)
	assert.Empty(m.GetGrids(), "Initially, grids should be empty.")

	grid1 := newTestGrid("g1")
	m.AddGrid(grid1)
	assert.Len(m.GetGrids(), 1)
	assert.Contains(m.GetGrids(), grid1)

	grid2 := newTestGrid("g2")
	m.AddGrid(grid2)
	assert.Len(m.GetGrids(), 2)
	assert.Contains(m.GetGrids(), grid1)
	assert.Contains(m.GetGrids(), grid2)
}

func TestConcreteMap_GetGridByName(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	m := NewMap(10, 10)
	grid1 := newTestGrid("gridAlpha")
	grid2 := newTestGrid("gridBeta")
	m.AddGrid(grid1)
	m.AddGrid(grid2)

	// Test getting existing grid
	retrievedGrid, err := m.GetGridByName("gridAlpha")
	require.NoError(err, "Should not error when getting existing grid by name.")
	assert.Equal(grid1, retrievedGrid, "Retrieved grid should match the one added.")

	retrievedGrid, err = m.GetGridByName("gridBeta")
	require.NoError(err)
	assert.Equal(grid2, retrievedGrid)

	// Test getting non-existent grid
	_, err = m.GetGridByName("nonExistentGrid")
	assert.Error(err, "Should error when getting non-existent grid by name.")
	assert.Contains(err.Error(), "grid with name nonExistentGrid not found", "Error message mismatch.")
}

func TestConcreteMap_GetGridByIndex(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	m := NewMap(10, 10)
	grid1 := newTestGrid("grid1")
	grid2 := newTestGrid("grid2")
	m.AddGrid(grid1)
	m.AddGrid(grid2)

	// Test getting existing grid by index
	retrievedGrid, err := m.GetGridByIndex(0)
	require.NoError(err, "Should not error when getting existing grid by index 0.")
	assert.Equal(grid1, retrievedGrid, "Retrieved grid at index 0 should match.")

	retrievedGrid, err = m.GetGridByIndex(1)
	require.NoError(err)
	assert.Equal(grid2, retrievedGrid, "Retrieved grid at index 1 should match.")

	// Test getting grid by out-of-bounds index
	_, err = m.GetGridByIndex(-1)
	assert.Error(err, "Should error for negative index.")
	assert.Contains(err.Error(), "index -1 out of bounds", "Error message mismatch for negative index.")

	_, err = m.GetGridByIndex(2)
	assert.Error(err, "Should error for index equal to length.")
	assert.Contains(err.Error(), "index 2 out of bounds", "Error message mismatch for out of bounds index.")

	_, err = m.GetGridByIndex(100)
	assert.Error(err, "Should error for large out-of-bounds index.")
	assert.Contains(err.Error(), "index 100 out of bounds", "Error message mismatch for large out of bounds index.")
}

func TestConcreteMap_RemoveGrid(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	m := NewMap(10, 10)
	grid1 := newTestGrid("gridOne")
	grid2 := newTestGrid("gridTwo")
	grid3 := newTestGrid("gridThree")
	m.AddGrid(grid1)
	m.AddGrid(grid2)
	m.AddGrid(grid3)

	// Remove existing grid
	err := m.RemoveGrid("gridTwo")
	require.NoError(err, "Should not error when removing an existing grid.")
	assert.Len(m.GetGrids(), 2, "Map should have 2 grids after removing one.")
	assert.NotContains(m.GetGrids(), grid2, "Removed grid should no longer be in the map.")
	assert.Contains(m.GetGrids(), grid1, "Grid1 should still be in the map.")
	assert.Contains(m.GetGrids(), grid3, "Grid3 should still be in the map.")

	// Try to remove already removed grid
	err = m.RemoveGrid("gridTwo")
	assert.Error(err, "Should error when trying to remove a non-existent grid.")
	assert.Contains(err.Error(), "grid with name gridTwo not found")
	assert.Len(m.GetGrids(), 2, "Grid count should remain 2.")

	// Remove another existing grid
	err = m.RemoveGrid("gridOne")
	require.NoError(err)
	assert.Len(m.GetGrids(), 1)
	assert.NotContains(m.GetGrids(), grid1)
	assert.Contains(m.GetGrids(), grid3)

	// Remove last grid
	err = m.RemoveGrid("gridThree")
	require.NoError(err)
	assert.Empty(m.GetGrids(), "Map should be empty after removing all grids.")

	// Try to remove from empty map
	err = m.RemoveGrid("gridOne")
	assert.Error(err, "Should error when trying to remove from an empty map.")
}

func TestConcreteMap_RemoveGridByIndex(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	m := NewMap(10, 10)
	grid1 := newTestGrid("g1")
	grid2 := newTestGrid("g2")
	grid3 := newTestGrid("g3")
	m.AddGrid(grid1)
	m.AddGrid(grid2)
	m.AddGrid(grid3) // grids: [g1, g2, g3]

	// Remove grid by index (middle)
	err := m.RemoveGridByIndex(1) // remove g2
	require.NoError(err, "Should not error when removing an existing grid by index.")
	assert.Len(m.GetGrids(), 2, "Map should have 2 grids after removing one.")
	assert.Equal(grid1, m.GetGrids()[0], "First element should be g1.")
	assert.Equal(grid3, m.GetGrids()[1], "Second element should be g3.")
	assert.NotContains(m.GetGrids(), grid2, "g2 should be removed.")

	// Remove grid by index (start)
	err = m.RemoveGridByIndex(0) // remove g1
	require.NoError(err)
	assert.Len(m.GetGrids(), 1)
	assert.Equal(grid3, m.GetGrids()[0], "First element should now be g3.")
	assert.NotContains(m.GetGrids(), grid1, "g1 should be removed.")

	// Add one back to test removing last element
	m.AddGrid(grid1)             // grids: [g3, g1]
	err = m.RemoveGridByIndex(1) // remove g1
	require.NoError(err)
	assert.Len(m.GetGrids(), 1)
	assert.Equal(grid3, m.GetGrids()[0])
	assert.NotContains(m.GetGrids(), grid1)

	// Remove last remaining grid
	err = m.RemoveGridByIndex(0) // remove g3
	require.NoError(err)
	assert.Empty(m.GetGrids(), "Map should be empty.")

	// Test removing by out-of-bounds index
	err = m.RemoveGridByIndex(-1)
	assert.Error(err, "Should error for negative index.")
	assert.Contains(err.Error(), "index -1 out of bounds")

	err = m.RemoveGridByIndex(0) // Map is empty
	assert.Error(err, "Should error for index 0 on empty map.")
	assert.Contains(err.Error(), "index 0 out of bounds")

	err = m.RemoveGridByIndex(100) // Map is empty
	assert.Error(err, "Should error for large out-of-bounds index on empty map.")
	assert.Contains(err.Error(), "index 100 out of bounds")

	// Test with one element
	m.AddGrid(grid1)             // grids: [g1]
	err = m.RemoveGridByIndex(1) // out of bounds
	assert.Error(err)
	assert.Contains(err.Error(), "index 1 out of bounds")
	assert.Len(m.GetGrids(), 1) // Should not have changed
}

func TestConcreteMap_String(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(5, 8)
	// Cast to *concreteMap to access String() method if it's not on the interface
	concreteM, ok := m.(*concreteMap)
	require.True(t, ok, "Map should be a *concreteMap")

	expectedString := "Map(width: 5, height: 8, grids: 0)"
	assert.Equal(expectedString, concreteM.String(), "String representation mismatch for empty map.")

	concreteM.AddGrid(newTestGrid("g1"))
	expectedString = "Map(width: 5, height: 8, grids: 1)"
	assert.Equal(expectedString, concreteM.String(), "String representation mismatch for map with one grid.")

	concreteM.AddGrid(newTestGrid("g2"))
	expectedString = "Map(width: 5, height: 8, grids: 2)"
	assert.Equal(expectedString, concreteM.String(), "String representation mismatch for map with two grids.")
}

// TestGetGrids_Implementation ensures GetGrids returns actual grids, not nil.
func TestGetGrids_Implementation(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(1, 1)
	grid1 := newTestGrid("grid1")
	m.AddGrid(grid1)

	grids := m.GetGrids()
	assert.NotNil(grids, "GetGrids() should not return nil if grids have been added.")
	assert.Len(grids, 1, "GetGrids() should return the added grids.")
	assert.Equal(grid1, grids[0])
}

// TestGetGridByName_Implementation ensures GetGridByName returns actual grid or error.
func TestGetGridByName_Implementation(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(1, 1)
	grid1 := newTestGrid("gridAlpha")
	m.AddGrid(grid1)

	// Test found
	foundGrid, err := m.GetGridByName("gridAlpha")
	assert.NoError(err)
	assert.Equal(grid1, foundGrid)

	// Test not found
	_, err = m.GetGridByName("gridBeta")
	assert.Error(err)
}

// TestGetGridByIndex_Implementation ensures GetGridByIndex returns actual grid or error.
func TestGetGridByIndex_Implementation(t *testing.T) {
	assert := assert.New(t)
	m := NewMap(1, 1)
	grid1 := newTestGrid("gridOne")
	m.AddGrid(grid1)

	// Test found
	foundGrid, err := m.GetGridByIndex(0)
	assert.NoError(err)
	assert.Equal(grid1, foundGrid)

	// Test not found (out of bounds)
	_, err = m.GetGridByIndex(1)
	assert.Error(err)

	_, err = m.GetGridByIndex(-1)
	assert.Error(err)
}
