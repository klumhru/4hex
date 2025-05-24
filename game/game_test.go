package game

import (
	"testing"

	"github.com/klumhru/4hex/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGame(t *testing.T) {
	assert := assert.New(t)

	game := NewGame()
	assert.NotNil(game, "NewGame() should return a non-nil Game instance.")

	cg, ok := game.(*concreteGame)
	require.True(t, ok, "NewGame() should return a *concreteGame")

	assert.Nil(cg.gameMap, "Newly created game should have a nil map.")
	assert.Empty(cg.players, "Newly created game should have no players.")
}

func TestConcreteGame_SetMap(t *testing.T) {
	assert := assert.New(t)

	game := NewGame().(*concreteGame) // Assuming NewGame works as tested

	// Use hex.NewMap to create a concrete map instance
	testMap := hex.NewMap(10, 10)
	game.SetMap(testMap)

	assert.Equal(testMap, game.gameMap, "gameMap field should be set to the provided map.")

	// Test setting another map
	anotherMap := hex.NewMap(5, 5)
	game.SetMap(anotherMap)
	assert.Equal(anotherMap, game.gameMap, "gameMap field should be updated to the new map.")

	// Test setting map to nil
	game.SetMap(nil)
	assert.Nil(game.gameMap, "gameMap field should be nillable.")
}

func TestConcreteGame_AddPlayer(t *testing.T) {
	assert := assert.New(t)

	game := NewGame().(*concreteGame) // Assuming NewGame works

	// Use NewPlayer to create concrete player instances
	player1 := NewPlayer("Alice")
	game.AddPlayer(player1)

	assert.Len(game.players, 1, "Should have 1 player after adding one.")
	assert.Contains(game.players, player1, "The added player should be in the players slice.")

	player2 := NewPlayer("Bob")
	game.AddPlayer(player2)

	assert.Len(game.players, 2, "Should have 2 players after adding another.")
	assert.Contains(game.players, player1, "Player1 should still be in the players slice.")
	assert.Contains(game.players, player2, "Player2 should be in the players slice.")

	// Test adding the same player again (if allowed, or check for specific behavior if not)
	// Current implementation allows duplicates
	game.AddPlayer(player1)
	assert.Len(game.players, 3, "Should have 3 players after adding player1 again.")
	count := 0
	for _, p := range game.players {
		if p == player1 {
			count++
		}
	}
	assert.Equal(2, count, "Player1 should appear twice in the players slice.")
}

func TestConcreteGame_Integration(t *testing.T) {
	assert := assert.New(t)

	game := NewGame()

	// Set a map
	gameMap := hex.NewMap(20, 20)
	game.SetMap(gameMap)

	cg, ok := game.(*concreteGame)
	require.True(t, ok)
	assert.Equal(gameMap, cg.gameMap, "Map should be set correctly.")

	// Add players
	playerAlice := NewPlayer("Alice")
	playerBob := NewPlayer("Bob")

	game.AddPlayer(playerAlice)
	assert.Len(cg.players, 1)
	assert.Contains(cg.players, playerAlice)

	game.AddPlayer(playerBob)
	assert.Len(cg.players, 2)
	assert.Contains(cg.players, playerAlice)
	assert.Contains(cg.players, playerBob)
}
