package main

import (
	"fmt"
	"testing"

	"github.com/klumhru/4hex/game"
	"github.com/klumhru/4hex/hex"
	"github.com/stretchr/testify/assert"
)

// Test_GameInitialization tests the initialization of a game, including map, player, and unit creation.
// It serves as a basic integration test to ensure that the components interact correctly.
func Test_GameInitialization(t *testing.T) {
	t.Run("Game Initialization", func(t *testing.T) {
		assert := assert.New(t)
		// Create a new game
		gameInstance := game.NewGame()

		// Create a new map
		m := hex.NewMap(10, 10)

		// Add the map to the game
		gameInstance.SetMap(m)

		// Create a new player
		p := game.NewPlayer("Player 1")

		// Add the player to the game
		gameInstance.AddPlayer(p)

		// Create a new unit
		warrior := hex.NewUnit("Warrior")

		// Add the unit to the player's army
		p.AddUnit(warrior)

		// Move the unit to a new position
		warrior.Move(hex.NewPosition(1, 1))

		// Print the unit's position
		fmt.Println("Warrior position:", warrior.Position())
		fmt.Println("Player name:", p.GetName())
		fmt.Println("Player units count:", p.GetUnitCount())

		width, height := m.GetDimensions()
		fmt.Printf("Map dimensions: %dx%d\n", width, height)
		// Check if the unit's position is correct
		assert.Equal(warrior.Position().Q, 1, "Expected unit Q position to be 1")
		assert.Equal(warrior.Position().R, 1, "Expected unit R position to be 1")

		// Check if the player's name is correct
		assert.Equal(p.GetName(), "Player 1", "Expected player name to be 'Player 1'")
		// Check if the player has one unit
		assert.Equal(p.GetUnitCount(), 1, "Expected player to have 1 unit")

		// Check the unit retrieved from the player
		retrievedUnit, err := p.GetUnitAt(0)
		assert.NoError(err, "Expected no error when getting unit at index 0")
		assert.NotNil(retrievedUnit, "Expected unit at index 0 to not be nil")
		if retrievedUnit != nil {
			assert.Equal(retrievedUnit.GetName(), "Warrior", "Expected retrieved unit name to be 'Warrior'")
			assert.Equal(retrievedUnit.Position(), warrior.Position(), "Expected retrieved unit position to match original unit")
		}

		// Check if the unit's name is correct (original warrior variable)
		assert.Equal(warrior.GetName(), "Warrior", "Expected unit name to be 'Warrior'")
		// Check if the map dimensions are correct
		assert.Equal(width, 10, "Expected map width to be 10")
		assert.Equal(height, 10, "Expected map height to be 10")
	})
}
