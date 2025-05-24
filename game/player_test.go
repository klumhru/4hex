package game

import (
	"testing"

	"github.com/klumhru/4hex/hex"
	"github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
	name := "Player1"
	p := NewPlayer(name)

	assert.NotNil(t, p, "NewPlayer should not return nil")
	cp, ok := p.(*concretePlayer)
	assert.True(t, ok, "NewPlayer should return a concretePlayer")
	assert.Equal(t, name, cp.GetName(), "GetName should return the correct name")
	assert.Equal(t, 0, cp.GetUnitCount(), "A new player should have 0 units")
}

func TestConcretePlayer_AddUnit(t *testing.T) {
	p := NewPlayer("TestPlayer").(*concretePlayer)
	u1 := hex.NewUnit("Warrior")
	u2 := hex.NewUnit("Archer")

	p.AddUnit(u1)
	assert.Equal(t, 1, p.GetUnitCount(), "Unit count should be 1 after adding one unit")
	unit, err := p.GetUnitAt(0)
	assert.NoError(t, err)
	assert.Equal(t, u1, unit, "GetUnitAt(0) should return the first unit added")

	p.AddUnit(u2)
	assert.Equal(t, 2, p.GetUnitCount(), "Unit count should be 2 after adding two units")
	unit, err = p.GetUnitAt(1)
	assert.NoError(t, err)
	assert.Equal(t, u2, unit, "GetUnitAt(1) should return the second unit added")
}

func TestConcretePlayer_GetName(t *testing.T) {
	name := "PlayerX"
	p := NewPlayer(name).(*concretePlayer)
	assert.Equal(t, name, p.GetName(), "GetName should return the player's name")
}

func TestConcretePlayer_GetUnitCount(t *testing.T) {
	p := NewPlayer("TestPlayer").(*concretePlayer)
	assert.Equal(t, 0, p.GetUnitCount(), "Initially, unit count should be 0")

	u1 := hex.NewUnit("Warrior")
	p.AddUnit(u1)
	assert.Equal(t, 1, p.GetUnitCount(), "Unit count should be 1 after adding one unit")

	u2 := hex.NewUnit("Archer")
	p.AddUnit(u2)
	assert.Equal(t, 2, p.GetUnitCount(), "Unit count should be 2 after adding two units")
}

func TestConcretePlayer_GetUnitAt(t *testing.T) {
	p := NewPlayer("TestPlayer").(*concretePlayer)
	u1 := hex.NewUnit("Warrior")
	u2 := hex.NewUnit("Archer")

	p.AddUnit(u1)
	p.AddUnit(u2)

	// Test valid indices
	unit, err := p.GetUnitAt(0)
	assert.NoError(t, err, "GetUnitAt(0) should not return an error")
	assert.Equal(t, u1, unit, "GetUnitAt(0) should return the first unit")

	unit, err = p.GetUnitAt(1)
	assert.NoError(t, err, "GetUnitAt(1) should not return an error")
	assert.Equal(t, u2, unit, "GetUnitAt(1) should return the second unit")

	// Test invalid indices
	_, err = p.GetUnitAt(-1)
	assert.Error(t, err, "GetUnitAt(-1) should return an error for negative index")

	_, err = p.GetUnitAt(2)
	assert.Error(t, err, "GetUnitAt(2) should return an error for out-of-bounds index")

	// Test with no units
	pEmpty := NewPlayer("EmptyPlayer").(*concretePlayer)
	_, err = pEmpty.GetUnitAt(0)
	assert.Error(t, err, "GetUnitAt(0) on an empty player should return an error")
}

func TestConcretePlayer_SetUnitAt(t *testing.T) {
	p := NewPlayer("TestPlayer").(*concretePlayer)
	u1 := hex.NewUnit("Warrior")
	u2 := hex.NewUnit("Archer")
	u3 := hex.NewUnit("Mage")

	p.AddUnit(u1)
	p.AddUnit(u2)

	// Test valid index
	err := p.SetUnitAt(0, u3)
	assert.NoError(t, err, "SetUnitAt(0) should not return an error")
	unit, _ := p.GetUnitAt(0)
	assert.Equal(t, u3, unit, "GetUnitAt(0) should return the new unit u3")

	unit, _ = p.GetUnitAt(1)
	assert.Equal(t, u2, unit, "GetUnitAt(1) should still return u2")

	// Test invalid indices
	err = p.SetUnitAt(-1, u3)
	assert.Error(t, err, "SetUnitAt(-1) should return an error for negative index")

	err = p.SetUnitAt(2, u3)
	assert.Error(t, err, "SetUnitAt(2) should return an error for out-of-bounds index")

	// Test with no units
	pEmpty := NewPlayer("EmptyPlayer").(*concretePlayer)
	err = pEmpty.SetUnitAt(0, u3)
	assert.Error(t, err, "SetUnitAt(0) on an empty player should return an error")
}
