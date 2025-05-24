package game

import (
	"testing"

	"github.com/klumhru/4hex/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPlayer(t *testing.T) {
	assert := assert.New(t)
	playerName := "PlayerOne"

	player := NewPlayer(playerName)
	assert.NotNil(player, "NewPlayer should return a non-nil Player instance.")

	cp, ok := player.(*concretePlayer)
	require.True(t, ok, "NewPlayer should return a *concretePlayer instance.")

	assert.Equal(playerName, cp.name, "Internal name field should be set.")
	assert.Equal(playerName, cp.GetName(), "GetName() should return the correct name.")
	assert.Empty(cp.units, "Newly created player should have no units.")
	assert.Empty(cp.GetUnits(), "GetUnits() should return an empty slice for a new player.")
}

func TestConcretePlayer_AddUnit(t *testing.T) {
	assert := assert.New(t)
	player := NewPlayer("TestPlayer").(*concretePlayer) // Assuming NewPlayer works

	unit1 := hex.NewUnit("Warrior")
	player.AddUnit(unit1)

	assert.Len(player.units, 1, "Player should have 1 unit after adding one.")
	assert.Contains(player.units, unit1, "The added unit should be in the units slice.")
	assert.Equal(unit1, player.units[0], "The first unit should be the one added.")

	unit2 := hex.NewUnit("Archer")
	player.AddUnit(unit2)

	assert.Len(player.units, 2, "Player should have 2 units after adding another.")
	assert.Contains(player.units, unit1, "Unit1 should still be in the units slice.")
	assert.Contains(player.units, unit2, "Unit2 should be in the units slice.")

	// Test adding nil unit (if this should be handled or if it's assumed valid unit)
	// Current implementation of AddUnit would append a nil to the slice.
	// Depending on desired behavior, a check might be needed in AddUnit.
	// For now, testing current behavior:
	player.AddUnit(nil)
	assert.Len(player.units, 3, "Player should have 3 units after adding nil.")
	assert.Nil(player.units[2], "The third unit should be nil.")
}

func TestConcretePlayer_GetName(t *testing.T) {
	assert := assert.New(t)

	name1 := "Gandalf"
	player1 := NewPlayer(name1)
	assert.Equal(name1, player1.GetName(), "GetName() should return the name set at creation.")

	name2 := "Frodo"
	player2 := NewPlayer(name2)
	assert.Equal(name2, player2.GetName(), "GetName() should return the name set at creation for a different player.")
}

func TestConcretePlayer_GetUnits(t *testing.T) {
	assert := assert.New(t)
	player := NewPlayer("TestPlayerGetUnits").(*concretePlayer)

	assert.Empty(player.GetUnits(), "GetUnits() should return an empty slice initially.")

	unit1 := hex.NewUnit("Scout")
	player.AddUnit(unit1)

	retrievedUnits1 := player.GetUnits()
	assert.Len(retrievedUnits1, 1, "GetUnits() should return a slice with 1 unit.")
	assert.Contains(retrievedUnits1, unit1, "The retrieved slice should contain unit1.")

	unit2 := hex.NewUnit("Tank")
	player.AddUnit(unit2)

	retrievedUnits2 := player.GetUnits()
	assert.Len(retrievedUnits2, 2, "GetUnits() should return a slice with 2 units.")
	assert.Contains(retrievedUnits2, unit1, "The retrieved slice should still contain unit1.")
	assert.Contains(retrievedUnits2, unit2, "The retrieved slice should contain unit2.")

	// Test that GetUnits returns a copy of the slice, not a reference to the internal one.
	// Modifying the returned slice should not affect the player's internal units slice.
	if len(retrievedUnits2) > 0 {
		retrievedUnits2[0] = hex.NewUnit("ModifiedUnit") // Modify the copy
		internalUnits := player.GetUnits()               // Get a fresh copy of internal units
		assert.NotEqual(retrievedUnits2[0], internalUnits[0], "Modifying the returned slice should not affect the internal units.")
		assert.Equal(unit1, internalUnits[0], "The first internal unit should still be the original unit1.")
	}

	// Test GetUnits after adding a nil unit
	player.AddUnit(nil)
	retrievedUnits3 := player.GetUnits()
	assert.Len(retrievedUnits3, 3, "GetUnits() should return 3 units after adding nil.")
	assert.Nil(retrievedUnits3[2], "The third unit in the retrieved slice should be nil.")
}
