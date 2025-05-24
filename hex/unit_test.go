package hex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	assert := assert.New(t)

	unitName := "test_warrior"
	unit := NewUnit(unitName)

	assert.NotNil(unit, "NewUnit should return a non-nil unit.")

	// Check if it's actually a concreteUnit and get its properties
	cu, ok := unit.(*concreteUnit)
	assert.True(ok, "NewUnit should return a *concreteUnit instance.")
	assert.Equal(unitName, cu.GetName(), "GetName() should return the initial name.")
	assert.Equal(unitName, cu.name, "Internal name field should match.")

	// Default position should be the zero value for Position
	expectedPos := Position{}
	assert.Equal(expectedPos, cu.Position(), "Initial position should be the zero value.")
	assert.Equal(expectedPos, cu.position, "Internal position field should be the zero value.")
}

func TestConcreteUnit_MoveAndPosition(t *testing.T) {
	assert := assert.New(t)

	unit := NewUnit("test_scout")
	cu, _ := unit.(*concreteUnit) // Already tested NewUnit, so direct cast is fine here for convenience

	pos1 := Position{Q: 1, R: 2}
	cu.Move(pos1)
	assert.Equal(pos1, cu.Position(), "Position() should return the updated position after Move.")
	assert.Equal(pos1, cu.position, "Internal position field should be updated after Move.")

	pos2 := Position{Q: -3, R: 5}
	cu.Move(pos2)
	assert.Equal(pos2, cu.Position(), "Position() should return the new position after a second Move.")
	assert.Equal(pos2, cu.position, "Internal position field should reflect the second Move.")

	// Test moving to zero position
	zeroPos := Position{}
	cu.Move(zeroPos)
	assert.Equal(zeroPos, cu.Position(), "Position() should return the zero position after moving to it.")
	assert.Equal(zeroPos, cu.position, "Internal position field should be zero after moving to it.")
}

func TestConcreteUnit_GetName(t *testing.T) {
	assert := assert.New(t)

	unitName1 := "archer"
	unit1 := NewUnit(unitName1)
	assert.Equal(unitName1, unit1.GetName(), "GetName() should return the correct name for unit1.")

	unitName2 := "mage"
	unit2 := NewUnit(unitName2)
	assert.Equal(unitName2, unit2.GetName(), "GetName() should return the correct name for unit2.")

	assert.NotEqual(unit1.GetName(), unit2.GetName(), "Names of different units should be different.")
}

// Test to ensure Position method returns a copy, not a reference (if Position were a pointer type, which it isn't here)
// For struct types like Position, this is less of an issue as they are typically passed by value.
// However, good to be mindful.
func TestConcreteUnit_PositionReturnsCopy(t *testing.T) {
	assert := assert.New(t)
	unit := NewUnit("test_unit").(*concreteUnit)

	initialPos := Position{Q: 10, R: 20}
	unit.Move(initialPos)

	retrievedPos := unit.Position()
	assert.Equal(initialPos, retrievedPos, "Position should match the set position.")

	// Modify the retrievedPos. If Position() returned a reference, unit.position would also change.
	// Since Position is a struct, retrievedPos is a copy.
	retrievedPos.Q = 100
	retrievedPos.R = 200

	assert.Equal(initialPos, unit.Position(), "Internal unit position should not change when the retrieved copy is modified.")
	assert.Equal(initialPos, unit.position, "Internal unit.position field should remain unchanged.")
}
